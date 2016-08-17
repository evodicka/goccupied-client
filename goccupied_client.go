package main

import (
	"github.com/evodicka/goccupied-client/client"
	"github.com/evodicka/goccupied-client/handler"
	"github.com/evodicka/goccupied-client/io"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigchan := make(chan os.Signal, 10)
	signal.Notify(sigchan, os.Interrupt, os.Kill, syscall.SIGTERM)

	handler.Init()

	config := io.ReadConfig()
	if config.Global.Mode == io.MQTT {
		client.Connect()
		client.Subscribe(handler.SwitchState)
	} else {
		panic("Only mode mqtt is currently supported")
	}

	<-sigchan
	client.Disconnect()
	handler.Close()
}
