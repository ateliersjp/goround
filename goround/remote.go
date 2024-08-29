package main

import (
	"io"
	"net"
	"strings"

	"go4.org/strutil"
	"github.com/ateliersjp/http"
)

type remoteConn interface {
	RemoteAddr() net.Addr
}

func getRemoteAddr(conn net.Conn) (connReader io.Reader, remoteAddr string) {
	connReader = conn

	if SOCK == "" {
		remoteAddr = conn.(remoteConn).RemoteAddr().String()
	}

	if HTTP {
		if m, err := http.ReadMsg(conn); err == nil {
			for _, line := range m.Headers {
				if strutil.HasPrefixFold(line, IP_HEADER) {
					remoteAddr = strings.TrimLeft(line[len(IP_HEADER):], ": ")
					break
				}
			}
			connReader = m.Reader()
		}
	}

	if remoteAddr == "" {
		remoteAddr = "127.0.0.1"
	}

	return
}
