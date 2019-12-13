package security

import (
	"../../services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
	"time"
)

func CreateToken(res *services.AccountLoginResponse) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        res.GetId(),
		"name":      res.GetName(),
		"email":     res.GetEmail(),
		"loginAt":   res.GetLoginAt(),
		"expiresAt": time.Now().Add(time.Minute * 15).Unix(),
	})

	if tokenString, err := token.SignedString([]byte("SECRET")); err != nil {
		fmt.Println(err)
		return "", status.Error(codes.Internal, "Couldn't create a token, try again later")
	} else {
		return tokenString, nil
	}
}

func AuthenticationRequired(authorizationType ...string) gin.HandlerFunc {
	// TODO: Check op expired time
	// TODO: Check op auth type
	// TODO: After x tries, wont be able to access for x seconds
	return func(ctx *gin.Context) {
		if token, err := getAuthToken(ctx); err != nil || !isTokenValid(token) {
			ctx.AbortWithStatusJSON(http.StatusNetworkAuthenticationRequired, gin.H{"error": "Authentication required"})
		} else {
			ctx.Next()
		}
	}
}

func getAuthToken(ctx *gin.Context) (string, error) {
	tokenString := ctx.GetHeader("Authorization")

	if strings.Contains(tokenString, "Bearer ") {
		return strings.Split(tokenString, " ")[1], nil
	} else {
		return "", status.Error(codes.FailedPrecondition, "No or wrong authentication token found")
	}
}

func isTokenValid(tokenString string, validAuthTypes ...string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("SECRET"), nil
	})

	return err == nil && token.Valid
}
