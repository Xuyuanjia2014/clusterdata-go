package main

import (
	"time"
	"clusterdata-go/util"
)

var Batch_instance ="E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_instance.csv"
var Batch_task = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_task.csv"
var Container_meta = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\container_meta.csv"
var Container_usage = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\container_usage.csv"
var Machine_meta = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\machine_meta.csv"
var Machine_usage = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\machine_usage.csv"

func main() {
	//util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_task.csv")
	//util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_instance.csv")
	//go util.ReadCsv(Batch_task)
	//time.Sleep(3 * time.Second)
	//go util.ConvertTasks()

	//go util.ReadCsv(Machine_meta)
	//time.Sleep(1 * time.Second)
	//go util.ConvertMachines()

	//middle.InitMachineMeta("Machines.yaml")
	//log.Println("For example, m_3:", middle.Machines["m_3"])
	//go util.ReadCsv(Machine_usage)
	//time.Sleep(3 * time.Second)
	//go util.ConvertMachinesCounts()
	//
	go util.ReadCsv(Container_meta)
	time.Sleep(3 * time.Second)
	go util.ConvertContainerMeta()

	//middle.InitContainerMeta("containersMeta.yaml")
	//go util.ReadCsv(Container_usage)
	//time.Sleep(3 * time.Second)
	//go util.ConvertContainersUsage()

	time.Sleep(7200 * time.Second)
}