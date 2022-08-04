package employee

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createEmployeeFunc func(context.Context, Employee) (primitive.ObjectID, error)

func (fn createEmployeeFunc) CreateEmployee(ctx context.Context, reqEmployee Employee) (primitive.ObjectID, error) {
	return fn(ctx, reqEmployee)
}

func CreateEmployeeHandler(svc createEmployeeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqEmployee Employee
		if err := c.Bind(&reqEmployee); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		newId, err := svc.CreateEmployee(c.Request().Context(), reqEmployee)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		resEmployee := Employee{
			Id:    newId,
			Name:  reqEmployee.Name,
			Email: reqEmployee.Email,
		}
		return c.JSON(http.StatusOK, resEmployee)
	}
}
