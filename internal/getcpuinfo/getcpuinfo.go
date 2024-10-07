package getcpuinfo

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUInfo struct {
	VendorID  string
	Family    string
	Cores     int32
	ModelName string
	Freq      float64
}

func (c *CPUInfo) GetInfo() error {
	cpuInfo, err := cpu.Info()

	if err != nil {
		return err
	}

	c.VendorID = cpuInfo[0].VendorID
	c.Family = cpuInfo[0].Family
	c.Cores = cpuInfo[0].Cores
	c.ModelName = cpuInfo[0].ModelName
	c.Freq = cpuInfo[0].Mhz

	return nil
}
