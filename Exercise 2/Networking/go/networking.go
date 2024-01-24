package main

import (
	"fmt"
	"net"
	"runtime"
)

func writeToServerUDP() {
	raddr, err := net.ResolveUDPAddr("udp", "10.100.23.186:20007")

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	message := []byte("Test group 8")
	conn.Write(message)

	defer conn.Close()
}

func readfromServerUDP() {
	raddr, err := net.ResolveUDPAddr("udp", ":20007")

	msg, err := net.ListenUDP("udp", raddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	

	buffer := make([]byte, 1024)

	msg.Read(buffer[0:])

	fmt.Print("From server: ", string(buffer[0:]))
	defer msg.Close()
}

func writeToServerTCP(){
	addr, err := net.ResolveTCPAddr("tcp", "10.100.23.186:33546") //change port to 33546 for \0 messages

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	

	message := []byte("Connect to: 10.100.23.18:33546\000") //change port to 33546 for \0 messages
	conn.Write(message)

	//message2 := []byte("TCP test 8\000")
	//conn.Write(message2)
	


	

	conn.Close()
}

func readfromServerTCP() {
	raddr, err := net.ResolveTCPAddr("tcp", "10.100.23.186:34933") 
	msg, err := net.ListenTCP("tcp", raddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	

	buffer := make([]byte, 1024)

	//msg.Read(buffer[0:])

	fmt.Print("From server: ", string(buffer[0:]))
	defer msg.Close()
}





func main() {
	runtime.GOMAXPROCS(2)

	go writeToServerTCP()
	//go readfromServerTCP()

	select {}

}
