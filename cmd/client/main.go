package main

import (
	"fmt"
	"log"

	"github.com/nguyenhoaibao/logsvc/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewLogServiceClient(conn)

	req := &pb.Log{
		ClientIp: "xxx",
		ServerIp: "xxx",
		Tags: &pb.Tags{
			Tags: map[string]string{
				"key": "value",
			},
		},
		Msg: "ahihi",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err = client.Write(ctx, req)
	if err != nil {
		fmt.Printf("could not write log: %v", err)
	}

	resp, err := client.Get(ctx, &pb.GetRequest{
		// ClientIp: "xxx",
		Tags: &pb.Tags{
			Tags: map[string]string{
				"key": "value",
			},
		},
	})
	if err != nil {
		fmt.Printf("could not get logs: %v", err)
	}
	fmt.Println(resp)
}
