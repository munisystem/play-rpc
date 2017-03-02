package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:6262")
	if err != nil {
		log.Fatal(err)
	}

	var reply int
	num := 50
	err = client.Call("Arith.Fib", num, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the %dth Fibonacci is %d", num, reply)
}
