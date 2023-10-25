package config

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NatsConnection() nats.JetStreamContext{
	nc, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil{
		log.Fatal("Failed to connect jetstream ",err)
	}
	return js
}
