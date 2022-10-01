package user

import (
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getUserListFunc func(context.Context) ([]User, error)

func (fn getUserListFunc) GetUserList(ctx context.Context) ([]User, error) {
	return fn(ctx)
}

func GetUserListHandler(svc getUserListFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)

		rs, err := svc.GetUserList(c.Request().Context())
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, rs)
	}
}

type getUserFunc func(context.Context, string) (*User, error)

func (fn getUserFunc) GetUser(ctx context.Context, id string) (*User, error) {
	return fn(ctx, id)
}

func GetUserHandler(svc getUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)
		id := c.Param("id")

		rs, err := svc.GetUser(c.Request().Context(), id)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, rs)
	}
}
