package goutility

import (
	"errors"
	"net"
)

// GetFirstNonLoopbackIP get first non-loopback ip address of the current machine
func GetFirstNonLoopbackIP() (string, error) {

	// https://gist.github.com/mowings/017c80c188d1024ba3e7
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {

		ip := GetAvailableIPFromAddress(addr)
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

// GetAvailableIPFromAddress retrieves the IP from a net.Addr if one is available
func GetAvailableIPFromAddress(address net.Addr) net.IP {

	var ip net.IP
	switch v := address.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	return ip
}
