package main

import (
	"github.com/beanstalkd/go-beanstalk"
	"log"
	"sync"
	"time"
)

func main() {
	var c, err = beanstalk.Dial("tcp", "127.0.0.1:11300")
	var (
		id   uint64
		body []byte
		wg   = &sync.WaitGroup{}
	)
	wg.Add(1)
	go func() {
		id, body, err = c.Reserve(5 * time.Second)
		log.Println(id, string(body), err)
		wg.Done()
	}()
	//c.Delete(id)

	wg.Add(1)
	go func() {
		id, body, err = c.Reserve(5 * time.Second)
		log.Println(id, string(body), err)
		wg.Done()
	}()
	//c.Delete(id)

	wg.Wait()
}
