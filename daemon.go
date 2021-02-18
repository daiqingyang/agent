package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runSubProcessCheck(daemon bool) {
	if daemon {
		mainProcess := os.Args[0]
		args := os.Args[1:]
		var newArgs []string
		//子进程去除-d参数
		for _, i := range args {
			if i != "-d" {
				newArgs = append(newArgs, i)
			}
		}
		//分发子进程，退出主进程
		cmd := exec.Command(mainProcess, newArgs...)
		if e := cmd.Start(); e != nil {
			fmt.Println(e)
		}
		os.Exit(0)

	}
}
