/*package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc func(v interface{})bool
)

//特别向信号注册相关代码结构
type Publisher struct {
	m     			sync.Mutex
	buffer 			int
	timeout    		time.Duration
	subscriber   	map[subscriber]topicFunc
}

func NerPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer: buffer,
		timeout: publishTimeout,
		subscriber: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p * Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscriber[ch]=topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) QuitSub(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	//删除
	delete(p.subscriber, sub)
	close(sub)
}

func (p *Publisher) Closed() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscriber {
		delete(p.subscriber, sub)
		close(sub)
	}
}


func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup

	for sub, topic := range p.subscriber {
		wg.Add(1)
		go p.sendtopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) sendtopic(sub chan interface{}, 
							  topic topicFunc, 
							  v interface{}, 
							  wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	//if topic != nil {
		//return
	//}

	select{
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func main() {
	p := NerPublisher(100* time.Millisecond, 5)
	defer p.Closed()
	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		//格式转换：由interface{} -> string
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello, world")
	p.Publish("hello, golang")

	go func ()  {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	} ()

	go func ()  {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	} ()

	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//fmt.Printf("quit (%v)\n", <-sig)

	time.Sleep(10 * time.Second)
}
*/