package user

import (
	exc "Chandara-Sin/supergo-api/exception"
	"Chandara-Sin/supergo-api/logger"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type deleteUserFunc func(context.Context, string) error

func (fn deleteUserFunc) DeleteUser(ctx context.Context, id string) error {
	return fn(ctx, id)
}

func DeleteUserHandler(svc deleteUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)
		ID := c.Param("id")

		err := svc.DeleteUser(c.Request().Context(), ID)
		if err != nil {
			errRes := exc.ErrorRes{
				Code:    DeleteUserError,
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
