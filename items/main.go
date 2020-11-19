package items

import (
	"context"
	"sync"
)

type LoadFunc func() int

type ReciveFunc func(int)

type ModifiyFunc func(int) int

type Head interface {
	GetOut() chan int
	SetOut(chan int)

	SetLoad(LoadFunc)

	StartLoading(context.Context, *sync.WaitGroup)
}

type Tail interface {
	GetIn() chan int
	SetIn(chan int)

	SetRecive(ReciveFunc)

	StartReciving(*sync.WaitGroup)
}

type Item interface {
	GetIn() chan int
	GetOut() chan int

	SetOut(chan int)
	SetIn(chan int)

	SetModifier(ModifiyFunc)

	StartModifying(*sync.WaitGroup)
}
