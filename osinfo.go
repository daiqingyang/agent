package main

import (
	"os"
)

type Os struct {
	Hostname string
	Platform string
}

func ScanOs() (o *Os) {
	name, e := os.Hostname()
	if e != nil {
		logger.Println(e)
		return
	}
	o = &Os{
		name,
		"",
	}
	return
}
