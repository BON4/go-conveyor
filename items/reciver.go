package items

import (
	"context"
	"sync"
)

type Reciver struct {
	In     chan int
	Recive ReciveFunc
}

func NewReciver(reciveFunc ReciveFunc) Tail {
	return &Reciver{Recive: reciveFunc}
}

func (r *Reciver) GetIn() chan int {
	return r.In
}

func (r *Reciver) SetIn(chanIn chan int) {
	r.In = chanIn
}

func (r *Reciver) SetRecive(reciveFunc ReciveFunc) {
	r.Recive = reciveFunc
}

func (r *Reciver) StartReciving(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	if r.In != nil {
		for {
			select {
			case <-ctx.Done():
				//close(r.In)
				return
			case x := <-r.In:
				r.Recive(x)
			}
		}
	}
}
