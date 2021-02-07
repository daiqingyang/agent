package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Agent struct {
	*Network
	*Os
	*Disks
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
	Pprint(agent)
}
func NewAgent() *Agent {
	nw := ScanNetwork()
	os := ScanOs()
	disk := ScanDisks()
	return &Agent{
		nw,
		os,
		disk,
	}
}
func Pprint(agent *Agent) {
	bs, e := json.Marshal(agent)
	if e != nil {
		logger.Println(e)
	}
	var buf bytes.Buffer
	json.Indent(&buf, bs, "", "\t")
	fmt.Printf("%+v", buf.String())

}
