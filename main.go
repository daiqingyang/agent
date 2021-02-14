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
type Config struct {
	ServerAdress string
	PingUrl      string
	PingInterval int
}

var (
	agent  *Agent
	logger *log.Logger
	debug  bool
	config Config = Config{
		ServerAdress: "localhost:10240",
		PingUrl:      "/ping",
		PingInterval: 6,
	}
)

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)
	configParse()
}
func main() {
	go SendPing()
	StartHttp()
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
