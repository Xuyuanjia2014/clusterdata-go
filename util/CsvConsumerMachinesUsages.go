package util

import (
	"log"
	"time"
	"clusterdata-go/middle"
)

func ConvertMachines()  {
	timestamp := time.Now().Unix()
	log.Println("Comsumer starts at",timestamp)
	var size int64;
	size = 0
	startTime := CurrentTime()
	for x := range Channel{
		//PrintFirst100(size,x)
		middle.MachineProcess(x,size)
		size++;
	}
	endTime := CurrentTime()
	log.Println("Total Consumer Seconds:",(endTime-startTime))
	log.Println("Mahcine meta lines: ",size)
	middle.ExportMachinesYaml()
}

func ConvertUsage()  {
	timestamp := time.Now().Unix()
	log.Println("Comsumer starts at",timestamp)
	var size int64;
	size = 0
	startTime := CurrentTime()
	for x := range Channel{
		//PrintFirst100(size,x)
		middle.MachineProcess(x,size)
		size++;
	}
	endTime := CurrentTime()
	log.Println("Total Consumer Seconds:",(endTime-startTime))
	middle.ExportMachinesYaml()
}