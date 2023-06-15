package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func requestCounter() gin.HandlerFunc {
	var (
		opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
			Name: "myapp_processed_ops_total",
			Help: "The total number of processed events",
		})
	)

	return func(c *gin.Context) {
		opsProcessed.Inc()
		c.Next()
	}
}

func main() {
	api_server := gin.Default()
	api_server.Use(requestCounter())
	api_server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	api_server.Run()
	// con := network.Create("client")
	// con.Connect("127.0.0.1:9000")

	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	var data []byte
	// 	for {
	// 		read, err := con.Read()
	// 		if err != nil {
	// 			break
	// 		}

	// 		data = append(data, read...)
	// 		buf := bytes.NewBuffer(data)
	// 		for {
	// 			var header packet.Header
	// 			err := binary.Read(buf, binary.BigEndian, &header)
	// 			if err != nil {
	// 				break
	// 			}

	// 			if buf.Len() < int(header.SIZE) {
	// 				break
	// 			}

	// 			s_data := make([]byte, header.SIZE)
	// 			if binary.Read(buf, binary.BigEndian, &s_data) != nil {
	// 				break
	// 			}

	// 			handler.Handle[header.ID](s_data)
	// 			data = data[len(data)-buf.Len():]
	// 		}
	// 	}
	// }()
	// wg.Wait()
	// Test()

	// ch := make(chan bool)

	// wg := &sync.WaitGroup{}

	// wg.Add(1)
	// go func() {
	// 	<-ch
	// 	wg.Done()
	// }()
	// wg.Wait()
}

func Test() {
	go func() {
		for {
			fmt.Print("test")
			time.Sleep(time.Second * 3)
		}
	}()
}
