package getgeneralinfo

import (
	"github.com/shirou/gopsutil/v3/host"
)

type SystemInfo struct {
	Hostname       string
	RunningProcs   uint64
	OSName         string
	Platform       string
	PlatformFamily string
	PlatformVer    string
	KernalVersion  string
	KernelArch     string
	HostId         string
}

func (s *SystemInfo) GetInfo() error {
	hostInfo, err := host.Info()

	if err != nil {
		return err
	}

	s.Hostname = hostInfo.Hostname
	s.RunningProcs = hostInfo.Procs
	s.OSName = hostInfo.OS
	s.Platform = hostInfo.Platform
	s.PlatformFamily = hostInfo.PlatformFamily
	s.PlatformVer = hostInfo.PlatformVersion
	s.KernalVersion = hostInfo.KernelVersion
	s.KernelArch = hostInfo.KernelArch
	s.HostId = hostInfo.HostID

	return nil
}
