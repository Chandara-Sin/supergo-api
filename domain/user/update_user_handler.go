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

type updateUserFunc func(context.Context, User) error

func (fn updateUserFunc) UpdateUser(ctx context.Context, usr User) error {
	return fn(ctx, usr)
}

func UpdateUserHandler(svc updateUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)
		usr := User{}

		if err := c.Bind(&usr); err != nil {
			errRes := exc.ErrorRes{
				Code:    UpdateUserError,
				Message: err.Error(),
			}
			log.Error("error",
				zap.String("code", errRes.Code),
				zap.Error(err),
			)
			return c.JSON(http.StatusBadRequest, errRes)
		}

		usr.UpdatedAt = time.Now()

		err := svc.UpdateUser(c.Request().Context(), usr)
		if err != nil {
			errRes := exc.ErrorRes{
				Code:    UpdateUserError,
				Message: err.Error(),
			}
			log.Error("error",
				zap.String("code", errRes.Code),
				zap.Error(err),
			)
			return c.JSON(http.StatusInternalServerError, errRes)
		}

		return c.JSON(http.StatusOK, usr)
	}
}
