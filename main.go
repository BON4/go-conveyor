package main

import (
	"conveyor/conveyor"
	"conveyor/items"
	"conveyor/line"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

type closeHandler func()

func SetupCloseHandler(closeFun closeHandler) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	closeFun()
	os.Exit(0)
}

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

	loader := items.NewLoader(load)
	reciver := items.NewReciver(recive)

	myconv, err := conveyor.NewConveyor(loader, reciver, myLine)

	if err != nil {
		log.Panicln(err)
	}

	myconv.Strat()

	SetupCloseHandler(myconv.Stop)
}
