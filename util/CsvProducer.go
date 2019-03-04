package util

import (
	"log"
	"io"
	"time"
	"os"
	"bufio"
	"strings"
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

func ReadCsvObject(path string)  {
	fileObj, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	startTime := CurrentTime()
	content:=""
	for {
		line,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if(strings.HasPrefix(line,"xyjM:") && content ==""){
			content+=line
			continue
		}
		if(strings.HasPrefix(line,"xyjM:") && strings.HasPrefix(content,"xyjM:")){
			Channel <- strings.Replace(content,"xyjM:","",1)
			content=line
			continue
		}
		content+=line
	}
	close(Channel)
	endTime := CurrentTime()
	log.Println("Total Producer Seconds:",(endTime-startTime))
}
