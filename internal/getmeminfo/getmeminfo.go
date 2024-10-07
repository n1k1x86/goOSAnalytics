package getmeminfo

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemInfo struct {
	Total       uint64
	Free        uint64
	UsedPercent float64
}

func (m *MemInfo) GetInfo() error {
	v, err := mem.VirtualMemory()

	if err != nil {
		return err
	}

	m.Total = v.Total / 1024 / 1024
	m.Free = v.Free / 1024 / 1024
	m.UsedPercent = v.UsedPercent

	return nil
}
