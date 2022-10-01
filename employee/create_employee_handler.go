package employee

import (
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createEmployeeFunc func(context.Context, Employee) (*Employee, error)

func (fn createEmployeeFunc) CreateEmployee(ctx context.Context, reqEmployee Employee) (*Employee, error) {
	return fn(ctx, reqEmployee)
}

func CreateEmployeeHandler(svc createEmployeeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqEmp Employee
		log := logger.Unwrap(c)

		if err := c.Bind(&reqEmp); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		}

		resEmp, err := svc.CreateEmployee(c.Request().Context(), reqEmp)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, resEmp)
	}
}
