package conveyor

import (
	"context"
	items "conveyor/items"
	line "conveyor/line"
	"errors"
	"log"
	"sync"
)

type Conveyor struct {
	conveyorLine line.Line
	loader       items.Head
	reciver      items.Tail
	wg           sync.WaitGroup
	cancel       context.CancelFunc
}

func (c *Conveyor) Strat() {
	//Set the waitGroup and pass it to the loader, reciver and to all Items in conveyorLine
	//Same with context
	_, c.cancel = context.WithCancel(context.Background())
}

func (c *Conveyor) Stop() {
	c.cancel()
	c.wg.Wait()
	log.Panicln("Conveyor stoped !")
}

func NewConveyor(loader items.Head, reciver items.Tail, line line.Line) (*Conveyor, error) {
	//SetUp the line
	for e := *line.Front(); e != nil; e = *e.Next() {
		if e.GetItem() == nil {
			return nil, errors.New("Some of the node dont have corresponded items.Item")
		}
		//SetUp First
		if e.Prev() == nil {
			if e.Next() != nil {
				//Make chan to connect two items
				c := make(chan int)

				//Get node.curr and node.Prev.Item
				curr := e.GetItem()
				next := (*e.Next()).GetItem()

				//Connect
				curr.SetOut(c)
				next.SetIn(c)
			} else {
				break
			}
			//SetUp Last
		} else if e.Next() == nil {
			if e.Prev() != nil {
				//Make chan to connect two items
				c := make(chan int)

				//Get node.curr and node.Next.Item
				curr := e.GetItem()
				prev := (*e.Prev()).GetItem()

				//Connect
				curr.SetIn(c)
				prev.SetOut(c)
			} else {
				break
			}
			//SetUp Middle
		} else {
			if e.Next() != nil && e.Prev() != nil {
				//Create two channels to connect items
				//{item1 (Out)}->(chanIn)->{(In) current (Out)}->(chanOut)->{(In) item2}
				chanIn := make(chan int)
				chanOut := make(chan int)
				curr := e.GetItem()
				prev := (*e.Prev()).GetItem()
				next := (*e.Next()).GetItem()
				if curr.GetIn() != nil {
					curr.SetIn(chanIn)
					prev.SetOut(chanIn)
				}
				if curr.GetOut() != nil {
					curr.SetOut(chanOut)
					next.SetIn(chanOut)
				}
			} else {
				break
			}
		}
	}

	//SetUp Head
	if line.Front() != nil {
		chanOut := make(chan int)
		first := (*line.Front()).GetItem()
		loader.SetOut(chanOut)
		first.SetIn(chanOut)
	} else {
		return nil, errors.New("Line is empty")
	}

	if line.Back() != nil {
		chanIn := make(chan int)
		last := (*line.Back()).GetItem()
		reciver.SetIn(chanIn)
		last.SetOut(chanIn)
	} else {
		return nil, errors.New("Line is empty")
	}

	return &Conveyor{conveyorLine: line, loader: loader, reciver: reciver}, nil
}
