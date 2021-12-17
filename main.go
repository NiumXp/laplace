package main

import (
	"log"

	laplace "github.com/NiumXp/laplace/internal"
)

func main() {
	map_ := laplace.NewMap(11, 9)
	game := laplace.NewGame(map_, nil)

	for index := uint(0); index < 1; index++ {
		state := game.StateAt(index)
		log.Println(state)
	}
}
