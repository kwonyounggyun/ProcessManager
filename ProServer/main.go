package main

import (
	"ProcessManager/agent/network"
	"ProcessManager/agent/network/packet"
	"ProcessManager/api"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//var wg sync.WaitGroup

	// ch1 := Adding(&wg, "1")
	// ch2 := Adding(&wg, "2")
	// ch3 := Adding(&wg, "3")

	// time.Sleep(time.Second * 3)

	// (*ch1) <- false
	// (*ch2) <- false
	// (*ch3) <- true
	// time.Sleep(time.Second * 2)
	// close(*ch1)
	// close(*ch2)
	// p := agent.CreateProcess("D:\\Projects\\StudyLib\\CoreLib\\x64\\Debug\\BATestServerD")
	// //p := agent.CreateProcess("go", "env")
	// p.Run(&wg)
	// //time.Sleep(time.Second * 10)
	// p.Stop()

	//wg.Wait()

	manager := network.CreateManager()
	manager.Listen(9000)
	time.Sleep(time.Second * 5)
	store := cookie.NewStore([]byte("secret"))

	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))
	i := 1
	agent := r.Group("/agent")
	{
		agent.GET("/status", func(c *gin.Context) {
			session := sessions.Default(c)

			val := session.Get("id")
			if val == nil {
				session.Set("id", i)
				session.Save()
				i++
			}

			str := fmt.Sprintf("%d", session.Get("id"))
			c.JSON(200, gin.H{
				"message": str,
			})
		})

		agent.POST("/executeprocess", func(c *gin.Context) {
			var data api.ExecuteProcessData
			err := c.BindJSON(&data)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter."})
				return
			}

			pack := packet.ReqeustExecute{Path: "", Args: ""}
			pack2 := packet.MakePacket(packet.ReqeustExecuteID, &pack)

			manager.SendPacket(data.Node, pack2)

			c.JSON(http.StatusOK, gin.H{"message": "Execute"})
		})
	}

	user := r.Group("/user")
	{
		user.POST("/login", api.Login)
		//user.POST("/signin")
	}

	r.Run()
	manager.Stop()
}

// func Adding(wg *sync.WaitGroup, prin string) *chan bool {
// 	wg.Add(1)

// 	ch := make(chan bool)
// 	go func(ch chan bool) {
// 		run := true

// 		defer wg.Done()

// 		for run {
// 			select {
// 			case <-time.After(time.Second * 1):
// 				fmt.Println(prin)
// 			case stop := <-ch:
// 				run = stop
// 			}
// 		}
// 	}(ch)

// 	return &ch
// }
