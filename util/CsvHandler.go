package util

import (
	"log"
	"fmt"
	"io"
	"time"
	"os"
	"bufio"
)

func ReadCsv(path string)  {
	fileObj, err := os.Open(path)
	size :=0
	if err != nil {
		log.Fatal(err)
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	buf := make([]byte,1024*1024*1024*40)
	timestamp := time.Now().Unix()
	fmt.Println("startAt:",timestamp)
	for {
		_,err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		size++
		fmt.Println("Part size:",size)
	}
	fmt.Println("Total size:",size)
	timestamp2 := time.Now().Unix()
	fmt.Println("endAt:",timestamp2)
	fmt.Println("Total Seconds:",(timestamp2-timestamp))
}
