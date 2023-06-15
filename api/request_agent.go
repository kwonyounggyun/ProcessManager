package api

import (
	"ProcessManager/agent/network"
	"ProcessManager/agent/network/packet"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API_GetStatus() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func API_ExecuteProcess(network *network.NetManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ExecuteProcessData
		err := c.BindJSON(&data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter."})
			return
		}

		pack := packet.ReqeustExecute{Path: "", Args: ""}
		pack2 := packet.MakePacket(packet.ReqeustExecuteID, &pack)

		network.SendPacket(data.Node, pack2)

		c.JSON(http.StatusOK, gin.H{"message": "Execute"})
	}
}

func API_StopProcess(c *gin.Context) {
	var data StopProcessData
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter."})
	}
}
