package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)


func PingHandler(c echo.Context) error {
	fmt.Println("Hello There!")
	return c.String(http.StatusOK, "Pong!")
}