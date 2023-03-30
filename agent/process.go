package agent

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

type Process struct {
	pro    *exec.Cmd
	input  io.Writer
	output io.Reader
	ch     chan os.Signal
	done   chan bool
}

func CreateProcess(process string, arg ...string) *Process {
	p := new(Process)
	p.pro = exec.Command(process, arg...)

	var err error
	p.input, err = p.pro.StdinPipe()
	if err != nil {
		fmt.Print("fail")
	}
	p.output, err = p.pro.StdoutPipe()
	if err != nil {
		fmt.Print("fail")
	}

	p.ch = make(chan os.Signal, 1)
	p.done = make(chan bool, 1)

	return p
}

func (p *Process) Run(wg *sync.WaitGroup) {
	wg.Add(2)

	p.pro.Start()
	go func() {
		defer wg.Done()

		read_buf := bufio.NewReader(p.output)
		for {
			str, err := read_buf.ReadString('\n')
			if err != nil {
				p.done <- true
				return
			}

			fmt.Println(str)
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case sig := <-p.ch:
			p.pro.Process.Signal(sig)
		case <-p.done:
			return
		}
	}()
}

func (p *Process) Stop() {
	p.ch <- os.Interrupt
}

func (p *Process) ForceStop() {
	p.ch <- os.Kill
}

func (p *Process) Release() {
	defer close(p.ch)
	defer close(p.done)
}
