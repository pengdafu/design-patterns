package main

import "fmt"

type GamePlay interface {
	Play()
	GameOperation
}
type GameOperation interface {
	Initialize()
	StartPlay()
	EndPlay()
}

type Game struct {
	GameOperation
}

func (g *Game) Play() {
	g.Initialize()
	g.StartPlay()
	g.EndPlay()
}

type Cricket struct {
}

func (c Cricket) Initialize() {
	fmt.Println("Cricket Game Initialized!")
}

func (c Cricket) StartPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game!")
}

func (c Cricket) EndPlay() {
	fmt.Println("Cricket Game Finished!")
}

type Football struct {
}

func (f Football) Initialize() {
	fmt.Println("Football Game Initialized!")
}

func (f Football) StartPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}

func (f Football) EndPlay() {
	fmt.Println("Football Game Finished!")
}

func main() {
	cricketGame := Game{Cricket{}}
	cricketGame.Play()
	fmt.Println()
	footballGame := Game{Football{}}
	footballGame.Play()
}
