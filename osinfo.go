package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
)

type Os struct {
	*load.AvgStat
	Hostname           string
	Uptime             uint64
	Procs              uint64
	OS                 string
	Platform           string
	PlatrformVersion   string
	KernalVersion      string
	VirtualizationRole string
	HostID             string //uuid
	PhysLogicalCpuNum  int
	LogicalCpuNum      int
	CpuModelName       string
	MemTotal           uint64
	MemUsed            uint64
	MemFree            uint64
	SwapTotal          uint64
	SwapUsed           uint64
	SwapFree           uint64
}

func ScanOs() (o *Os) {
	hInfo := getHostInfo()
	PhysicalCpuNum := getPhysicalCpuNum()
	LogicalCpuNum := getLogicalCpuNum()
	CpuModelName := getCpuModelName()
	avgStat := getLoad()
	MemTotal, MemUsed, MemFree := getMem()
	SwapTotal, SwapUsed, SwapFree := getSwap()

	o = &Os{
		avgStat,
		hInfo.Hostname,
		hInfo.Uptime,
		hInfo.Procs,
		hInfo.OS,
		hInfo.Platform,
		hInfo.PlatformVersion,
		hInfo.KernelVersion,
		hInfo.VirtualizationRole,
		hInfo.HostID,
		PhysicalCpuNum,
		LogicalCpuNum,
		CpuModelName,
		MemTotal,
		MemUsed,
		MemFree,
		SwapTotal,
		SwapUsed,
		SwapFree,
	}
	return

}
func getLoad() *load.AvgStat {
	avgStat, e := load.Avg()
	if e != nil {
		logger.Println(e)
	}
	return avgStat
}
func getHostInfo() *host.InfoStat {
	hInfo := &host.InfoStat{}
	hInfo, e := host.Info()
	if e != nil {
		logger.Println(e)
	}
	return hInfo
}
func getPhysicalCpuNum() int {
	var PhysicalCpuNum int
	PhysicalCpuNum, e := cpu.Counts(false)
	if e != nil {
		logger.Println(e)
	}
	return PhysicalCpuNum
}

func getLogicalCpuNum() int {
	var LogicalCpuNum int
	LogicalCpuNum, e := cpu.Counts(true)
	if e != nil {
		logger.Println(e)
	}
	return LogicalCpuNum

}
func getCpuModelName() string {
	var cInfo []cpu.InfoStat
	cInfo, e := cpu.Info()
	if e != nil {
		logger.Println(e)
	}
	return cInfo[0].ModelName

}
