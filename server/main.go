package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"
)

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func highLoadSimulate(stop chan bool) int {
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case <-stop:
					return
				default:
				}
			}
		}()
	}

	return n
}

func main() {
	machineIP := getOutboundIP().String()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		stop := make(chan bool)
		numOfCore := highLoadSimulate(stop)
		time.Sleep(time.Second * 2)
		close(stop)

		rw.WriteHeader(http.StatusOK)
		msg := fmt.Sprintf("Hello from %s. Running %d core of cpu", machineIP, numOfCore)
		rw.Write([]byte(msg))
	})
	log.Println("Server is starting at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
