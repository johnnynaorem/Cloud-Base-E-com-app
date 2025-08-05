package jwt

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

type UserClaims struct {
	jwt.StandardClaims
	UserEmail string
}

func VerifyToken(tokenStr, secret string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func AuthMiddleware(secret string) iris.Handler {
	return func(ctx iris.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			ctx.StopWithStatus(401)
			return
		}
		tokenStr := strings.Split(auth, " ")[1]
		claims, err := VerifyToken(tokenStr, secret)
		if err != nil {
			ctx.StopWithStatus(401)
			return
		}
		ctx.Values().Set("user_email", claims.UserEmail)
		ctx.Next()
	}
}
