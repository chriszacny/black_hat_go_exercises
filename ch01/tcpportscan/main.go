package main

import (
	"fmt"
	"net"
	"sort"
)

const hostName="scanme.nmap.org"
const protocol="tcp"
const numPortsToScan=1024
const numWorkers=100

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", hostName, p)
		conn, err := net.Dial(protocol, address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, numWorkers)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func(){
		for i := 1; i <= numPortsToScan; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= numPortsToScan; i++ {
		portResult := <- results
		if portResult != 0 {
			openports = append(openports, portResult)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open.\n", port)
	}
}