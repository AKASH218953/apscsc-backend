package controllers

import (
	"apscsc-backend/models"
	"apscsc-backend/responces"
	"net/http"

	"github.com/labstack/echo"
)

// FaceRecognitionHandler handles face recognition requests
func FaceRecognitionHandler(c echo.Context) error {
	// Parse the request body
	var request models.Face
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, responces.FaceResponce{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: "Invalid request format",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responces.FaceResponce{
		Status:  http.StatusOK,
		Success: true,
		Message: "Face recognition successful",
		Data:    request,
	})
}
