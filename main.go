package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func main() {
	gracefulStop := make(chan os.Signal, 1)

	log.Println("Login ")
	err := client.Login("")
	if err != nil {
		panic(err)
	}

	tiemstamp := time.Now()

	log.Println("SetActivity")
	err = client.SetActivity(client.Activity{
		Details: "with LDPlayer 9.0",
		// State:      "Heyy!!!",
		// LargeImage: "largeimageid",
		// LargeText:  "This is the large image :D",
		Timestamps: &client.Timestamps{Start: &tiemstamp},
		Secrets:    &client.Secrets{},
		Buttons:    []*client.Button{},
	})

	if err != nil {
		panic(err)
	}
	signal.Notify(gracefulStop, os.Interrupt, syscall.SIGTERM)
	<-gracefulStop
	log.Println("Graceful Exiting...")
}
