package employee

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type getEmployeeFunc func(context.Context, string) (*Employee, error)

func (fn getEmployeeFunc) GetEmployee(ctx context.Context, Id string) (*Employee, error) {
	return fn(ctx, Id)
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
