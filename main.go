package main

import (
	"clusterdata-go/util"
	time "time"
)

func main() {
	//util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_task.csv")
	//util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_instance.csv")
	go util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_instance.csv")
	time.Sleep(3 * time.Second)
	go util.ConvertCsv()

	time.Sleep(20 * time.Second)
}