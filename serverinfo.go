package xgo

import (
    "fmt"
    "runtime"
    "time"
	"encoding/json"
	"strconv"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/load"
    "github.com/shirou/gopsutil/mem"
)

const (
    B  = 1
    KB = 1024 * B
    MB = 1024 * KB
    GB = 1024 * MB
)

type Mem struct{
	UsedMB int `json:"usedMB"`
	TotalMB int `json:"totalMB"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct{
	UsedMB int `json:"usedMB"`
	TotalMB int `json:"totalMB"`
	UsedPercent int `json:"usedPercent"`
}

type Osinfo struct{
	Ostype string `json:"os"`
	Compiler string `json:"compiler"`
	NumCpu int `json:"numCpu"`
	Goversion string `json:"goversion"`
	NumGoroutine int `json:"numGoroutine"`
}

type Cpu struct{
	UsedMB int `json:"usedMB"`
	TotalMB int `json:"totalMB"`
	UsedPercent int `json:"usedPercent"`
	Cores int `json:"cores"`
	Cpus int `json:"cpus"`
	Load1 float64 `json:"load1"`
	Load5 float64 `json:"load5"`
	Load15 float64 `json:"load15"`
	Cpulist []float64 `json:"list"`
}

type Jiankong struct{
	Cpu0 Cpu `json:"cpu"`
	Osinfo0 Osinfo `json:"osinfo"`
	Disk0 Disk `json:"disk"`
	Mem0 Mem `json:"mem"`
}

func Getsystem() string {
	// 磁盘
	d, _ := disk.Usage("/")
    // d_usedMB := int(d.Used) / MB
    d_usedGB := int(d.Used) / GB
    // d_totalMB := int(d.Total) / MB
    d_totalGB := int(d.Total) / GB
    d_usedPercent := int(d.UsedPercent)
	var di0 Disk = Disk{UsedMB:d_usedGB,TotalMB:d_totalGB,UsedPercent:d_usedPercent}

	var os0 Osinfo = Osinfo{Ostype:runtime.GOOS,Compiler:runtime.Compiler,NumCpu:runtime.NumCPU(),Goversion:runtime.Version(),NumGoroutine:runtime.NumGoroutine()}

	// CPU
	cores, _ := cpu.Counts(false)
    cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true)
	sli := []float64{}
    if err == nil {
        for _, c := range cpus {
			num,_ := strconv.ParseFloat(fmt.Sprintf("%.2f", c), 64)
			sli = append(sli, num)
        }
    }
    a, _ := load.Avg()
	var cpu0 Cpu = Cpu{UsedMB:20,TotalMB:21,UsedPercent:22,Cores:cores,Cpus:0,Load1:a.Load1,Load5:a.Load5,
		Load15:a.Load15,
		Cpulist:sli,
	}

	// 内存
	m, _ := mem.VirtualMemory()
    m_usedMB := int(m.Used) / MB
    m_totalMB := int(m.Total) / MB
    m_usedPercent := int(m.UsedPercent)
	var mem0 Mem = Mem{UsedMB:m_usedMB,TotalMB:m_totalMB,UsedPercent:m_usedPercent}

	var jian0 Jiankong = Jiankong{Disk0:di0,Osinfo0:os0,Cpu0:cpu0,Mem0:mem0}
	jian1, _ := json.Marshal(jian0)
	return string(jian1)
}