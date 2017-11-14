package main

import (
	"log"
)

func main() {
	err := xdg.Open("https://me.gsora.xyz")
	if err != nil {
		log.Fatal(err)
	}
}
