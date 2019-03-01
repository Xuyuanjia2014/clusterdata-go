package util

import (
	"log"
	"time"
)

func ConvertCsv()  {
	timestamp := time.Now().Unix()
	log.Println("Comsumer starts at",timestamp)
	var size int64;
	size = 0
	startTime := CurrentTime()
	for x := range Channel{
		//log.Println(x)
		size++;
		if size%10000000 == 0{
			log.Println("size ten million:",size)
			log.Println(x)
		}
	}
	endTime := CurrentTime()
	log.Println("Total Consumer Seconds:",(endTime-startTime))
}