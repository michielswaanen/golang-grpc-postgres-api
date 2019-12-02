package service

import (
	"../../services"
	"../database"
	"../database/queries"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"log"
)

type accountServer struct{}

func (s *accountServer) Register(context.Context, *services.AccountRegisterRequest) (*services.AccountRegisterResponse, error) {
	panic("implement me")
}

func (s *accountServer) Login(ctx context.Context, request *services.AccountLoginRequest) (*services.AccountLoginResponse, error) {
	email, pass := request.GetEmail(), request.GetPassword()

	db := database.GetConnection()

	stmt := queries.AuthenticateAccount(db)
	defer stmt.Close()

	rows, err := stmt.Query(email, pass)

	if err != nil {
		return &services.AccountLoginResponse{
			Name:         nil,
			Email:        nil,
			LoginAt:      nil,
		}, errors.New("email or password is wrong")
	}

	var name string
	var loginAt timestamp.Timestamp

	err = rows.Scan(&name, &email, &loginAt)

	if err != nil {
		log.Fatal(err)
	}

	return &services.AccountLoginResponse{
		Name:         name,
		Email:        email,
		LoginAt:      &loginAt,
	}, nil
}

func InitializeAccountServer() *accountServer {
	return new(accountServer)
}
