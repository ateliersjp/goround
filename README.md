# goround
Go implementation of a lightweight load balancer that tunnels to one of the replicated backends chosen based on hashing the remote IP, using the UNIX time as the salt. This way has the advantage of a distributed load on each backend, and especially for scraping use, distributed IP addresses.

## Usage
```
  -c <command>
    	program that lists all replicated backends
    	(required)
  -p <port>
    	listening and default backend port
    	(default 80)
  -u	UDP mode
  -h	HTTP mode
  -H <header>
    	HTTP mode with real IP header
    	(default "X-Real-IP")
  -s <filename>
    	UNIX socket mode
    	(default "/var/run/goround/goround.sock")
  -S	UNIX socket mode with the default filename
```
