package main

import (
	"fmt"
	"net/http"
	"time"
)

func SendPing() {
	serverAddres := "http://www.baidu.com"
	url := "/ping"
	interval := 3 //second
	for {
		resp, _ := http.Get(serverAddres + url)
		fmt.Println(resp.StatusCode, time.Now())
		time.Sleep(time.Second * time.Duration(interval))

	}
}
