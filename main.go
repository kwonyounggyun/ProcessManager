package main

import (
	"ProcessManager/agent"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

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
	p := agent.CreateProcess("D:\\Projects\\StudyLib\\CoreLib\\x64\\Debug\\BATestServerD")
	//p := agent.CreateProcess("go", "env")
	p.Run(&wg)
	//time.Sleep(time.Second * 10)
	p.Stop()

	wg.Wait()
	//p.Release()
}

func Adding(wg *sync.WaitGroup, prin string) *chan bool {
	wg.Add(1)

	ch := make(chan bool)
	go func(ch chan bool) {
		run := true

		defer wg.Done()

		for run {
			select {
			case <-time.After(time.Second * 1):
				fmt.Println(prin)
			case stop := <-ch:
				run = stop
			}
		}
	}(ch)

	return &ch
}
