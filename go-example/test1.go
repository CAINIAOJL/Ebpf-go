/*package main

import (
	"log"
	"sync"
)

type Total struct {
	sync.Mutex
	sum int
}

var total Total

func doprint(tg *sync.WaitGroup) {
	defer tg.Done()

	for i := 0; i < 100; i++ {
		total.Mutex.Lock()
		total.sum++
		total.Mutex.Unlock()
	}
	log.Println(total.sum)
}

func main() {
	var tg sync.WaitGroup
	tg.Add(2)
	go doprint(&tg)
	go doprint(&tg)
	tg.Wait()

	log.Println("main gorouting over")

}*/