package main

import (
	"log"
	"os"
)

type Agent struct {
	*Network
	*Os
}

var (
	agent  *Agent
	logger *log.Logger
)

func init() {

	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)
}
func main() {
	agent = NewAgent()
	logger.Println(agent.Os)
}
func NewAgent() *Agent {
	nw := ScanNetwork()
	os := ScanOs()
	return &Agent{
		nw,
		os,
	}
}
