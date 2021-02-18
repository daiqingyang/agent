package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	LogFile      string
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
	configParse()
	CreateLog()
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
func CreateLog() {
	var out io.Writer
	if config.LogFile != "" {
		oFile, e := os.OpenFile(config.LogFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if e != nil {
			fmt.Println(e)
		}
		out = io.MultiWriter(oFile, os.Stdout)
		// out = oFile
	} else {
		out = os.Stdout
	}
	logger = log.New(out, "", log.Ldate|log.Ltime|log.Llongfile)

}
