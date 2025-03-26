/*package main

import (
	"io"
	"net/http"
	"sync"
)

func dohttp(curl string) (interface{}, error) {
	resp, err := http.Get(curl)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}
type Func func (key string)(interface{}, error)
type result struct {
	value interface {}
	err   error
}
type request struct {
	key 		string
	response    chan <- result
}
type entry struct {
	res result
	ready chan struct{} 
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo :=&Memo{requests: make(chan request)}
	go memo.server(f) //开启goroutine //监控goroutine
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <- response
	return res.value,res.err //返回文本，与错误
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	//遍历每一个request
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			//第一次调用
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response<- e.res
}