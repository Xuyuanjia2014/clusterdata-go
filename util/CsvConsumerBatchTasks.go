package util

import (
	"log"
	"time"
	"clusterdata-go/middle"
)

func ConvertTasks()  {
	timestamp := time.Now().Unix()
	log.Println("Comsumer starts at",timestamp)
	var size int64;
	size = 0
	startTime := CurrentTime()
	for x := range Channel{
		//PrintFirst100(size,x)
		middle.BatchTaskProcess(x,size)
		middle.CheckDuplication(x,size)
		size++;
	}
	endTime := CurrentTime()
	log.Println("Total tasks' lines:",size)
	middle.ExportDuplication()
	log.Println("Total Consumer Seconds:",(endTime-startTime))
	middle.ExportBatchTaskYaml()
}