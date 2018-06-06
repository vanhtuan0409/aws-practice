package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
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

func main() {
	machineIP := getOutboundIP().String()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		msg := fmt.Sprintf("Hello from %s", machineIP)
		rw.Write([]byte(msg))
	})
	log.Println("Server is starting at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
