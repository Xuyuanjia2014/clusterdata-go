package main

import (
	"clusterdata-go/util"
	"runtime/debug"
)

func main() {
	debug.SetMaxStack(1024*1024*20)
	//util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_task.csv")
	util.ReadCsv("E:\\benchmark\\alibaba_clusterdata2018\\alibaba_clusterdata_v2018\\batch_instance.csv")
}
