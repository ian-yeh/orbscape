package main

import (
	"log"
	"github.com/ian-yeh/orbscape/game"
)

func main() {
	if err := game.NewGame(); err != nil {
		log.Fatal(err)
	}
}

