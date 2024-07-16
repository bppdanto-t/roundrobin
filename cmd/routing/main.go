package main

import (
	"flag"

	"github.com/bppdanto-t/roundrobin/internal/app/routing/handler"
	"github.com/bppdanto-t/roundrobin/internal/pkg/routing/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.String("port", "8000", "port for simple service (without colon, ex: 8000)")
	flag.Parse()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler.Hello)
	e.POST("/register", handler.Register)
	e.POST("/request", handler.SimpleRequest)

	router.Init()

	// Start server
	e.Logger.Fatal(e.Start(":" + *port))
}