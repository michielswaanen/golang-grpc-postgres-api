package service

import (
	"../../services"
	"../database"
	"../database/queries"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type accountServer struct{}

func (s *accountServer) Login(ctx context.Context, request *services.AccountLoginRequest) (*services.AccountLoginResponse, error) {
	// Get data from request body
	email, pass := request.GetEmail(), request.GetPassword()

	// Create a database connection
	db := database.GetConnection()

	// Create statement: Check if an account is available with a specific email and password
	stmt := queries.AuthenticateAccount(db)
	defer stmt.Close()

	// Execute statement: Check if an account is available
	rows, err := stmt.Query(email, pass)

	// No account found with the specific account and password
	if err != nil {
		return &services.AccountLoginResponse{}, errors.New("email or password is wrong")
	}

	// Set name and email variables to the result retrieved by the query
	var name string
	err = rows.Scan(&name, &email)
	loginAt, _ := ptypes.TimestampProto(time.Now())

	checkError(err)

	return &services.AccountLoginResponse{
		Name:    name,
		Email:   email,
		LoginAt: loginAt,
	}, nil
}

func (a *accountServer) Register(ctx context.Context, request *services.AccountRegisterRequest) (*services.AccountRegisterResponse, error) {
	// Get data from the request body
	email := request.GetEmail()

	// Create a new database connection
	db := database.GetConnection()

	// Create statement: check if account is available
	stmt := queries.IsAccountAvailable(db)
	defer stmt.Close()

	// Execute statement: check if account is available
	rows, err := stmt.Query(email)

	checkError(err)

	// If data was returned by the query (so the email exists), then there's no error.
	if err := rows.Scan(&email); err != nil {
		return &services.AccountRegisterResponse{}, errors.New("email already registered")
	}

	// Get data from the request body
	name, pass, passConfirm := request.GetName(), request.GetPassword(), request.GetPasswordConfirmation()

	// Check if both passwords are equal
	if pass != passConfirm {
		return &services.AccountRegisterResponse{}, errors.New("password does not match the password confirmation")
	}

	// Create statement: Insert account into database
	stmt = queries.RegisterAccount(db)
	defer stmt.Close()

	// Execute statement: Insert account into database and return the last inserted id
	var accountID int64
	err = stmt.QueryRow(name, email, pass).Scan(&accountID)

	checkError(err)

	// Get the current time
	registeredAt, _ := ptypes.TimestampProto(time.Now())

	return &services.AccountRegisterResponse{
		Id:         accountID,
		Name:       name,
		Email:      email,
		RegisterAt: registeredAt,
	}, nil
}

func InitializeAccountServer() *accountServer {
	return new(accountServer)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
