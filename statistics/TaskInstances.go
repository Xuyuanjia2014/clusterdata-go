package statistics

import (
	"clusterdata-go/middle"
	"log"
)

var TaskInstanceCounts map[int]int =make(map[int]int)

func Calculate()  {
	var max int
	var size int
	var i int
	size = 0;
	for i =0;i<=100;i++{
		TaskInstanceCounts[i] = 0;
	}
	for _,value := range middle.Jobs.AllJobs {
		for _,task := range value.SelfTasks {
			if task.InstanceNumber >max {
				max = task.InstanceNumber
			}
			size++
			ResetMap(task.InstanceNumber)
		}
	}
	log.Println(TaskInstanceCounts)
	log.Println("task size:",size," ; max:",max)
	TotalInstances()
}

func ResetMap(number int)  {
	TaskInstanceCounts[number]++
}

func TotalInstances()  {
	var sum int
	sum = 0
	for key,value := range TaskInstanceCounts {
		sum += key  * value
	}
	log.Println("instance sum:",sum)
}
