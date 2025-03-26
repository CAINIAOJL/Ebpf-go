/*package test

import (
	"log"
	"time"
	"math/rand"
	"sync"
)
var wg sync.WaitGroup

func doprint(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done()
}

func main() {
	//rand.Seed(123456)
	log.SetFlags(0)
	wg.Add(2)
	go doprint("hi, jiang", 20)
	go doprint("hi, lei", 20)
	time.Sleep(2 * time.Second)
	wg.Wait()
}