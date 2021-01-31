package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

func main() {

	hostname := flag.String("hostname", "", "hostname to test")
	startPort := flag.Int("start-port", 80, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 100, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")
	flag.Parse()

	ports := []int{}
	fmt.Println("*********** TCP SCANNER **************************")

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for port := *startPort; port < *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}

			wg.Done()
		}(port)

		wg.Wait()
		fmt.Printf("opened ports: %v\n", ports)

	}

}
