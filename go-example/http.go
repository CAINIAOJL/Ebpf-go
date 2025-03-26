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
type entry struct {
	res result
	ready chan struct{} 
}

type Memo struct {
	f			Func
	mu 			sync.Mutex
	cache 		map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{
		f: f,
		cache: make(map[string]*entry),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value,e.res.err
}
