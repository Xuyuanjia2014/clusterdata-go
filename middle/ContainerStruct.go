package middle

type AppDu struct {
	Containers map[string]Container `yaml:"containers,flow"`
}

type Container struct {
	Id string `yaml:"id,flow"`
	MachineId string `yaml:"mid,flow"`
	TimeStamp int `yaml:"ts,flow"`
	CpuRequest int `yaml:"cr,flow"`
	CpuLimit int `yaml:"cl,flow"`
	MemSize float64 `yaml:"ms,flow"`
	Status string `yaml:"sta,flow"`
	Usages []CUsage `yaml:"Usages,flow"`
}

type CUsage struct {
	MachineId string `yaml:"mid,flow"`
	TimeStamp int `yaml:"ts,flow"`
	CpuPercent int `yaml:"cp,flow"`
	Mpki int `yaml:"mpki,flow"`
	Cpi float64 `yaml:"cpi,flow"`
	MemPercent int `yaml:"mp,flow"`
	MemGps float64 `yaml:"mg,flow"`
	DiskIoPercent float64 `yaml:"dip,flow"`
	NetIn float64 `yaml:"ni,flow"`
	NetOut float64 `yaml:"no,flow"`
}


