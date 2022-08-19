package xsync

import (
	"sync"
)

//type Pool[T any] interface {
//	Put(x T)
//	Get() T
//}

type Pool[T any] struct {
	p   sync.Pool
	New func() T
}

func (p *Pool[T]) Put(x T) {
	p.p.Put(x)
}

func (p *Pool[T]) Get() (t T) {
	v := p.p.Get()
	if v != nil {
		return v.(T)
	}
	if v == nil && p.New != nil {
		t = p.New()
	}
	return t
}
