package main

import (
	"flag"
	"fmt"
	"net/rpc"

	//	"time"
	"net"
)

var nextAddr string
var n int
var conn *rpc.Client

type runInfo struct {
	ID      int
	bottles int
}

type Message struct {
	text string
}

type nextHandler struct{}

func (s *nextHandler) next(info runInfo, message *Message) (err error) {
	var identity int

	if n > -1 {
		identity = 1
	} else {
		identity = info.ID
	}

	if info.bottles > 0 {
		fmt.Println("Buddy ", identity, ": ", info.bottles, "bottles of beer on the wall, ", info.bottles, "bottles of beer. take one down, pass it around...")
		conn.Call("nextHandler.next", runInfo{ID: identity + 1, bottles: info.bottles - 1}, Message{text: "success"})
	} else {
		fmt.Println("Buddy ", identity, "NO MORE BOTTLES OF BEEER WOOOOOOOO")
	}

	return
}

func main() {
	//	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	flag.IntVar(&n, "bottles", -1, "number of bottles to count down from, set if this is the start of the counting")
	//	bottles := flag.Int("n",0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.
	rpc.Register(&nextHandler{})
	listener, _ := net.Listen("tcp", ":8040")
	defer listener.Close()

	if n > -1 {
		conn, _ = rpc.Dial("tcp", nextAddr)
		defer conn.Close()
		rpc.Accept(listener)
		fmt.Println("Buddy 1: ", n, "bottles of beer on the wall, ", n, "bottles of beer. take one down, pass it around...")
		conn.Call("nextHandler.next", runInfo{ID: 2, bottles: n - 1}, Message{text: "yeay"})
	} else {
		rpc.Accept(listener)
		conn, _ = rpc.Dial("tcp", nextAddr)
		defer conn.Close()
	}

}
