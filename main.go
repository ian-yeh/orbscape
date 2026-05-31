package main

import (
	"log"
	"github.com/ian-yeh/orbscape/src"
)

func main() {
	if err := src.NewGame(); err != nil {
		log.Fatal(err)
	}
}

