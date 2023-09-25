package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"Chandara-Sin/supergo-api/auth"
	"Chandara-Sin/supergo-api/config"
	"Chandara-Sin/supergo-api/domain/user"
	"Chandara-Sin/supergo-api/logger"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	initConfig()

	zaplog, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer zaplog.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := config.InitMongoDB(ctx)
	defer func() {
		if err := db.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mongodb := db.Client.Database(viper.GetString("mongo.db"))

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(logger.Middleware(zaplog))
	CORSConfig := middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	e.Use(middleware.CORSWithConfig(CORSConfig))

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "v1 OK")
	})

	e.GET("/v1/token", auth.AccessToken(viper.GetString("auth.sign")), middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:" + viper.GetString("api.key"),
		Validator: auth.ValidatorOnlyAPIKey(viper.GetString("api.public")),
	}))

	JWTConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SigningKey: []byte(viper.GetString("auth.sign")),
	}
	u := e.Group("/v1/users", echojwt.WithConfig(JWTConfig))
	u.POST("", user.CreateUserHandler(user.Create(mongodb)))
	u.GET("/:id", user.GetUserHandler(user.GetUser(mongodb)))
	u.GET("", user.GetUserListHandler(user.GetUserList(mongodb)))
	u.PUT("", user.UpdateUserHandler(user.Update(mongodb)))
	u.DELETE("/:id", user.DeleteUserHandler(user.DeleteUser(mongodb)))

	go func() {
		if err := e.Start(":" + viper.GetString("app.port")); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
