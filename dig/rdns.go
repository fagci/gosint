package dig

import "net"

func ReverseDNS(ip string) ([]string, error) {
	return net.LookupAddr(ip)
}
