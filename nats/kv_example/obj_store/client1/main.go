package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

type ObjectMeta struct {
	Name string
}

var metadata nats.ObjectMeta

// metadata = &ObjectMeta{Name: "testfile"}

func main() {
	//"nats://demo.nats.io:4222"
	nc, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal("can't connect ", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("can't connect to js ", err)
	}

	filePath := "/home/vasenth/Documents/nats/kv_example/kv_grpc/test3.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("can't open local file: ", err)
	}
	defer file.Close()

	reader := io.Reader(file)
	metadata = nats.ObjectMeta{
		Name: "newfile3",
	}

	var obj nats.ObjectStore
	if stream, _ := js.StreamInfo("testos"); stream == nil {
		obj, _ = js.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket: "newos",
		})
		fmt.Println("1")
	} else {
		obj, _ = js.ObjectStore("newos")
		fmt.Println("2")
	}

	// var info1 *nats.ObjectInfo
	// var info1 *nats.ObjectInfo
	
	info, err := obj.Put(&metadata, reader)
	if err != nil {
		log.Println("Can't put file in object store: ", err)
	}
	log.Println("obj info: ", info.ObjectMeta)

	res, err := obj.Get("newfile3")
	if err != nil {
		log.Println("error getting file in object store: ", err)
	}

	inf, err := res.Info()
	if err != nil {
		fmt.Println("can't show info: ", err)
	}
	fmt.Println("file got info: ", inf)

	buffer := make([]byte, 60)
	batchNumber := 1
	for {
		num, err := res.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error reading file: ", err)
		}
		chunk := buffer[:num]
		// fmt.Print(string(buffer[:bytes]))
		log.Printf("Sent - batch #%v - size - %v\ncontent - %v\n\n", batchNumber, len(chunk), string(chunk))
		batchNumber += 1
	}

	// fmt.Print(string(buffer[:bytes]))
	// fmt.Println("got file: ",res.Read())

}


