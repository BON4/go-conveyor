package main

import (
	"conveyor/conveyor"
	"conveyor/items"
	"conveyor/line"
	"log"
	"math/rand"
	"time"
)

func modify1(x int) int {
	return x + 1
}

func modify2(x int) int {
	return x * 2
}

func load() int {
	return rand.Intn(10)
}

func recive(x int) {
	log.Println("Recived: ", x)
}

func main() {
	myLine := line.NewConveyorLine()

	myLine.AppendItem(items.NewAdder(modify1))
	myLine.AppendItem(items.NewAdder(modify2))

	log.Println(*myLine.Front())
	log.Println(*myLine.Back())

	loader := items.NewLoader(load)
	reciver := items.NewReciver(recive)

	myconv, err := conveyor.NewConveyor(loader, reciver, myLine)

	if err != nil {
		log.Panicln(err)
	}

	myconv.Strat()

	time.Sleep(time.Second * 10)

	myconv.Stop()
}
