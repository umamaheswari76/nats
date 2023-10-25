package main

import (
	"fmt"
	"log"

	// natsServer "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)


func main() {

	nc, err := nats.Connect("nats://demo.nats.io:4222")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()


	var kv nats.KeyValue
	js, _ := nc.JetStream()
	if stream, _ := js.StreamInfo("KV_discovery"); stream == nil {
		kv, _ = js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket: "discovery",
		})
	} else {
		kv, _ = js.KeyValue("discovery")
	}

	kv.Put("services.orders", []byte("https://localhost:8080/orders"))
	fmt.Println("stored key value in bucket")
	entry, _ := kv.Get("services.orders")
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

}
