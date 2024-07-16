package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bppdanto-t/roundrobin/internal/pkg/routing/router"
	"github.com/bppdanto-t/roundrobin/pkg/httpclient"
	"github.com/labstack/echo"
)

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func SimpleRequest(c echo.Context) error {
	if b, err := io.ReadAll(c.Request().Body); err == nil {
		var err error
		var res *http.Response
		for try := 0; try < 5; try++ {
			address := router.GetAddress()
			res, err = httpclient.SendPostHTTPRequest(string(b), address + "/request")
			if err != nil || res.StatusCode != http.StatusOK {
				fmt.Printf("retry request due to error; try: %d; address: %s\n", try, address)
				continue
			}
			break
		}
		if err != nil {
			return c.String(res.StatusCode, err.Error())
		}
		resStr, err := readResponse(res)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
    return c.String(res.StatusCode, resStr)
	} else {
		return c.String(http.StatusInternalServerError, err.Error())
	}
}

func Register(c echo.Context) error {
	if b, err := io.ReadAll(c.Request().Body); err == nil {
		router.Register(string(b))
		return c.String(http.StatusOK, fmt.Sprintf("Successfully register %s", string(b)))
	} else {
		return c.String(http.StatusInternalServerError, err.Error())
	}
}

func readResponse(res *http.Response) (string, error) {
	if res == nil {
		return "", errors.New("nil response")
	}
	if b, err := io.ReadAll(res.Body); err == nil {
		return string(b), nil
	} else {
		return "", err
	}
}