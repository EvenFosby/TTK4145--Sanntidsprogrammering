package main

import(
	"fmt"
	"net"
	//"os"
)

var addr = "123"

func main(){
	fmt.Println()
	net.Dial("udp",addr)
}
