package service

import (
	"../../services"
	"../database"
	"../database/queries"
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type accountServer struct{}

func (s *accountServer) Fetch(ctx context.Context, request *services.AccountFetchRequest) (*services.AccountFetchResponse, error) {
	id := request.GetId()

	db := database.Connection

	stmt := queries.FetchAccount(db)
	defer stmt.Close()

	var name, email string
	var createdAt time.Time

	if err := stmt.QueryRow(id).Scan(&name, &email, &createdAt); err == nil {
		convertedTime, _ := ptypes.TimestampProto(createdAt)
		return &services.AccountFetchResponse{
			Id:        id,
			Name:      name,
			Email:     email,
			CreatedAt: convertedTime,
		}, nil
	} else if err == sql.ErrNoRows {
		return nil, status.Error(codes.Unknown, "No account found with that email and password")
	} else {
		return nil, status.Error(codes.Internal, "Couldn't access the database, please try again later")
	}
}

func (s *accountServer) Login(ctx context.Context, request *services.AccountLoginRequest) (*services.AccountLoginResponse, error) {
	// Get data from request body
	email, pass := request.GetEmail(), request.GetPassword()

	// Get a database connection
	db := database.Connection

	// Create statement: Check if an account is available with a specific email and password
	stmt := queries.AuthenticateAccount(db)
	defer stmt.Close()

	var name string

	// Execute statement: Check if an account is available
	if err := stmt.QueryRow(email, pass).Scan(&name, &email); err == nil {
		loginAt, _ := ptypes.TimestampProto(time.Now())
		return &services.AccountLoginResponse{
			Name:    name,
			Email:   email,
			LoginAt: loginAt,
		}, nil
	} else if err == sql.ErrNoRows {
		return nil, status.Error(codes.Unknown, "No account found with that email and password")
	} else {
		return nil, status.Error(codes.Internal, "Couldn't access the database, please try again later")
	}
}

func (a *accountServer) Register(ctx context.Context, request *services.AccountRegisterRequest) (*services.AccountRegisterResponse, error) {
	// Get data from the request body
	email := request.GetEmail()

	// Get a new database connection
	db := database.Connection

	// Create statement: check if account is available
	stmt := queries.IsAccountAvailable(db)
	defer stmt.Close()

	// Execute statement: check if account is available
	if err := stmt.QueryRow(email).Scan(&email); err == nil {
		return &services.AccountRegisterResponse{}, status.Error(codes.AlreadyExists, "Email already registered")
	}

	// Get data from the request body
	name, pass, passConfirm := request.GetName(), request.GetPassword(), request.GetPasswordConfirmation()

	// Check if both passwords are equal
	if pass != passConfirm {
		return nil, status.Error(codes.FailedPrecondition, "Passwords does not match")
	}

	// Create statement: Insert account into database
	stmt = queries.RegisterAccount(db)
	defer stmt.Close()

	// Execute statement: Insert account into database and return the last inserted id
	var accountID int64

	if err := stmt.QueryRow(name, email, pass).Scan(&accountID); err != nil {
		return nil, status.Error(codes.Internal, "Couldn't register account, please try again later")
	}

	// Get the current time
	if registeredAt, err := ptypes.TimestampProto(time.Now()); err != nil {
		return nil, status.Error(codes.Internal, "Couldn't retrieve server time, please try again later")
	} else {
		return &services.AccountRegisterResponse{
			Id:         accountID,
			Name:       name,
			Email:      email,
			RegisterAt: registeredAt,
		}, nil
	}
}

func InitializeAccountServer() *accountServer {
	return new(accountServer)
}
