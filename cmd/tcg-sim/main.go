package main

import (
	"fmt"
	"github.com/evevioletrose-hash/tcg-simulator/internal/game"
)

func main() {
	fmt.Println("Welcome to the TCG Simulator!")
	g := game.NewGame()
	g.Run()
}
