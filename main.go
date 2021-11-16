package main

import (
	"log"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	dir, _ := os.Getwd()

	file, err := os.OpenFile(dir+"/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}

	logger = log.New(file, "", log.LstdFlags)
}

func main() {
	pid := os.Getpid()
	logger.Printf("service[%d] started...", pid)

	for i := 1; i > 0; i++ {
		time.Sleep(time.Second * 1)
		logger.Printf("service[%d]: %d", pid, i)
	}

	logger.Printf("service[%d] stopped", pid)
}
