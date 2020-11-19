package items

import (
	"context"
	"sync"
)

type Loader struct {
	Out  chan int
	Load LoadFunc
}

func NewLoader(loadfunc LoadFunc) Head {
	return &Loader{Load: loadfunc}
}

func (l *Loader) GetOut() chan int {
	return l.Out
}

func (l *Loader) SetOut(chanOut chan int) {
	l.Out = chanOut
}

func (l *Loader) SetLoad(loadfunc LoadFunc) {
	l.Load = loadfunc
}

func (l *Loader) StartLoading(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(l.Out)
	if l.Out != nil {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				l.Out <- l.Load()
			}
		}
	}
}
