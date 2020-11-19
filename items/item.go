package items

import (
	"sync"
)

type Adder struct {
	In  chan int
	Out chan int

	Modifyer ModifiyFunc
}

func NewAdder(modifyFunc ModifiyFunc) Item {
	return &Adder{Modifyer: modifyFunc}
}

func (a *Adder) GetIn() chan int {
	return a.In
}

func (a *Adder) GetOut() chan int {
	return a.Out
}

func (a *Adder) SetOut(out chan int) {
	a.Out = out
}

func (a *Adder) SetIn(in chan int) {
	a.In = in
}

func (a *Adder) SetModifier(modifiyFunc ModifiyFunc) {
	a.Modifyer = modifiyFunc
}

func (a *Adder) StartModifying(wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(a.Out)
	if a.In != nil && a.Out != nil {
		for {
			select {
			case x, ok := <-a.In:
				if ok {
					a.Out <- a.Modifyer(x)
				} else {
					return
				}
			}
		}
	}
}
