package main

import (
	"log"

	"github.com/gsora/xdg"
)

func main() {
	err := xdg.Open("https://me.gsora.xyz")
	if err != nil {
		log.Fatal(err)
	}
}
