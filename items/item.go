package items

import "context"

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

func (a *Adder) StartModifying(ctx context.Context) {
	if a.In != nil && a.Out != nil {
		for {
			select {
			case <-ctx.Done():
				//close(a.In)
				//close(a.Out)
				return
			case x := <-a.In:
				a.Out <- a.Modifyer(x)
			}
		}
	}
}
