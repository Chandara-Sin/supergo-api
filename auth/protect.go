package auth

import (
	exc "Chandara-Sin/supergo-api/exception"
	"Chandara-Sin/supergo-api/logger"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Protect(signature []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := logger.Unwrap(c)

			auth := c.Request().Header.Get("Authorization")
			tokenString := strings.TrimPrefix(auth, "Bearer ")

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return signature, nil
			})
			if err != nil {
				errRes := exc.ErrorRes{
					Code:    AuthError,
					Message: err.Error(),
				}
				log.Error("error",
					zap.String("code", AuthError),
					zap.Error(err),
				)
				return c.JSON(http.StatusUnauthorized, errRes)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				aud := claims["aud"]
				c.Set("aud", aud)
			}
			return next(c)
		}
	}
}
