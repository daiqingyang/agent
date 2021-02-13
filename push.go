package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func SendPing() {
	serverAdress := "localhost:10240"
	url := "/ping"
	interval := 3 //second
	if !strings.HasPrefix(serverAdress, "http://") || !strings.HasPrefix(serverAdress, "https://") {
		serverAdress = "http://" + serverAdress
	}
	for {
		resp, err := http.Get(serverAdress + url)
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println(resp.StatusCode, time.Now())
		}
		time.Sleep(time.Second * time.Duration(interval))

	}
}
