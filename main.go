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
	"Chandara-Sin/supergo-api/employee"
	"Chandara-Sin/supergo-api/logger"

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
	
	// config
	config := middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	e.Use(middleware.CORSWithConfig(config))

	e.GET("/healths", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/token", auth.AccessToken(viper.GetString("auth.sign")))

	em := e.Group("/v1/employees", auth.Protect([]byte(viper.GetString("auth.sign"))))
	em.POST("", employee.CreateEmployeeHandler(employee.Create(mongodb)))
	em.GET("/:id", employee.GetEmployeeHandler(employee.GetEmployee(mongodb)))
	em.GET("", employee.GetEmployeeListHandler(employee.GetEmployeeList(mongodb)))
	em.PUT("", employee.UpdateEmployeeHandler(employee.Update(mongodb)))
	em.DELETE("/:id", employee.DelEmployeeHandler(employee.DelEmployee(mongodb)))

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
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err) // Handle errors reading the config file
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
