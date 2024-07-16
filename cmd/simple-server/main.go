package main

import (
	"flag"

	"github.com/bppdanto-t/roundrobin/internal/app/simple-server/handler"
	"github.com/bppdanto-t/roundrobin/pkg/httpclient"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.String("port", "8100", "port for simple service (without colon, ex: 8100)")
	routerURL := flag.String("router", "http://localhost:8000", "base URL for router")
	flag.Parse()

	// Register
	_, err := httpclient.SendPostHTTPRequest("http://localhost:"+*port, *routerURL+"/register")
	if err != nil {
		panic(err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler.Hello)
	e.POST("/request", handler.SimpleRequest)
	e.POST("/config/delay", handler.SetDelay)
	e.POST("/config/maintenance", handler.SetMaintenance)

	// Start server
	e.Logger.Fatal(e.Start(":" + *port))
}
