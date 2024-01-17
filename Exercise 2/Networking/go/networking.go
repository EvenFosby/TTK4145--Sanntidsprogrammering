package main

import (
	"fmt"
	"net"
	"runtime"
)


func writeToServer(){
	raddr, err := net.ResolveUDPAddr("udp", "10.100.23.129:20008")

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	message := []byte("Serriøst har vi brukt fire tima på det herre!?")
	conn.Write(message)

	defer conn.Close()
}


func readfromServer() {
	raddr, err := net.ResolveUDPAddr("udp", "10.100.23.129:20008")

	msg, err := net.ListenUDP("udp", raddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	buffer := make([]byte, 1024)

	msg.Read(buffer[0:])

	fmt.Print("From server: ", string(buffer[0:]))

}

func main() {
	runtime.GOMAXPROCS(2)

	go writeToServer()
	go readfromServer()
	

}
	


