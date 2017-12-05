package openstack

import (
	"bytes"
	"net"
)

type IPRange struct {
	from net.IP
	to   net.IP
}

func NewIPRange(from string, to string) IPRange {
	return IPRange{net.ParseIP(from).To4(), net.ParseIP(to).To4()}
}

var privateIPRanges []IPRange = []IPRange{
	NewIPRange("10.0.0.0", "10.255.255.255"),
	NewIPRange("172.16.0.0", "172.31.255.255"),
	NewIPRange("192.168.0.0", "192.168.255.255"),
}

func IsPrivateIP(ip net.IP) bool {
	for _, r := range privateIPRanges {
		if bytes.Compare(ip, r.from) >= 0 && bytes.Compare(ip, r.to) <= 0 {
			return true
		}
	}
	return false
}
