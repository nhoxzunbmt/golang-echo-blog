package main

import (
	"github.com/xesina/golang-echo-realworld-example-app/db"
	"github.com/xesina/golang-echo-realworld-example-app/handler"
	"github.com/xesina/golang-echo-realworld-example-app/router"
	"github.com/xesina/golang-echo-realworld-example-app/store"

	echoSwagger "github.com/swaggo/echo-swagger"                 // echo-swagger middleware
	_ "github.com/xesina/golang-echo-realworld-example-app/docs" // docs is generated by Swag CLI, you have to import it.

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title Swagger Example API
// @version 1.0
// @description Conduit API
// @title Conduit API

// @host 127.0.0.1:8585
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := router.New()

	r.Use(echoPrometheus.MetricsMiddleware())
	r.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewArticleStore(d)
	h := handler.NewHandler(us, as)
	h.Register(v1)
	r.Logger.Fatal(r.Start(":9001"))
}
