package main

import (
	"fmt"

	pb "test/test_go_grpc"
)

func main() {
	fmt.Println("hello world")
	test := pb.GetMessageRequest{}
	fmt.Println(test)
}
