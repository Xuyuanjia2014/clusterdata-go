package statistics

import (
	"clusterdata-go/middle"
	"log"
)

var Counts2Apps map[int]int =make(map[int]int)

func SimpleFindCountsInApp()  {
	for _,value := range middle.Apps {
		Counts2Apps[value.Counts]++;
	}
	log.Println(len(Counts2Apps))
	log.Println(Counts2Apps)
}
