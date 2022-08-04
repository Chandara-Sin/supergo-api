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

	"Chandara-Sin/supergo-api/employee"
	"Chandara-Sin/supergo-api/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	// To configure auth via URI instead of a Credential, use
	// "mongodb://user:password@localhost:27017".
	credential := options.Credential{
		Username: viper.GetString("mongo.user"),
		Password: viper.GetString("mongo.password"),
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("mongo.uri")).SetAuth(credential))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	mongodb := client.Database(viper.GetString("mongo.db"))

	e := echo.New()
	// Default Level is Error
	e.Logger.SetLevel(log.INFO)

	// Log with each request
	e.Use(middleware.Logger())

	// Recover from panic
	e.Use(middleware.Recover())
	e.Use(logger.Middleware(zaplog))

	e.GET("/healths", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	em := e.Group("/v1/employees")
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
