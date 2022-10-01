package user

import (
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type deleteUserFunc func(context.Context, string) error

func (fn deleteUserFunc) DeleteUser(ctx context.Context, id string) error {
	return fn(ctx, id)
}

func DeleteUserHandler(svc deleteUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)
		id := c.Param("id")

		err := svc.DeleteUser(c.Request().Context(), id)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "successfully delete employee with id: " + id,
		})
	}
}
