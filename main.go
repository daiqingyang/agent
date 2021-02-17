package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	test()
	os.Exit(1)
	go SendPing()
	StartHttp()
}
func test() {
	r := strings.NewReader("helloworld,helloworld,helloworld,")
	var b []byte = make([]byte, 8, 8)
	for {
		n, e := r.Read(b)
		if e != nil {
			fmt.Println("err:", e)
			if e == io.EOF {
				break
			}
		}

		i, _ := r.Seek(0, os.SEEK_CUR)
		fmt.Println(n, i, string(b))
	}
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
