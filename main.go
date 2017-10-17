package main

import (
	"fmt"
	"os"
)

const header = "Content-type: text/plain\nauthor: erdii\n";

func main() {
	// read ip from environment variable
	addr := os.Getenv("REMOTE_ADDR")

	// default ip to "127.0.0.1"
	if len(addr) == 0 {
		addr = "127.0.0.1"
	}

	// send http response
	fmt.Printf("%s\n%s", header, addr)
}
