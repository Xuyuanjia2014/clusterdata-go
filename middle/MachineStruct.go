package middle

import (
	strings "strings"
	"strconv"
	"gopkg.in/yaml.v2"
	"log"
)

type Machine struct {
	TimeStamp int `yaml:"ts"`
	Level1 int `yaml:"l1"`
	Level2 string `yaml:"l2"`
	CpuNum int `yaml:"cn"`
	MemSize int `yaml:"ms"`
	Status string `yaml:"status"`
	Usages map[int]MUsage `yaml:"Usages,flow"`
}

type MUsage struct {
	CpuPercent int `yaml:"cp"`
	MemPercent int `yaml:"mp"`
	MemGps float64 `yaml:"mg"`
	Mpki int `yaml:"mpki"`
	DiskIoPercent float64 `yaml:"dip"`
	NetIn float64 `yaml:"ni"`
	NetOut float64 `yaml:"no"`
}

var machines = make(map[string]Machine)

func MachineProcess(line string, size int64)  {
	var machine Machine
	csv := strings.Split(line,",")
	if len(csv) != 7{
		log.Println(line)
		return
	}
	ts ,_:= strconv.Atoi(csv[1])
	l1 ,_:= strconv.Atoi(csv[2])
	cn ,_:= strconv.Atoi(csv[4])
	ms ,_:= strconv.Atoi(csv[5])

	_,ok := machines[csv[0]]

	if ok {
		machine = machines[csv[0]]
	} else {
		machine = Machine{Usages:make(map[int]MUsage),TimeStamp:ts,Level1:l1,Level2:csv[3],CpuNum:cn,MemSize:ms,Status:csv[6]}
	}
	if ts <machine.TimeStamp{
		machine.TimeStamp = ts
	}
	machines[csv[0]] = machine
}

func ExportMachinesYaml(){
	d, err := yaml.Marshal(&machines)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("Statistical information of machines:")
	log.Println("Machine counts: ",len(machines))
	WriteWithFileWrite(Prefix+"Machines.yaml",string(d))
}

