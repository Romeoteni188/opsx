package main

import (
	"log"

	"github.com/Romeoteni188/opsx/internal/cli"
)

func main() {
	if err := cli.NewRootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
