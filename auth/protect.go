package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Protect(signature []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			tokenString := strings.TrimPrefix(auth, "Bearer ")

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return signature, nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"status": "unauthorized",
					"error":  err.Error(),
				})
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				aud := claims["aud"]
				c.Set("aud", aud)
			}
			return next(c)
		}
	}
}
