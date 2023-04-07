package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Execute(c *gin.Context) {
	var data ExecuteData
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter."})
		return
	}

}
