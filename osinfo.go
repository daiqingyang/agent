package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

type Os struct {
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
}

// CpuModeName        string

func ScanOs() (o *Os) {
	hInfo := getHostInfo()
	PhysicalCpuNum := getPhysicalCpuNum()
	LogicalCpuNum := getLogicalCpuNum()
	CpuModelName := getCpuModelName()
	o = &Os{
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
	}
	return
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
