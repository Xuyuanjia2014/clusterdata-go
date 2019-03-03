package middle

import (
	"strings"
	"strconv"
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
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

var Machines = make(map[string]Machine)

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

	_,ok := Machines[csv[0]]

	if ok {
		machine = Machines[csv[0]]
	} else {
		machine = Machine{Usages:make(map[int]MUsage),TimeStamp:ts,Level1:l1,Level2:csv[3],CpuNum:cn,MemSize:ms,Status:csv[6]}
	}
	if ts <machine.TimeStamp{
		machine.TimeStamp = ts
	}
	Machines[csv[0]] = machine
}

func ExportMachinesYaml(){
	d, err := yaml.Marshal(&Machines)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("Statistical information of machines:")
	log.Println("Machine counts: ",len(Machines))
	WriteWithFileWrite(Prefix+"Machines.yaml",string(d))
}

func InitMachineMeta(name string){
	contents := ReadAll(Prefix+name)
	if contents ==nil{
		return
	}
	err := yaml.Unmarshal(contents,&Machines)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func ReadAll(path string) []byte{
	if contents,err := ioutil.ReadFile(path);err == nil {
		return contents
	}
	return nil
}

