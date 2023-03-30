package agent

import "github.com/gin-gonic/gin"

const serviceName = "PMS Agent"
const serviceDescription = "Simple service, just for fun"

func RegistService() {
	r := gin.Default()
	Agent := r.Group("agent")
	{
		Agent.POST("/execute", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Gin success",
			})
		})
	}
	r.Run()
}
