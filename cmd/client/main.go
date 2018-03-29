package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nguyenhoaibao/logsvc/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", "127.0.0.1:8080", "Server addr")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewLogServiceClient(conn)

	var (
		clientIP = "127.0.0.1"
		serverIP = "127.0.1.1"
		tags     = map[string]string{"key": "val"}
		msg      = "log message"
	)

	req := &pb.Log{
		ClientIp: clientIP,
		ServerIp: serverIP,
		Tags: &pb.Tags{
			Tags: tags,
		},
		Msg: msg,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err = client.Write(ctx, req)
	if err != nil {
		fmt.Printf("could not write log: %v", err)
	}

	resp, err := client.Get(ctx, &pb.GetRequest{
		Tags: &pb.Tags{
			Tags: tags,
		},
	})
	if err != nil {
		fmt.Printf("could not get logs: %v", err)
	}
	fmt.Println(resp)
}
