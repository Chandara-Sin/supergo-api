package employee

import (
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type delEmployeeFunc func(context.Context, string) error

func (fn delEmployeeFunc) DelEmployee(ctx context.Context, id string) error {
	return fn(ctx, id)
}

func DelEmployeeHandler(svc delEmployeeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		log := logger.Unwrap(c)

		err := svc.DelEmployee(c.Request().Context(), id)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "successfully delete employee with id : " + id,
		})
	}
}
