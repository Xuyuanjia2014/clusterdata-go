package middle

import (
	"strings"
	"log"
	"strconv"
	"gopkg.in/yaml.v2"
	"os"
	"fmt"
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
	Usages map[int]CUsage `yaml:"Usages,flow"`
}

type CUsage struct {
	MachineId string `yaml:"mid"`
	//TimeStamp int `yaml:"ts"`
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
var Containers = make(map[string]Container)

func ContainersProcess(line string, size int64)  {
	var app AppDu
	csv := strings.Split(line,",")
	if len(csv) != 8 {
		log.Println("for line: ", size," ; string:",line,"  ERROR!")
		return
	}

	ts ,_:= strconv.Atoi(csv[2])

	_,ok := Apps[csv[3]]
	if ok {
		app = Apps[csv[3]]
	} else {
		app = AppDu{Containers:make(map[string]Container)}
	}
	container,ok2 :=app.Containers[csv[0]]
	if ok2 && container.TimeStamp < ts {
		ts = container.TimeStamp
	}
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
	for _,value1 := range Apps {
		for key2,value2 := range value1.Containers {
			Containers[key2] = value2
		}
	}
	log.Println("Load containers' meta information:", len(Containers))
}

func ContainerUsageProcess(line string, size int64)  {
	var container Container
	var usages map[int]CUsage

	csv := strings.Split(line,",")
	ts ,_:= strconv.Atoi(csv[2])

	if len(csv) != 11 {
		log.Println("Bad line:"+line)
		return
	}
	_,ok := Containers[csv[0]]
	if !ok {
		return
	}

	container = Containers[csv[0]]

	if container.Usages != nil {
		usages = container.Usages
	} else {
		usages = make(map[int]CUsage)
	}

	if len(usages) >100 {
		return
	}

	cup ,_:= strconv.Atoi(csv[3])
	mpki,_ := strconv.Atoi(csv[7])
	cpi ,_:= strconv.ParseFloat(csv[5],64)
	mup ,_:= strconv.Atoi(csv[4])
	mg ,_:= strconv.ParseFloat(csv[6],64)
	dip ,_:= strconv.ParseFloat(csv[10],64)
	ni ,_:= strconv.ParseFloat(csv[8],64)
	no ,_:= strconv.ParseFloat(csv[9],64)

	usages[ts] = CUsage{MachineId: csv[1],CpuPercent:cup,Mpki:mpki,Cpi:cpi,MemPercent:mup,MemGps:mg,NetIn:ni,NetOut:no,DiskIoPercent: dip}
	container.Usages = usages
	container.Ucounts = len(usages)

	Containers[csv[0]] = container
}

func ContainerUsageCountProcess(line string, size int64)  {
	var container Container

	csv := strings.Split(line,",")

	container = Containers[csv[0]]

	container.Ucounts++

	Containers[csv[0]] = container
}

func ExportContainersYaml(name string){
	size :=0
	fileObj,err := os.OpenFile(Prefix+name+".yaml",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	defer fileObj.Close()

	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}

	for key,value := range Containers {
		TempMap := make(map[string]Container)
		TempMap["xyjC:"+key] = value
		d, err2 := yaml.Marshal(&TempMap)
		if err2 != nil {
			log.Fatalf("error: %v", err)
		}
		n, _ := fileObj.Seek(0,os.SEEK_END)
		_,err2 = fileObj.WriteAt(d,n)
		size++
	}
	log.Println("for Container:", size)
}

