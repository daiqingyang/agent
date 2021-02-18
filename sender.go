package main

import (
	"net/http"
	"strings"
	"time"
)

func SendPing() {
	serverAdress := config.ServerAdress
	url := config.PingUrl
	interval := config.PingInterval //second
	//min 5 second
	if interval < 5 {
		interval = 5
	}
	if !strings.HasPrefix(serverAdress, "http://") || !strings.HasPrefix(serverAdress, "https://") {
		serverAdress = "http://" + serverAdress
	}
	for {
		resp, err := http.Get(serverAdress + url)
		if err != nil {
			logger.Println(err)

		} else {
			logger.Println(resp.StatusCode, time.Now())
		}
		time.Sleep(time.Second * time.Duration(interval))

	}
}
