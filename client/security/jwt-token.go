package security

import (
	"../../services"
	"errors"
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
		"role":      "user",
		"expiresAt": time.Now().Add(time.Minute * 15),
	})

	if tokenString, err := token.SignedString([]byte("SECRET")); err != nil {
		return "", status.Error(codes.Internal, "Couldn't create a token, try again later")
	} else {
		return tokenString, nil
	}
}

func AuthenticationRequired(authorizationType ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token, err := getAuthToken(ctx); err == nil {
			if valid, err := isTokenValid(token, authorizationType...); valid {
				ctx.Next()
			} else {
				ctx.AbortWithStatusJSON(http.StatusNetworkAuthenticationRequired, gin.H{"error": err.Error()})
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusNetworkAuthenticationRequired, gin.H{"error": err.Error()})
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

func isTokenValid(tokenString string, validAuthTypes ...string) (bool, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("SECRET"), nil
	})

	if !hasPermission(claims, validAuthTypes...) {
		fmt.Println("Token err", err)
		return false, errors.New("No permission to access this endpoint")
	} else {
		return err == nil && token.Valid, errors.New("No or wrong authentication token provided")
	}

}

func hasPermission(claims jwt.MapClaims, validAuthTypes ...string) bool {
	fmt.Println("Authenticating token...")

	result := false

	for _, auth := range validAuthTypes {
		if claims["role"] == auth {
			result = true
		}
	}

	if expiresAt, ok := claims["expiresAt"].(string); ok {
		if t, err := time.Parse(time.RFC3339, expiresAt); err == nil {
			result = t.After(time.Now())
		}
	}

	fmt.Println("Result", result)

	return result
}
