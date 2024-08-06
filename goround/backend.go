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
		cmd := exec.Command(LS_PROGRAM)
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
	if size := len(backends); size > 0 {
		now := time.Now().Unix()
		i := MurmurHash2(key, uint32(now)) % uint32(size)
		backend = backends[int(i)]
	}
	return
}
