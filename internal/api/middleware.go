package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type jwtClaims struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	jwt.StandardClaims
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET not set")
	}

	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			return c.Redirect(http.StatusUnauthorized, "/login")
		}
		if cookie.Value == "" {
			return c.Redirect(http.StatusUnauthorized, "/login")
		}

		claim := &jwtClaims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claim, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || token == nil || !token.Valid {
			return c.Redirect(http.StatusUnauthorized, "/login")
		}
		claims, ok := token.Claims.(*jwtClaims)
		if !ok {
			return c.Redirect(http.StatusUnauthorized, "/login")
		}
		c.Set("id", claims.ID)
		return next(c)
	}
}
