package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"remotetest/test"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "123.57.211.11:9005", "the address to connect to")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := test.NewTestClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.GetTestResult(ctx, &test.AmendableTest{TestNumber: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(response)
}
