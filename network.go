package main

import (
	"net"
	"strings"
)

type Network struct {
	Interfaces []MyInterface
}

type MyInterface struct {
	Name     string
	HardWare string
	Addrs    []string
}

func ScanNetwork() (nw *Network) {
	var interfaces []MyInterface
	inters, err := net.Interfaces()
	if err != nil {
		logger.Println(err)
		return
	}
	for _, inter := range inters {
		name := inter.Name
		if name == "lo" {
			continue
		}
		hw := inter.HardwareAddr.String()
		ads := []string{}
		addrs, e := inter.Addrs()
		if e != nil {
			logger.Println(e)
			return
		}
		for _, addr := range addrs {
			if strings.Contains(addr.String(), ":") {
				continue
			}
			ads = append(ads, addr.String())
		}
		my := MyInterface{
			name,
			hw,
			ads,
		}
		interfaces = append(interfaces, my)
	}
	nw = &Network{interfaces}
	return
}
