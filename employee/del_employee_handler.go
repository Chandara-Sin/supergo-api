package employee

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type delEmployeeFunc func(context.Context, string) error

func (fn delEmployeeFunc) DelEmployee(ctx context.Context, id string) error {
	return fn(ctx, id)
}

func DelEmployeeHandler(svc delEmployeeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := svc.DelEmployee(c.Request().Context(), id)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "successfully delete employee with id : " + id,
		})
	}
}
