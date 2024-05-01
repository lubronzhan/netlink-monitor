package main

import (
	"fmt"
	"log"

	"github.com/vishvananda/netlink"
)

func main() {
	// Subscribe to netlink messages
	ch := make(chan netlink.LinkUpdate)
	done := make(chan struct{})
	if err := netlink.LinkSubscribe(ch, done); err != nil {
		log.Fatalf("failed to subscribe to netlink messages: %v", err)
	}

	// Handle incoming netlink messages
	var msg netlink.LinkUpdate
	for {
		msg = <-ch
		fmt.Print("msg: ")
		fmt.Println(msg)
		fmt.Print("msg.Header: ")
		fmt.Println(msg.Header)
		fmt.Print("msg.Header.Pid: ")
		fmt.Println(msg.Header.Pid)
		fmt.Print("Serialize(): ")
		fmt.Println(msg.Serialize())
	}
}
