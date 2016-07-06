package goutility

import (
	"errors"
	"net"
)

// GetFirstNonLoopbackIP get first non-loopback ip address of the current machine
// https://gist.github.com/mowings/017c80c188d1024ba3e7
func GetFirstNonLoopbackIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip == nil || ip.IsLoopback() {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}
		return ip.String(), nil
	}
	return "", errors.New("Unable to find external ip address")
}
