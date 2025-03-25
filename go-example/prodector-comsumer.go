/*package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	//"sync"
	//"time"
)

//生产者
func prodector(factor int, out chan<- int) {
	for i := 0;;i++ {
		out <- i * factor
	}
}

func consumer(in <-chan int) {
	for v := range in {
		log.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)

	go prodector(3, ch)
	go prodector(5, ch)
	go consumer(ch)
	sig := make(chan os.Signal, 1)


	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//time.Sleep(20 * time.Second)
	fmt.Printf("quit (%v)\n", <-sig)
}*/