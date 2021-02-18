package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func configParse() {
	var daemon bool
	var configFile string
	var home string
	var e error
	if home, e = os.UserHomeDir(); e != nil {
		home = "/root"
	}
	defaultConfigFile := home + "/.gagent.cfg"
	flag.StringVar(&configFile, "f", defaultConfigFile, "define agent config file path")
	flag.BoolVar(&debug, "v", false, "open verbose mode")
	flag.BoolVar(&daemon, "d", false, "run it in backgroud")
	flag.Parse()
	runSubProcessCheck(daemon)
	if debug {
		fmt.Println("$HOME:", home)
		fmt.Println(configFile)
	}
	if configFile == defaultConfigFile {
		if _, e := os.Stat(configFile); e != nil {
			if os.IsNotExist(e) {
				if f, e := os.Create(defaultConfigFile); e != nil {
					fmt.Println(e)
				} else {
					toml.NewEncoder(f).Encode(&config)
				}
			}
		}

	}
	if _, e := toml.DecodeFile(configFile, &config); e != nil {
		fmt.Println(e)
		os.Exit(2)
	}
	if debug {
		fmt.Println(config)
	}
}
