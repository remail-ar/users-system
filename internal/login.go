package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func saveUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	log.Printf("POST /login - Username: %s ", username)
	log.Printf("POST /login - Password: %s ", password)
	return c.String(http.StatusOK, "Recieved")
}
