package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExecuteProcess(c *gin.Context) {
	var data ExecuteProcessData
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter."})
		return
	}
}
