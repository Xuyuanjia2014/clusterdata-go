package middle

type Container struct {
	Id string
	MachineId string
	App string
	TimeStamp int
	CpuRequest int
	CpuLimit int
	MemSize float64
	Status string
	Usages []CUsage `yaml:"Usages,flow"`
}

type CUsage struct {
	MachineId string
	TimeStamp int
	CpuPercent int
	Mpki int
	Cpi float64
	MemPercent int
	MemGps float64
	DiskIoPercent float64
	NetIn float64
	NetOut float64
}
