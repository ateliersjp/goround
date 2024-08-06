package main

import (
	"log"
	"net"
)

func main() {
	parse()
	ln, err := listen()
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	connReader, remoteAddr := getRemoteAddr(conn)
	if backend := backendManager.getByHash([]byte(remoteAddr)); backend != "" {
		if remote, err := net.Dial(PROTOCOL, backend); err == nil {
			defer remote.Close()
			wg := NewWaitGroup()
			go wg.Copy(remote, connReader)
			go wg.Copy(conn, remote)
			wg.Wait()
		}
	}
}