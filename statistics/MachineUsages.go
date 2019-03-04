package statistics

import (
	"clusterdata-go/middle"
	"log"
)

type ResourceUsage struct {
	max int
	min int
	avg int
}

var Tuages map[string][12]ResourceUsage =make(map[string][12]ResourceUsage)

var Result map[string][12]int = make(map[string][12]int)

func GetOneDayUsage()  {
	for key,value := range middle.Machines {
		var machine12 [12]ResourceUsage

		var buckets [12]map[int]middle.MUsage

		for i,_ := range buckets {
			buckets[i] =make(map[int]middle.MUsage)
		}

		for ts,usage := range value.Usages {
			buckets[ts/3600][ts] = usage
		}

		for index,value := range buckets {
			sum :=0
			max :=0
			min :=100
			for _,usage := range value {
				sum+=usage.CpuPercent
				if usage.CpuPercent > max{
					max = usage.CpuPercent
				}
				if usage.CpuPercent < min {
					min = usage.CpuPercent
				}
			}
			if len(value) <= 0{
				sum = 0
			} else{
				sum = sum/len(value)
			}
			machine12[index] = ResourceUsage{max:max,min:min,avg:sum}
		}
		Tuages[key] = machine12
	}
	log.Println("finish the calculation of each machine' 12 hours avg cpu.")
	var i int
	var maxl [12]int
	var minl [12]int
	var avgl [12]int
	for i =0;i <12 ;i++ {
		sum :=0
		max :=0
		min :=100
		for _,value := range Tuages {
			sum += value[i].avg
			if value[i].max > max{
				max = value[i].max
			}
			if value[i].min < min && value[i].min > 0 {
				min = value[i].min
			}
		}
		if len(Tuages) <= 0{
			sum = 0
		} else{
			sum = sum/len(Tuages)
		}
		maxl[i] = max
		minl[i] = min
		avgl[i] = sum
	}
	log.Println("finish the calculation of total avg cpu.")
	Result["max"] = maxl
	Result["min"] = minl
	Result["avg"] = avgl

	log.Println(Result)
}
