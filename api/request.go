package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data LoginData
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data.ID})
}

func Status(c *gin.Context) {

}

func Execute(c *gin.Context) {
	var data ExecuteData
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter."})
		return
	}

}
