package user

import (
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type updateUserFunc func(context.Context, User) error

func (fn updateUserFunc) UpdateUser(ctx context.Context, usr User) error {
	return fn(ctx, usr)
}

func UpdateUserHandler(svc updateUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log, _ := logger.Unwrap(c)
		usr := User{}

		if err := c.Bind(&usr); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		}

		usr.UpdatedAt = time.Now()

		err := svc.UpdateUser(c.Request().Context(), usr)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, usr)
	}
}
