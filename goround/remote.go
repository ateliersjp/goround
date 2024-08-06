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

	if !SOCK {
		remoteAddr, _, _ = strings.Cut(conn.(remoteConn).RemoteAddr().String(), ":")
	}

	if HTTP && REAL_IP_HEADER != "" {
		if m, err := http.ReadMsg(conn); err == nil {
			for _, line := range m.Headers {
				if strutil.HasPrefixFold(line, REAL_IP_HEADER) {
					remoteAddr = strings.TrimSpace(line[16:])
					break
				}
			}
			connReader = m.Reader()
		}
	}

	return
}
