package controllers

import (
	"apscsc-backend/responces"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GunnyBagCount(c echo.Context) error {
	// Get the API key from environment variables
	apiKey := os.Getenv("APIkey")
	if apiKey == "" {
		logrus.Error("API key is empty or not found in .env")
		return c.JSON(http.StatusInternalServerError, responces.GunnyBagResponce{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: "API key is not set",
			Data:    nil,
		})
	}

	// Return the gunny bag count in the response
	return c.JSON(http.StatusOK, responces.GunnyBagResponce{
		Status:  http.StatusOK,
		Success: true,

		Message: "Gunny bag count retrieved successfully",
		Data:    nil,
	})
}

// CreateGunnyBag handles the creation of a new gunny bag
