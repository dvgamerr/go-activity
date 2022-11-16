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
	err := client.Login(os.Args[1])
	if err != nil {
		panic(err)
	}

	tiemstamp := time.Now()

	log.Println("SetActivity")
	err = client.SetActivity(client.Activity{
		Details: os.Args[2],
		// State:      "Heyy!!!",
		SmallImage: "image",
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
