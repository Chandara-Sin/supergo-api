package user

import (
	exc "Chandara-Sin/supergo-api/exception"
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type createUserFunc func(context.Context, User) error

func (fn createUserFunc) CreateUser(ctx context.Context, reqUser User) error {
	return fn(ctx, reqUser)
}

func CreateUserHandler(svc createUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		usr := User{}
		log := logger.Unwrap(c)

		// Todo
		// Get trace_id from context
		if err := c.Bind(&usr); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, exc.ErrorRes{
				Code:    CreateUserError,
				TraceId: "trace_id",
				Message: err.Error(),
			})
		}

		usr.CreatedAt = time.Now()
		usr.UpdatedAt = time.Now()

		err := svc.CreateUser(c.Request().Context(), usr)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "status ok",
		})
	}
}
