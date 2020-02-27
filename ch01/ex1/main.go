package main

import (
	"fmt"
	"net"
	"sync"
)

//const Address = "scanme.nmap.org"
const Address = "127.0.0.1"
const Ports = 65535

func main () {
	var wg sync.WaitGroup
	for i := 1; i <= Ports; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%v:%d", Address, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed
				return
			}
			conn.Close()
			fmt.Printf("%d is open.\n", j)
		}(i)
	}
	wg.Wait()
}