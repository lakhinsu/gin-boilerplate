package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lakhinsu/gin-boilerplate/models"
	"github.com/rs/zerolog/log"
)

func Hello(c *gin.Context) {
	var user models.User

	request_id := c.GetString("x-request-id")

	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(&user); binderr != nil {

		log.Error().Err(binderr).Str("request_id", request_id).
			Msg("Error occurred while binding request data")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello %s %s", user.FirstName, user.LastName),
	})

}
