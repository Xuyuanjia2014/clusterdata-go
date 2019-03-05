package statistics

import (
	"clusterdata-go/middle"
	"sort"
	"log"
)

type BResourceUsage struct {
	Max int
	Min int
	Avg int
}

var Clines [101]BResourceUsage

func FindContainers()  {
		for index:=0;index<=100;index++ {
			var ru = BResourceUsage{Avg:0,Max:0,Min:100}
			sum :=0
			size :=0
			for _,value := range middle.Containers {
				if len(value.Usages) != 101{
					continue
				}
				var tss []int
				tss = make([]int,101)
				i :=0
				for ts,_ := range value.Usages {
					tss[i] = ts
					i++
				}
				sort.Ints(tss)
				sum +=value.Usages[tss[index]].CpuPercent
				if value.Usages[tss[index]].CpuPercent >ru.Max && value.Usages[tss[index]].CpuPercent < 100 {
					ru.Max = value.Usages[tss[index]].CpuPercent
				}
				if value.Usages[tss[index]].CpuPercent < ru.Min && value.Usages[tss[index]].CpuPercent >0{
					ru.Min = value.Usages[tss[index]].CpuPercent
				}
				size ++
			}
			ru.Avg = sum/size
			Clines[index] = ru
		}
		log.Println(Clines)
}
