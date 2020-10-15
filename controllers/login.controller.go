package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hafif/echoFramework/helpers"
	"github.com/hafif/echoFramework/models"
	"github.com/labstack/echo"
)

func RegisterUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	hash, err := helpers.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := models.RegisterUser(username, hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, res)
}

func LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username

	if username == "admin" {
		claims["admin"] = true
	}

	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
