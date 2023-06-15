package api

import (
	"net/http"

	"ProcessManager/db"
	"ProcessManager/db/dbtask"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func API_Login(db *db.DBManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data LoginData
		err := c.BindJSON(&data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": API_ERR_INVALID_PARAM.string()})
			return
		}

		ch := make(chan bool)
		task := dbtask.LoginTask{ID: data.ID, PW: data.PW, CallFunc: func() {
			ch <- true
		}}

		db.PushTask(&task)
		<-ch

		if task.KEY == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "check you are id or pw"})
			return
		}

		session := sessions.Default(c)
		session.Options(sessions.Options{MaxAge: 60 * 30, 
		})
		session.Set("id", data.ID)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"result": session.ID()})
	}
}

func API_Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}

func API_GetNodes(db *db.DBManager) gin.HandlerFunc {
	return func(c *gin.Context) {

		ch := make(chan bool)
		task := dbtask.GetNodesTask{CallFunc: func() {
			ch <- true
		}}
		db.PushTask(&task)

		<-ch

		c.JSON(http.StatusOK, task.Nodes)
	}
}

func API_AddNode(db *db.DBManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data AddNodeData
		err := c.BindJSON(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": API_ERR_INVALID_PARAM.string()})
			return
		}

		ch := make(chan bool)
		task := dbtask.AddNodeTask{IP: data.IP, NodeName: data.Node, CallFunc: func() {
			ch <- true
		}}

		if db == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": API_ERR_DATABASE_DISCONNECT.string()})
			return
		}

		db.PushTask(&task)
		<-ch

		if task.ID != -1 {
			c.JSON(http.StatusOK, gin.H{"ip": task.IP, "name": task.NodeName, "id": task.ID})
			return
		}
		c.JSON(http.StatusExpectationFailed, gin.H{"ip": data.IP, "name": data.Node})
	}
}

func Status(c *gin.Context) {

}
