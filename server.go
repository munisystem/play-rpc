package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Arith int

func (a *Arith) Fib(n int, reply *int) {
	*reply = FibTail(n)
}

func FibTail(n int, a ...int) int {
	if len(a) == 0 {
		a = []int{1, 0}
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return a[0]
	}

	return FibTail(n-1, a[0]+a[1], a[0])
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":6262", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
