package service

import (
	"../../services"
	"../security"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

func Login(ctx *gin.Context) {
	client := newAccountServiceClient()

	var loginRequest services.AccountLoginRequest

	if err := ctx.BindJSON(&loginRequest); err != nil {
		panic(err)
	}

	if res, err := client.Login(ctx, &loginRequest); err == nil {
		if token, err := security.CreateToken(res); err == nil {
			ctx.Header("Authorization", "Bearer "+token)
			ctx.JSON(http.StatusOK, gin.H{
				"id":      res.GetId(),
				"name":    res.GetName(),
				"email":   res.GetEmail(),
				"loginAt": res.GetLoginAt(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
	} else if status.Code(err) == codes.Unknown {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
	} else {
		panic(err)
	}
}

func Register(ctx *gin.Context) {
	client := newAccountServiceClient()

	var registerRequest services.AccountRegisterRequest

	if err := ctx.BindJSON(&registerRequest); err != nil {
		panic(err)
	}

	if res, err := client.Register(ctx, &registerRequest); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"id":         res.GetId(),
			"name":       res.GetName(),
			"email":      res.GetEmail(),
			"registerAt": res.GetRegisterAt(),
		})
	} else if status.Code(err) == codes.AlreadyExists {
		ctx.AbortWithStatusJSON(http.StatusConflict, err)
	} else if status.Code(err) == codes.FailedPrecondition {
		ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, err)
	} else if status.Code(err) == codes.Internal {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		panic(err)
	}
}

func Fetch(ctx *gin.Context) {
	client := newAccountServiceClient()

	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		fetchRequest := services.AccountFetchRequest{
			Id: int64(id),
		}

		if res, err := client.Fetch(ctx, &fetchRequest); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"id":        res.GetId(),
				"name":      res.GetName(),
				"email":     res.GetEmail(),
				"createdAt": res.GetCreatedAt(),
			})
		} else if status.Code(err) == codes.Unknown {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		} else if status.Code(err) == codes.Internal {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		} else {
			panic(err)
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, status.Error(codes.FailedPrecondition, "Provided ID is not of type integer"))
	}
}

func newAccountServiceClient() services.AccountServiceClient {
	if conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure()); err != nil {
		panic(err)
	} else {
		return services.NewAccountServiceClient(conn)
	}
}
