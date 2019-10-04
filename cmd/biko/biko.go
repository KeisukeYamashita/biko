package main

import (
	"log"

	"github.com/KeisukeYamashita/biko/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
