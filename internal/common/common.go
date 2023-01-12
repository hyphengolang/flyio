package common

import (
	"net"
	"os"
)

func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func GetIPAddress() string {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
