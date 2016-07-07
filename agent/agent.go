package agent

import (
	"github.com/virgild/sixalert/alert"
	"log"
	"time"
)

type Agent struct {
	checkIntervalSeconds int
	uptimeSeconds        int
	doneCh               chan bool
	alertsCh             <-chan *alert.TTCAlert
}

func NewAgent(alertsCh <-chan *alert.TTCAlert) *Agent {
	doneCh := make(chan bool)

	agent := &Agent{
		10,
		0,
		doneCh,
		alertsCh,
	}
	return agent
}

func (a *Agent) Start() {
	log.Printf("Agent starting")
	stillRunning := true
	for stillRunning {
		select {
		case <-time.After(time.Second * 1):
			a.uptimeSeconds++
			if a.uptimeSeconds%a.checkIntervalSeconds == 0 {
				a.runTask()
			}
		case <-a.doneCh:
			stillRunning = false
		}
	}

	log.Printf("Agent stopping")
}

func (a *Agent) Stop() {
	a.doneCh <- true
}

func (a *Agent) CheckIntervalSeconds() int {
	return a.checkIntervalSeconds
}

func (a *Agent) UptimeSeconds() int {
	return a.uptimeSeconds
}

func (a *Agent) runTask() {
	log.Printf("Agent task run %d", a.UptimeSeconds())
}
