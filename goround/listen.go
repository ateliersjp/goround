package main

import (
	"os"
	"net"
)

func listen() (ln net.Listener, err error) {
	if SOCK {
		os.Remove(ADDRESS)
		ln, err = net.Listen("unix", ADDRESS)
		if err == nil {
			os.Chmod(ADDRESS, 0666)
		}
	} else {
		ln, err = net.Listen(PROTOCOL, ADDRESS)
	}
	return
}
