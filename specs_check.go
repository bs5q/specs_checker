package main

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	// host
	hostInfo, _ := host.Info()
	fmt.Printf("OS: %s %s (%s)\n", hostInfo.OS, hostInfo.Platform, hostInfo.PlatformFamily)

	// ram info
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("Memory: %.2f GB\n", float64(memInfo.Total)/1024/1024/1024)

	// cpu info
	cpuInfo, _ := cpu.Info()
	fmt.Printf("CPU: %s\n", cpuInfo[0].ModelName)

	// hdd/ssd/m.2 info
	diskInfo, _ := disk.Usage("/")
	fmt.Printf("Disk: %.2f GB free out of %.2f GB\n", float64(diskInfo.Free)/1024/1024/1024, float64(diskInfo.Total)/1024/1024/1024)
}
