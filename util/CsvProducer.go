package util

import (
	"log"
	"io"
	"time"
	"os"
	"bufio"
)

var Channel = make(chan string, 100000)

func CurrentTime() int64 {
	timestamp := time.Now().Unix()
	log.Println("Current start or end At:",timestamp)
	return timestamp;
}

func ReadCsv(path string)  {
	fileObj, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	startTime := CurrentTime()
	for {
		line,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		Channel <- line
	}
	endTime := CurrentTime()
	log.Println("Total Producer Seconds:",(endTime-startTime))
	close(Channel)
}
