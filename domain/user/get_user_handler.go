package user

import (
	exc "Chandara-Sin/supergo-api/exception"
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type getUserListFunc func(context.Context) ([]User, error)

func (fn getUserListFunc) GetUserList(ctx context.Context) ([]User, error) {
	return fn(ctx)
}

func GetUserListHandler(svc getUserListFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)

		usrList, err := svc.GetUserList(c.Request().Context())
		if err != nil {
			errRes := exc.ErrorRes{
				Code:    GetUserError,
				Message: err.Error(),
			}
			log.Error("error",
				zap.String("code", errRes.Code),
				zap.Error(err),
			)
			return c.JSON(http.StatusInternalServerError, errRes)
		}
		return c.JSON(http.StatusOK, usrList)
	}
}

type getUserFunc func(context.Context, string) (*User, error)

func (fn getUserFunc) GetUser(ctx context.Context, id string) (*User, error) {
	return fn(ctx, id)
}

func GetUserHandler(svc getUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)
		ID := c.Param("id")

		usr, err := svc.GetUser(c.Request().Context(), ID)
		if err != nil {
			errRes := exc.ErrorRes{
				Code:    GetUserListError,
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
