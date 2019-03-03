package middle

import (
	"strings"
	"log"
	"strconv"
	"gopkg.in/yaml.v2"
)

type AppDu struct {
	Containers map[string]Container `yaml:"containers,flow"`
	Counts int `yaml:"cts"`
}

type Container struct {
	MachineId string `yaml:"mid"`
	TimeStamp int `yaml:"ts"`
	CpuRequest int `yaml:"cr"`
	CpuLimit int `yaml:"cl"`
	MemSize float64 `yaml:"ms"`
	Status string `yaml:"sta"`
	Ucounts int `yaml:"ucs"`
	Usages []CUsage `yaml:"Usages,flow"`
}

type CUsage struct {
	MachineId string `yaml:"mid"`
	TimeStamp int `yaml:"ts"`
	CpuPercent int `yaml:"cp"`
	Mpki int `yaml:"mpki"`
	Cpi float64 `yaml:"cpi"`
	MemPercent int `yaml:"mp"`
	MemGps float64 `yaml:"mg"`
	DiskIoPercent float64 `yaml:"dip"`
	NetIn float64 `yaml:"ni"`
	NetOut float64 `yaml:"no"`
}

//var App AppDu = AppDu{Containers:make(map[string]Container)}
var Apps = make(map[string]AppDu)

func ContainersProcess(line string, size int64)  {
	var app AppDu
	csv := strings.Split(line,",")
	if len(csv) != 8 {
		log.Println("for line: ", size," ; string:",line,"  ERROR!")
		return
	}
	_,ok := Apps[csv[3]]
	if ok {
		app = Apps[csv[3]]
	} else {
		app = AppDu{Containers:make(map[string]Container)}
	}

	ts ,_:= strconv.Atoi(csv[2])
	cr ,_:= strconv.Atoi(csv[5])
	cl ,_:= strconv.Atoi(csv[6])
	ms ,_:= strconv.ParseFloat(csv[7],64)

	app.Containers[csv[0]] = Container{MachineId:csv[1],TimeStamp: ts,CpuRequest:cr,CpuLimit: cl,MemSize: ms,Status:csv[4]}
	app.Counts = len(app.Containers)

	Apps[csv[3]] = app
}

func ExportContainersYamlAll(name string){
	d, err := yaml.Marshal(&Apps)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("Statistical information of containers:")
	log.Println("apps counts: ",len(Apps))
	cc :=0
	for _,value := range Apps {
		cc += value.Counts
	}
	log.Println("container counts: ",cc)
	WriteWithFileWrite(Prefix+name+".yaml",string(d))
}

func InitContainerMeta(name string){
	contents := ReadAll(Prefix+name)
	if contents ==nil{
		return
	}
	err := yaml.Unmarshal(contents,&Apps)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("Load containers' meta information")
}

func ContainerUsageProcess(line string, size int64)  {
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