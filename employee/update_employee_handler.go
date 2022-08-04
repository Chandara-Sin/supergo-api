package employee

import (
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type updateEmployeeFunc func(context.Context, Employee) error

func (fn updateEmployeeFunc) UpdateEmployee(ctx context.Context, reqEmployee Employee) error {
	return fn(ctx, reqEmployee)
}

func UpdateEmployeeHandler(svc updateEmployeeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqEmp Employee
		log := logger.Unwrap(c)

		if err := c.Bind(&reqEmp); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		err := svc.UpdateEmployee(c.Request().Context(), reqEmp)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, reqEmp)
	}
}
