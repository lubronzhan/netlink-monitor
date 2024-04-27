package main

import (
	"fmt"
	"syscall"

	"github.com/lubronzhan/netlink-monitor/pkg/syscall_lib"
)

func main() {
	l, _ := ListenNetlink()

	for {
		msgs, err := l.ReadMsgs()
		if err != nil {
			fmt.Printf("could not read netlink: %v", err)
		}

		for _, m := range msgs {
			fmt.Print("pid: ")
			fmt.Println(m.Header.Pid)
			fmt.Println(m.Data)
			fmt.Println(syscall_lib.GetTypeName(m.Header.Type))
			if IsNewAddr(&m) {
				fmt.Println("New Addr")
			}

			if IsDelAddr(&m) {
				fmt.Println("Del Addr")
			}
		}
	}
}

type NetlinkListener struct {
	fd int
	sa *syscall.SockaddrNetlink
}

func ListenNetlink() (*NetlinkListener, error) {
	groups := (1 << (syscall.RTNLGRP_LINK - 1)) |
		(1 << (syscall.RTNLGRP_IPV4_IFADDR - 1)) |
		(1 << (syscall.RTNLGRP_IPV6_IFADDR - 1))

	s, err := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM,
		syscall.NETLINK_ROUTE)
	if err != nil {
		return nil, fmt.Errorf("socket: %s", err)
	}

	saddr := &syscall.SockaddrNetlink{
		Family: syscall.AF_NETLINK,
		Pid:    uint32(0),
		Groups: uint32(groups),
	}

	err = syscall.Bind(s, saddr)
	if err != nil {
		return nil, fmt.Errorf("bind: %s", err)
	}

	return &NetlinkListener{fd: s, sa: saddr}, nil
}

func (l *NetlinkListener) ReadMsgs() ([]syscall.NetlinkMessage, error) {
	defer func() {
		recover()
	}()

	pkt := make([]byte, 2048)

	n, err := syscall.Read(l.fd, pkt)
	if err != nil {
		return nil, fmt.Errorf("read: %s", err)
	}

	msgs, err := syscall.ParseNetlinkMessage(pkt[:n])
	if err != nil {
		return nil, fmt.Errorf("parse: %s", err)
	}

	return msgs, nil
}

func IsNewAddr(msg *syscall.NetlinkMessage) bool {
	return msg.Header.Type == syscall.RTM_NEWADDR
}

func IsDelAddr(msg *syscall.NetlinkMessage) bool {
	return msg.Header.Type == syscall.RTM_DELADDR
}

// rtm_scope is the distance to the destination:
//
// RT_SCOPE_UNIVERSE   global route
// RT_SCOPE_SITE       interior route in the
// local autonomous system
// RT_SCOPE_LINK       route on this link
// RT_SCOPE_HOST       route on the local host
// RT_SCOPE_NOWHERE    destination doesn't exist
//
// The values between RT_SCOPE_UNIVERSE and RT_SCOPE_SITE are
// available to the user.

func IsRelevant(msg *syscall.IfAddrmsg) bool {
	if msg.Scope == syscall.RT_SCOPE_UNIVERSE ||
		msg.Scope == syscall.RT_SCOPE_SITE {
		return true
	}

	return false
}
