package main

import (
	// "context"

	"context"
	"fmt"
	"log"
	"time"

	// "time"

	natsServer "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func createServer() {
	opts := &natsServer.Options{
		ServerName:     "local_nats_server",
		Host:           "localhost",
		Port:           15000,
		NoLog:          false,
		NoSigs:         false,
		MaxControlLine: 4096,
		MaxPayload:     65536,
	}

	server, err := natsServer.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	err = natsServer.Run(server)
	if err != nil {
		log.Fatal("Failed to start NATS server:", err)
	}

	log.Println("NATS server started")

}

func producer(ctx context.Context) {
	nc, err := nats.Connect("nats://127.0.0.1:15000")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()
	message := "hello world"
	subject := "logs"
	err = nc.Publish(subject, []byte(message))
	if err != nil {
		log.Println("Failed to publish message:", err)
	} else {
		log.Println("message published ", message)
	}
}

func subscriber(ctx context.Context) {
	nc, err := nats.Connect("nats://127.0.0.1:15000")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	fmt.Println("Connected to NATS server on port 15000")

	subject := "logs"

	_, err1 := nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("message received on subject: %v, data: %v\n", msg.Subject, string(msg.Data))
	})
	if err1 != nil {
		fmt.Println("Failed to subscribe to NATS server ", err1)
	}
	// producer()

	time.Sleep(1 * time.Minute)
	// sub.Unsubscribe()

}

func main() {
	ctx, _ := context.WithCancel((context.Background()))

	createServer()
	go subscriber(ctx)
	time.Sleep(2 * time.Second)
	
	go producer(ctx)
	time.Sleep(2 * time.Second)
}
