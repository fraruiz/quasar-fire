package main

import (
	"github.com/franciscoruizar/quasar-fire/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
