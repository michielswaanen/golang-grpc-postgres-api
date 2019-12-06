package service

import (
	"../../services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

func Login(ctx *gin.Context) {
	client := newAccountServiceClient()

	var loginRequest services.AccountLoginRequest

	err := ctx.BindJSON(loginRequest)

	if err != nil {
		panic(err)
	}

	if res, err := client.Login(ctx, &loginRequest); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"id":      res.GetId(),
			"name":    res.GetName(),
			"email":   res.GetEmail(),
			"loginAt": res.GetLoginAt(),
		})
	}
}

func Register(ctx *gin.Context) {
	client := newAccountServiceClient()

	var registerRequest services.AccountRegisterRequest

	err := ctx.BindJSON(registerRequest)

	if err != nil {
		panic(err)
	}

	if res, err := client.Register(ctx, &registerRequest); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"id":         res.GetId(),
			"name":       res.GetName(),
			"email":      res.GetEmail(),
			"registerAt": res.GetRegisterAt(),
		})
	}
}

func newAccountServiceClient() services.AccountServiceClient {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	return services.NewAccountServiceClient(conn)
}
