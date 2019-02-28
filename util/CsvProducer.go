package util

import (
	"log"
	"fmt"
	"io"
	"time"
	"os"
	"bufio"
)

var Channel = make(chan string, 100000)

func ReadCsv(path string)  {
	fileObj, err := os.Open(path)
	size :=0
	if err != nil {
		log.Fatal(err)
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	timestamp := time.Now().Unix()
	fmt.Println("startAt:",timestamp)
	for {
		line,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		size++
		Channel <- line
	}
	fmt.Println("Total size:",size)
	timestamp2 := time.Now().Unix()
	fmt.Println("endAt:",timestamp2)
	fmt.Println("Total Seconds:",(timestamp2-timestamp))
}
