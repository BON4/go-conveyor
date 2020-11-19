package items

import "context"

type LoadFunc func() int

type ReciveFunc func(int)

type ModifiyFunc func(int) int

type Head interface {
	GetOut() chan int
	SetOut(chan int)

	SetLoad(LoadFunc)

	StartLoading(context.Context)
}

type Tail interface {
	GetIn() chan int
	SetIn(chan int)

	SetRecive(ReciveFunc)

	StartReciving(context.Context)
}

type Item interface {
	GetIn() chan int
	GetOut() chan int

	SetOut(chan int)
	SetIn(chan int)

	SetModifier(ModifiyFunc)

	StartModifying(context.Context)
}
type Node interface {
	GetItem() Item

	Next() *Node
	Prev() *Node
}

type Line interface {
	Front() *Node
	Back() *Node
}
