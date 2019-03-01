package middle

type Machine struct {
	Id string
	TimeStamp int
	Level1 int
	Level2 int
	CpuNum int
	MemSize int
	Status string
	Usages []CUsage `yaml:"Usages,flow"`
}

type MUsage struct {
	TimeStamp int
	CpuPercent int
	MemPercent int
	MemGps float64
	Mpki int
	DiskIoPercent float64
	NetIn float64
	NetOut float64
}