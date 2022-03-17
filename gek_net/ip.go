package gek_net

import (
	"net/netip"
)

// IsLoopback IP是否为回环地址
func IsLoopback(ip string) (isLoopback bool, err error) {
	var blocks = []string{
		// Loopback
		"127.0.0.0/8",
		"::1/128",
	}
	// 解析ip
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false, err
	}
	// ip是否包含在块中
	for _, blk := range blocks {
		prefix, err := netip.ParsePrefix(blk)
		if err != nil {
			return false, err
		}
		if prefix.Contains(addr) {
			return true, nil
		}
	}

	return false, nil
}

// IsPrivate IP是否为专用地址
func IsPrivate(ip string) (isPrivate bool, err error) {
	var blocks = []string{
		// Private network
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"fc00::/7",
	}
	// 解析ip
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false, err
	}
	// ip是否包含在块中
	for _, blk := range blocks {
		prefix, err := netip.ParsePrefix(blk)
		if err != nil {
			return false, err
		}
		if prefix.Contains(addr) {
			return true, nil
		}
	}

	return false, nil
}

// IsLinkLocal IP是否为链路本地地址
func IsLinkLocal(ip string) (isLinkLocal bool, err error) {
	var blocks = []string{
		// Link-local address
		"169.254.0.0/16",
		"fe80::/10",
	}
	// 解析ip
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false, err
	}
	// ip是否包含在块中
	for _, blk := range blocks {
		prefix, err := netip.ParsePrefix(blk)
		if err != nil {
			return false, err
		}
		if prefix.Contains(addr) {
			return true, nil
		}
	}

	return false, nil
}

// IsPublic IP是否为公网地址
func IsPublic(ip string) (isPublic bool, err error) {
	isLoopback, err := IsLoopback(ip)
	if err != nil {
		return false, err
	}

	if isLoopback {
		return false, nil
	}

	isPrivate, err := IsPrivate(ip)
	if err != nil {
		return false, err
	}

	if isPrivate {
		return false, nil
	}

	isLinkLocal, err := IsLinkLocal(ip)
	if err != nil {
		return false, err
	}

	if isLinkLocal {
		return false, nil
	}

	return true, nil
}
