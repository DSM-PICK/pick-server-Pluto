package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"pluto/presenter/web/exception"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, err := jwtAuth(context)

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

func jwtAuth(context *gin.Context) (jwt.MapClaims, error) {
	authorization := context.Request.Header.Get("Authorization")
	if len(authorization) == 0 || authorization[:7] != "Bearer " {
		return nil, fmt.Errorf(exception.InvalidAuthorizationException)
	}

	token, err := jwt.Parse(authorization[7:], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil { return nil, fmt.Errorf(exception.InvalidAuthorizationException) }

	return token.Claims.(jwt.MapClaims), nil
}
