package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	PROTOCOL = "tcp"
	ADDRESS string
	PORT uint
	UDP bool
	HTTP bool
	SOCK bool
	REAL_IP_HEADER string
	LS_PROGRAM string
)

func init() {
	flag.UintVar(&PORT, "p", 80, "listening and default backend port")
	flag.BoolVar(&UDP, "u", false, "UDP mode")
	flag.BoolVar(&HTTP, "h", false, "HTTP mode")
	flag.StringVar(&REAL_IP_HEADER, "H", "X-Real-IP", "HTTP mode with real IP header")
	flag.BoolVar(&SOCK, "S", false, "UNIX socket mode with the default filename")
	flag.StringVar(&ADDRESS, "s", "/var/run/goround/goround.sock", "UNIX socket mode")
	flag.StringVar(&LS_PROGRAM, "c", "", "program that lists all replicated backends")
}

func parse() {
	flag.Parse()

	if LS_PROGRAM == "" {
		log.Fatal("-c option is mandatory")
	}

	if REAL_IP_HEADER == "" {
		if SOCK {
			log.Fatal("-s option must be set together with -H option")
		}
	
		if ADDRESS != "" {
			log.Fatal("-S option must be set together with -H option")
		}
	}

	if REAL_IP_HEADER != "" {
		HTTP = true
	}

	if ADDRESS != "" {
		SOCK = true
	} else {
		if SOCK {
			ADDRESS = "/var/run/goround/goround.sock"
		} else {
			ADDRESS = fmt.Sprintf(":%d", PORT)
		}
	}

	if UDP {
		PROTOCOL = "udp"
	}
}
