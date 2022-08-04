package employee

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type getEmployeeListFunc func(context.Context) ([]Employee, error)

func (fn getEmployeeListFunc) GetEmployeeList(ctx context.Context) ([]Employee, error) {
	return fn(ctx)
}

func GetEmployeeListHandler(svc getEmployeeListFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		empRes, err := svc.GetEmployeeList(c.Request().Context())
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, empRes)
	}
}

type getEmployeeFunc func(context.Context, string) (*Employee, error)

func (fn getEmployeeFunc) GetEmployee(ctx context.Context, id string) (*Employee, error) {
	return fn(ctx, id)
}

func GetEmployeeHandler(svc getEmployeeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		empRes, err := svc.GetEmployee(c.Request().Context(), id)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, empRes)
	}
}
