package main

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type Disks struct {
	DisksStatus []*disk.UsageStat
}

func ScanDisks() *Disks {
	var dsList []disk.PartitionStat
	dsList, e := disk.Partitions(false)
	if e != nil {
		logger.Println(e)
	}
	device := []string{}
	DisksStatus := []*disk.UsageStat{}
	for _, ds := range dsList {
		//排除/var/lib/docker/containers等
		if Contains(device, ds.Device) {
			continue
		}
		device = append(device, ds.Device)
		//UsageStat
		usageStat, e := disk.Usage(ds.Mountpoint)
		if e != nil {
			logger.Println(e)
		}
		DisksStatus = append(DisksStatus, usageStat)

	}
	disks := &Disks{
		DisksStatus,
	}
	return disks

}
func Contains(sl []string, s string) (contained bool) {
	contained = false
	for _, v := range sl {
		if v == s {
			contained = true
			break
		}
	}
	return
}
