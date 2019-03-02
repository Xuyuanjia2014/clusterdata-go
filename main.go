package main

import (
	"clusterdata-go/util"
	time "time"
)

var batch_instance ="E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_instance.csv"
var batch_task = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_task.csv"
var container_meta = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\container_meta.csv"
var container_usage = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\container_usage.csv"
var machine_meta = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\container_meta.csv"
var machine_usage = "E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\container_usage.csv"

func main() {
	//util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_task.csv")
	//util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_instance.csv")
	go util.ReadCsv(batch_instance)
	time.Sleep(3 * time.Second)
	go util.ConvertCsv()

	time.Sleep(7200 * time.Second)
}