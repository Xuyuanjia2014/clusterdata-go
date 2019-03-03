package middle

import (
	"strings"
	"strconv"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"fmt"
)

var Prefix = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\"

type Batch struct {
	AllJobs map[string]Job `yaml:"Jobs,flow"`
}

type Job struct {
	FatalTag string `yaml:"ft"`
	SelfTasks map[string]Task `yaml:"Tasks,flow"`
	TasksCount int `yaml:"tc"`
}

type Task struct {
	InstanceNumber int `yaml:"in"`
	Type string `yaml:"t"`
	TaskStatus string `yaml:"ts"`
	StartTime int `yaml:"st"`
	EndTime int `yaml:"et"`
	Cpu int `yaml:"c"`
	Mem int `yaml:"m"`
	SelfInstances map[string]Instance `yaml:"Instances,flow"`
}

type Instance struct {
	Type string
	InstanceStatus string
	StartTime int
	EndTime int
	MachineId string
	SeqNo int
	TotalSeqNo int
	CpuAvg float64
	CpuMax float64
	MemAvg float64
	MemMax float64
}

var Jobs Batch =Batch{AllJobs:make(map[string]Job)}

var OneInstance Instance = Instance{Type:"",InstanceStatus:"",StartTime:0,EndTime:0,MachineId:"",
								SeqNo:0,TotalSeqNo:0,CpuAvg:0,CpuMax:0,MemAvg:0,MemMax:0}

var JobTaskSet = make(map[string]int)

func BatchTaskProcess(line string,size int64)  {
	var OneJob Job
	csv := strings.Split(line,",")
	if len(csv) != 9 {
		log.Println("for line: ", size," ; string:",line,"  ERROR!")
		return
	}
	_,ok := Jobs.AllJobs[csv[3]]
	if ok {
		OneJob = Jobs.AllJobs[csv[3]]
	} else {
		OneJob = Job{FatalTag:"none",SelfTasks:make(map[string]Task)}
	}
	in ,_:= strconv.Atoi(csv[1])
	st ,_:= strconv.Atoi(csv[5])
	et ,_:= strconv.Atoi(csv[6])
	c ,_:= strconv.Atoi(csv[7])
	m ,_:= strconv.Atoi(csv[8])
	var OneTask = Task{InstanceNumber:in,Type:csv[1],TaskStatus:csv[4],StartTime:st,EndTime:et,Cpu:c,Mem:m}
	OneJob.SelfTasks[csv[0]] = OneTask
	Jobs.AllJobs[csv[3]] = OneJob
}

func WriteWithFileWrite(name,content string){
	fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _,err := fileObj.WriteString(content);err == nil {
		fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.")
	}
	contents := []byte(content)
	if _,err := fileObj.Write(contents);err == nil {
		fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.")
	}
}

func ExportBatchTaskYaml(){
	d, err := yaml.Marshal(&Jobs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	size := 0
	log.Println("Statistical information of Jobs, tasks and instances:")
	log.Println("Jobs' count:", len(Jobs.AllJobs) )
	for key,value := range Jobs.AllJobs {
		value.TasksCount = len(value.SelfTasks)
		size++;
		log.Println("job Name:",key," ; Tasks' count:", value.TasksCount)
	}
	WriteWithFileWrite(Prefix+"BatchTasks.yaml",string(d))
}

func CheckDuplication(line string,size int64) {
	csv := strings.Split(line,",")
	value,ok := JobTaskSet[csv[3]+csv[0]]
	if ok {
		JobTaskSet[csv[3]+csv[0]] = value+1
	} else {
		JobTaskSet[csv[3]+csv[0]] = 1
	}
}

func ExportDuplication()  {
	log.Println("Total tasks:",len(JobTaskSet))
}





