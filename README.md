# goround
A load balancer that assigns a server on a cluster by hashing the ```$remote_addr```, regularly reassigned. This leads to distributed loads, as well as distributed IPs during scraping.

## Usage
```
  -h	HTTP mode
  -u	UDP mode
  -s <endpoint>
  	UNIX socket mode
  -p <port>
  	listening port (default 80)
  -m <seconds>
  	reassignment interval (default 5)
```
### Environment variables
```
  LS_CMD	program that lists all servers on a cluster
  IP_HEADER	in HTTP mode, real IP header (default "X-Real-IP")
```
