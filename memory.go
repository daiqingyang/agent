package main

import (
	"github.com/shirou/gopsutil/v3/mem"
)

func getMem() (total uint64, used uint64, free uint64) {

	vmStat, e := mem.VirtualMemory()
	if e != nil {
		logger.Println(e)
		return

	}
	total = vmStat.Total
	used = vmStat.Used
	free = vmStat.Free
	return

}

func getSwap() (total uint64, used uint64, free uint64) {
	SwapStat, e := mem.SwapMemory()
	if e != nil {
		logger.Println(e)
		return
	}
	total = SwapStat.Total
	used = SwapStat.Used
	free = SwapStat.Free
	return
}
