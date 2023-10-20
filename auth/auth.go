package auth

import (
	exc "Chandara-Sin/supergo-api/exception"
	"Chandara-Sin/supergo-api/logger"
	b64 "encoding/base64"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const (
	TokenizeError = "SUP-AUTH-40001"
	AuthError     = "SUP-AUTH-40002"
)

func AccessToken(signature string) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			Issuer:    "https://supergo-api.",
			Audience:  jwt.ClaimStrings{"supergo-api"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		})

		at, err := token.SignedString([]byte(signature))
		if err != nil {
			errRes := exc.ErrorRes{
				Code:    TokenizeError,
				Message: err.Error(),
			}
			log.Error("error",
				zap.String("code", TokenizeError),
				zap.Error(err),
			)
			return c.JSON(http.StatusInternalServerError, errRes)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token":      at,
			"token_type": "Bearer",
		})
	}
}

func ValidatorOnlyAPIKey(apikey string) middleware.KeyAuthValidator {
	return func(apiKeyHeader string, c echo.Context) (bool, error) {
		apiKeyEnc := b64.StdEncoding.EncodeToString([]byte(apikey))
		return apiKeyHeader == apiKeyEnc, nil
	}
}
