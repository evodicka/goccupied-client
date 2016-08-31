package main

import (
	"github.com/evodicka/goccupied-client/client"
	"github.com/evodicka/goccupied-client/handler"
	"github.com/evodicka/goccupied-client/io"
	"github.com/evodicka/goccupied-server/scheduler"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	handler.Init()

	config := io.ReadConfig()
	if config.Global.Mode == io.MQTT {
		client.Connect()
		client.Subscribe(handler.SwitchState)
	} else if config.Global.Mode == io.HTTP {
		scheduler.StartPeriod(handler.HandlePoll, config.Http.PollIntervalSeconds)
	} else {
		panic("Only modes mqtt and http are currently supported")
	}
	shutdown()
}

func shutdown() {
	sigchan := make(chan os.Signal, 10)
	signal.Notify(sigchan, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sigchan
	scheduler.Stop()
	client.Disconnect()
	handler.Close()
}
