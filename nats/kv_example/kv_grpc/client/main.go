package main

import (
	"context"
	"fmt"
	"log"

	pro "kv_grpc/proto"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Can't dial grpc client: ", err)
	}

	client := pro.NewKvGrpcServiceClient(conn)
	getFileRequest := &pro.GetFileRequest{
		FileName: "hello1",
	}

	getFileResponse, err := client.GetFile(context.Background(), getFileRequest)
	if err != nil {
		log.Fatalf("GetFile error: %v", err)
	}
	log.Println(getFileResponse)

	// config.NatsConnection()
	nc, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("Failed to connect jetstream ", err)
	}
	var kv nats.KeyValue
	
	fmt.Println("1")
	if stream, _ := js.StreamInfo("test"); stream == nil {
		kv, _ = js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket: "new",
		})
		fmt.Println("2")
	} else {
		fmt.Println("no such a bucket!!!")
	}
	fmt.Println(getFileResponse.FileID)
	res, _ := kv.Get(getFileResponse.FileID)
	fmt.Println(string(res.Value()))
	
	res1, _ := kv.Keys()
	fmt.Println(res1)

}

// import (
// 	"log"
// 	pro "kv_example/kv_grpc/proto"

// 	"google.golang.org/grpc"
// )

// func main() {
// 	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Could not connect: %v", err)
// 	} else {
// 		log.Println("connected")
// 	}
// 	// func pro.NewFileTransferServiceClient(cc grpc.ClientConnInterface)
// 	// pro.FileTransferServiceClient

// 	client := pro.KvGrpcServiceClient(conn)
// 	getFileRequest := &pro.GetFileRequest{
// 		FileName: "sampleKVfile.txt",
// 	}

// }
