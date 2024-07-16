package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bppdanto-t/roundrobin/internal/app/simple-server/configs"
	"github.com/labstack/echo"
)

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func SimpleRequest(c echo.Context) error {
	if configs.IsMaintenance() {
		return c.String(http.StatusServiceUnavailable, "Service in maintenance")
	}

	configs.Delay()

	if b, err := io.ReadAll(c.Request().Body); err == nil {
    return c.String(http.StatusOK, string(b))
	} else {
		return c.String(http.StatusInternalServerError, err.Error())
	}
}

func SetDelay(c echo.Context) error {
	delayForm := c.FormValue("delay")
	delay, err := strconv.ParseInt(delayForm, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "incorrect delay value")
	}
	configs.SetDelay(delay)
	return c.String(http.StatusOK, fmt.Sprintf("Delay has been set to %d", delay))
}

func SetMaintenance(c echo.Context) error {
	maintenanceForm := c.FormValue("maintenance")
	isMaintenance, err := strconv.ParseBool(maintenanceForm)
	if err != nil {
		return c.String(http.StatusBadRequest, "incorrect maintenance value")
	}
	configs.SetMaintenance(isMaintenance)
	return c.String(http.StatusOK, fmt.Sprintf("Maintenance has been set to %t", isMaintenance))
}