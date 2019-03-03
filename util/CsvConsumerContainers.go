package util

import (
	"log"
	"time"
	"clusterdata-go/middle"
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
		middle.ContainerUsageProcess(x,size)
		size++;
		if(size %10000000 == 0){
			log.Println("size: ",size)
			log.Println(x)
		}
	}
	endTime := CurrentTime()
	log.Println("size: ",size)
	log.Println("Total Consumer Seconds:",(endTime-startTime))
	middle.ExportContainersYaml("containerusages")
}