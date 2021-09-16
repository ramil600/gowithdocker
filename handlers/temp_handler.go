package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)


func TempHandler(c echo.Context) error {

	temp := c.Param("temp")
	rawuuid := c.Param("uuid")

	uuid, err := uuid.Parse(rawuuid)

	if err != nil{
		log.Warn("The Wrong Format for uuid: ", rawuuid )
	}

	return c.String(http.StatusOK, fmt.Sprintf("Temperature from %s was accepted: %s", uuid, temp))


}