package main

import (
	"ProcessManager/agent/network"
	"ProcessManager/api"
	"ProcessManager/db"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	network_engine := network.CreateManager()
	network_engine.Listen(9000)

	con_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "BA", "hERpUzg_RL4_Ja-*", "127.0.0.1", 3306, "BAUser")

	database_engine, err := db.CreateManager("mysql", con_str)
	if err != nil {
		return
	}
	database_engine.Run(5)

	time.Sleep(time.Second * 5)
	store, _ := redis.NewStore(1000, "tcp", "localhost:6379", "", []byte("test"))

	r := gin.Default()
	r.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"http://localhost:3000"},
				AllowMethods:     []string{"POST", "GET"},
				MaxAge:           30 * time.Minute,
				AllowCredentials: true,
			}),
		sessions.Sessions("auth-token", store),
	)
	r.POST("/login", api.API_Login(database_engine))
	r.GET("/logout", api.API_Logout)

	agent := r.Group("/agent")
	agent.Use(Authendication())
	{
		agent.POST("/addnode", api.API_AddNode(database_engine))
		agent.GET("/getnodes", api.API_GetNodes(database_engine)) //이거 임시로 db조회로 만듬
		agent.POST("/executeprocess", api.API_ExecuteProcess(network_engine))
	}

	r.Run()
	network_engine.Stop()
}

func Authendication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("id")
		if id == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			c.Abort()
		}
		c.Next()
	}
}

//curl http://127.0.0.1:8080/agent/addnode --include --header 'Content-Type: application/json' --request 'POST' --data '{"ip" : "127.0.0.1",  "node":"test"}'
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
