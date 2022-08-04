package employee

import (
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createEmployeeFunc func(context.Context, Employee) (*primitive.ObjectID, error)

func (fn createEmployeeFunc) CreateEmployee(ctx context.Context, reqEmployee Employee) (*primitive.ObjectID, error) {
	return fn(ctx, reqEmployee)
}

func CreateEmployeeHandler(svc createEmployeeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqEmp Employee
		log := logger.Unwrap(c)

		if err := c.Bind(&reqEmp); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		resId, err := svc.CreateEmployee(c.Request().Context(), reqEmp)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		resEmp := Employee{
			ID:    *resId,
			Name:  reqEmp.Name,
			Email: reqEmp.Email,
		}
		return c.JSON(http.StatusOK, resEmp)
	}
}
