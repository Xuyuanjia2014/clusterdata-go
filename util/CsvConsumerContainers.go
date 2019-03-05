package util

import (
	"log"
	"time"
	"clusterdata-go/middle"
	"gopkg.in/yaml.v2"
	"clusterdata-go/statistics"
	"clusterdata-go/drawer"
)

func ConvertContainerMeta()  {
	timestamp := time.Now().Unix()
	log.Println("Comsumer starts at",timestamp)
	var size int64;
	size = 0
	startTime := CurrentTime()
	for x := range Channel{
		//PrintFirst100(size,x)
		middle.ContainersProcess(x,size)
		size++;
	}
	endTime := CurrentTime()
	log.Println("Total tasks' lines:",size)
	log.Println("Total Consumer Seconds:",(endTime-startTime))
	middle.ExportContainersYamlAll("containersMeta")
}

func ConvertContainersUsage()  {
	timestamp := time.Now().Unix()
	log.Println("Comsumer starts at",timestamp)
	var size int64;
	size = 0
	startTime := CurrentTime()
	for x := range Channel{
		//PrintFirst100(size,x)
		middle.ContainerUsageCountProcess(x,size)
		size++;
		if(size %10000000 == 0){
			log.Println("size: ",size)
			log.Println(x)
		}
	}
	endTime := CurrentTime()
	log.Println("size: ",size)
	log.Println("Total Consumer Seconds:",(endTime-startTime))
	middle.ExportContainersYaml("containerusagesCounts")
}

func ConvertBigContainerObject()  {
	timestamp := time.Now().Unix()
	log.Println("Comsumer starts at",timestamp)
	var size int64;
	size = 0
	startTime := CurrentTime()
	for x := range Channel{
		container := make(map[string]middle.Container)
		err := yaml.Unmarshal([]byte(x),&container)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		for key,value := range container {
			//log.Println("For machine:", key, " ; sth:", len(value.Usages))
			middle.Containers[key] = value
			if(size %1000 == 0){
				log.Println("size: ",size," ;usage size:", len(value.Usages))
			}
		}
		size++;
	}
	endTime := CurrentTime()
	log.Println("Total Consumer Seconds:",(endTime-startTime))
	log.Println("Container meta lines: ",size)
	statistics.FindContainers()
	drawer.Test()
}