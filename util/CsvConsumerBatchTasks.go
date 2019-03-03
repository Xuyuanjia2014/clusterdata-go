package util

import (
	"log"
	"time"
	"clusterdata-go/middle"
)

func ConvertCsv()  {
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
	middle.ExoprtBatchTaskYaml()
}

func PrintFirst100(size int64, x string)  {
	if(size <200){
		log.Println(x)
	}
}