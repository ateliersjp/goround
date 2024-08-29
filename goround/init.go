package main

import (
	"os"
	"flag"
	"fmt"
	"log"
)

var (
	PROTOCOL = "tcp"
	ADDRESS string
	LS_CMD string
	PORT uint64
	INTERVAL uint64
	UDP bool
	HTTP bool
	SOCK string
	IP_HEADER string
)

func init() {
	flag.Uint64Var(&PORT, "p", 80, "listening port")
	flag.Uint64Var(&INTERVAL, "m", 5, "reassignment interval, seconds")
	flag.BoolVar(&UDP, "u", false, "UDP mode")
	flag.BoolVar(&HTTP, "h", false, "HTTP mode")
	flag.StringVar(&SOCK, "s", "", "UNIX socket mode, endpoint")
}

func parse() {
	flag.Parse()

	LS_CMD = os.Getenv("LS_CMD")
	if LS_CMD == "" {
		log.Fatal("$LS_CMD is required")
	}

	if SOCK == "" {
		ADDRESS = fmt.Sprintf(":%d", PORT)
	}

	if HTTP {
		IP_HEADER = os.Getenv("IP_HEADER")
	}

	if UDP {
		PROTOCOL = "udp"
	}
}
