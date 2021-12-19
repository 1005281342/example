package main

import (
	"log"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func main() {
	var c, err = beanstalk.Dial("tcp", "127.0.0.1:11300")
	var taskID uint64

	taskID, err = c.Put([]byte("world"), 1024, 0, 1*time.Second)
	log.Println(taskID, err)

	taskID, err = c.Put([]byte("hello"), 1, 0, 1*time.Second)
	log.Println(taskID, err)
}
