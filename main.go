package main

import (
	"fmt"
)

func main() {

	tr := NewTCPTransport("localhost:4000")

    fmt.Println("We, Gucci!")
};