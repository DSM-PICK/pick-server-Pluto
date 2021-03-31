package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"pluto/presenter/web/exception"
)

func Auth(environment string) gin.HandlerFunc {
	return func(context *gin.Context) {
		_, err := jwtAuth(context, environment)

		if err != nil {
			switch err.Error() {
			case exception.InvalidAuthorizationException:
				exception.InvalidAuthorization(context)
			}
			return
		}

		context.Next()
	}
}

func jwtAuth(context *gin.Context, environment string) (jwt.MapClaims, error) {
	authorization := context.Request.Header.Get("Authorization")
	if len(authorization) == 0 || authorization[:7] != "Bearer " {
		return nil, fmt.Errorf(exception.InvalidAuthorizationException)
	}

	token, err := jwt.Parse(authorization[7:], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(environment)), nil
	})

	if err != nil {
		return nil, fmt.Errorf(exception.InvalidAuthorizationException)
	}

	return token.Claims.(jwt.MapClaims), nil
}
