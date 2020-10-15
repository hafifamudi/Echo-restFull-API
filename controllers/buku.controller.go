package controllers

import (
	"net/http"
	"strconv"

	"github.com/hafif/echoFramework/models"
	"github.com/labstack/echo"
)

func FetchAllData(c echo.Context) error {
	result, err := models.FetchAllData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreData(c echo.Context) error {
	nama := c.FormValue("nama")
	harga := c.FormValue("harga")
	jenis := c.FormValue("jenis")

	result, err := models.StoreData(nama, harga, jenis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateData(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	harga := c.FormValue("harga")
	jenis := c.FormValue("jenis")

	convID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateData(convID, nama, harga, jenis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteData(c echo.Context) error {
	id := c.FormValue("id")

	convID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteData(convID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
