package auth

import (
	b64 "encoding/base64"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AccessToken(signature string) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			Issuer:    "https://supergo-api.",
			Audience:  jwt.ClaimStrings{"supergo-api"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		})

		at, err := token.SignedString([]byte(signature))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
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
