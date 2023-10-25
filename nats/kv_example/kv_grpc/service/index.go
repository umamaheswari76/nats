package service

import (
	"context"
	"fmt"
	pro "kv_grpc/proto"
	"log"

	"github.com/nats-io/nats.go"
)

type Server struct {
	pro.UnimplementedKvGrpcServiceServer
}

func (s *Server) GetFile(ctx context.Context, req *pro.GetFileRequest) (*pro.GetFileResponse, error) {

	// fmt.Println("going to connect nats")
	// natsServer := config.NatsConnection()
	// fmt.Println("nats conneected")

	// var kv nats.KeyValue
	// fmt.Println("1")

	// if stream, _ := natsServer.StreamInfo("KV_discovery"); stream == nil {
	// 	kv, _ = natsServer.CreateKeyValue(&nats.KeyValueConfig{
	// 		Bucket: "discovery",
	// 	})
	// 	fmt.Println("2")
	// } else {
	// 	kv, _ = natsServer.KeyValue("discovery")
	// 	fmt.Println("3")
	// }
	// fmt.Println("4")
	// fmt.Println(req.FileName)

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
	// fmt.Println("1")
	if stream, _ := js.StreamInfo("test"); stream == nil {
		kv, _ = js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket: "new",
		})
		// fmt.Println("2")
	}else{
		kv,_=js.KeyValue("new")
	}
	
	id := "reggt1"
	// id1:= "1"
	// id2 := "2" 
	:= kv.Put(id, []byte(req.FileName))
	fmt.Println("stored key value in bucket")

	// entry, err := kv.Get("services.orders")

	return &pro.GetFileResponse{
		FileID: id,
	}, nil

}
