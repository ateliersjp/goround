package main

import (
	"os"
	"net"
	"fmt"
	"strings"
)

func listen() (ln net.Listener, err error) {
	if SOCK != "" {
		if i := strings.LastIndex(SOCK, "/"); i >= 0 {
			os.MkdirAll(SOCK[:i], 0777)
		} else {
			SOCK = fmt.Sprintf("/var/run/%s", SOCK)
		}
		os.Remove(SOCK)
		ln, err = net.Listen("unix", SOCK)
		if err == nil {
			os.Chmod(SOCK, 0666)
		}
	} else {
		ln, err = net.Listen(PROTOCOL, ADDRESS)
	}
	return
}
