package user

import (
	exc "Chandara-Sin/supergo-api/exception"
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type createUserFunc func(context.Context, User) error

func (fn createUserFunc) CreateUser(ctx context.Context, reqUser User) error {
	return fn(ctx, reqUser)
}

func CreateUserHandler(svc createUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		usr := User{}
		log := logger.Unwrap(c)

		if err := c.Bind(&usr); err != nil {
			errRes := exc.ErrorRes{
				Code:    CreateUserError,
				Message: err.Error(),
			}
			log.Error("error",
				zap.String("code", errRes.Code),
				zap.Error(err),
			)
			return c.JSON(http.StatusBadRequest, errRes)
		}

		usr.CreatedAt = time.Now()
		usr.UpdatedAt = time.Now()

		err := svc.CreateUser(c.Request().Context(), usr)
		if err != nil {
			errRes := exc.ErrorRes{
				Code:    CreateUserError,
				Message: err.Error(),
			}
			log.Error("error",
				zap.String("code", errRes.Code),
				zap.Error(err),
			)
			return c.JSON(http.StatusInternalServerError, errRes)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "status ok",
		})
	}
}
