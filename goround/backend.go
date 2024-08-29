package main

import (
	"bytes"
	"strings"
	"net/netip"
	"os/exec"
	"time"

	. "github.com/rryqszq4/go-murmurhash"
)

var backendManager = &backendCache{}

type backendCache struct {
	backends    []string
	expiry      time.Time
}

func (c *backendCache) getAll() (backends []string) {
	if time.Now().After(c.expiry) {
		cmd := exec.Command(LS_CMD)
		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		if err := cmd.Run(); err == nil {
			for _, backend := range strings.Split(stdout.String(), "\n") {
				backend = strings.TrimSpace(backend)
				if len(backend) == 0 {
					continue
				}
				if addr, err := netip.ParseAddrPort(backend); err == nil {
					backend = addr.String()
				} else if addr, err := netip.ParseAddr(backend); err == nil {
					backend = netip.AddrPortFrom(addr, uint16(PORT)).String()
				} else {
					continue
				}
				backends = append(backends, backend)
			}

			c.backends = backends
			c.expiry = time.Now().Add(time.Second)
		}
	}

	return c.backends
}

func (c *backendCache) getByHash(key []byte) (backend string) {
	backends := c.getAll()
	if size := uint64(len(backends)); size > 0 {
		salt := uint64(time.Now().Unix()) % INTERVAL
		backend = backends[int(MurmurHash64A(key, salt) % size)]
	}
	return
}
