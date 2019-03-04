package middle

import (
	"strings"
	"strconv"
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
	"os"
	"fmt"
)

type Machine struct {
	TimeStamp int `yaml:"ts"`
	Level1 int `yaml:"l1"`
	Level2 string `yaml:"l2"`
	CpuNum int `yaml:"cn"`
	MemSize int `yaml:"ms"`
	Status string `yaml:"status"`
	Usages map[int]MUsage `yaml:"Usages,flow"`
	UsageCounts int `yaml:"uc"`
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

func ExportMachinesYamlAll(name string){
		d, err := yaml.Marshal(&Machines)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Println("Statistical information of machines:")
		log.Println("Machine counts: ",len(Machines))
		WriteWithFileWrite(Prefix+name+".yaml",string(d))
}

func ExportMachinesYaml(name string){
	size :=0
	fileObj,err := os.OpenFile(Prefix+name+".yaml",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	defer fileObj.Close()

	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}

	for key,value := range Machines {
		TempMap := make(map[string]Machine)
		TempMap["xyjM:"+key] = value
		d, err2 := yaml.Marshal(&TempMap)
		if err2 != nil {
			log.Fatalf("error: %v", err)
		}
		n, _ := fileObj.Seek(0,os.SEEK_END)
		_,err2 = fileObj.WriteAt(d,n)
		log.Println("for Mahcine:", size)
		size++
	}
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
	log.Println("load one day data.")
}

func ReadAll(path string) []byte{
	if contents,err := ioutil.ReadFile(path);err == nil {
		return contents
	}
	return nil
}

func MachineUsageProcess(line string, size int64)  {
	var machine Machine
	var usages map[int]MUsage

	csv := strings.Split(line,",")
	ts ,_:= strconv.Atoi(csv[1])

	if ts >= 3600*12{
		return
	}
	if len(csv) != 9 {
		log.Println("Bad line:"+line)
		return
	}
	_,ok := Machines[csv[0]]
	if !ok {
		return
	}

	machine = Machines[csv[0]]
	if machine.Usages != nil {
		usages = machine.Usages
	} else {
		usages = make(map[int]MUsage)
	}

	cup ,_:= strconv.Atoi(csv[2])
	mup ,_:= strconv.Atoi(csv[3])
	mg ,_:= strconv.ParseFloat(csv[4],64)
	mpki,_ := strconv.Atoi(csv[5])
	ni ,_:= strconv.ParseFloat(csv[6],64)
	no ,_:= strconv.ParseFloat(csv[7],64)
	dip ,_:= strconv.ParseFloat(csv[8],64)

	usages[ts] = MUsage{CpuPercent:cup,MemPercent:mup,MemGps:mg,Mpki:mpki,NetIn:ni,NetOut:no,DiskIoPercent: dip}
	machine.Usages = usages

	Machines[csv[0]] = machine
}

func MachineUsageCountProcess(line string, size int64)  {
	var machine Machine

	csv := strings.Split(line,",")

	if len(csv) != 9 {
		log.Println("Bad line:"+line)
		return
	}
	machine,ok := Machines[csv[0]]
	if !ok {
		return
	}
	machine.UsageCounts++
	Machines[csv[0]] = machine
}

