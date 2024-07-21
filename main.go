package main

import (
	"fmt"
	_ "grpc-demo/gen/protoc"
	demo "grpc-demo/gen/protoc"
)

func main() {
	var x demo.BaseItem = demo.BaseItem{
		FirstNums: 123,
	}
	fmt.Printf("%+v\n", x)
}
