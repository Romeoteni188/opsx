package main

import (
	"log"

	"github.com/Romeoteni188/opsx/internal"
)

func main() {
	if err := internal.NewRootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
