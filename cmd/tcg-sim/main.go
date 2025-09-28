package main

import (
	"fmt"
	"github.com/evevioletrose-hash/tcg-simulator/internal"
)

func main() {
	fmt.Println("Welcome to the TCG Simulator!")
	g := internal.NewGame()
	g.Run()
}
