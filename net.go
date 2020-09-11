package mypkg

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// IPGet Get the IP address of the physical network adapter, ignore the virtual machine network adapter
func IPGet() map[string]string {
	IPMap := make(map[string]string)
	nfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, nface := range nfaces {
		addrs, _ := nface.Addrs()
		name := nface.Name
		up := nface.Flags&net.FlagUp != 0
		loopback := nface.Flags&net.FlagLoopback != 0
		// unicast := nface.Flags&net.FlagPointToPoint != 0
		// Multicast := nface.Flags&net.FlagMulticast != 0
		// broadcast := nface.Flags&net.FlagBroadcast != 0

		// Not a Loopback address
		if up && !loopback && !isVMAdapter(name) {
			for _, addr := range addrs {
				ipnet := addr.(*net.IPNet)

				// get net.IP in *Net.IPNet
				// ip := ipnet.IP

				// convert ip net.IP([]byte) to ip string
				// ipstring := ipnet.IP.String()

				// convert ip string to ip net.IP([]byte)
				// ip = net.ParseIP(ipstring)

				if ipnet.IP.To4() != nil {
					IPMap["ipv4"] = ipnet.IP.String()
				} else if ipnet.IP.To16() != nil {
					IPMap["ipv6"] = ipnet.IP.String()
				}
			}
		}
		// Loopback address
		if up && loopback && !isVMAdapter(name) {
			for _, addr := range addrs {
				ipnet := addr.(*net.IPNet)
				if ipnet.IP.To4() != nil {
					IPMap["loopbackv4"] = ipnet.IP.String()
				} else if ipnet.IP.To16() != nil {
					IPMap["loopbackv6"] = ipnet.IP.String()
				}
			}
		}
	}
	return IPMap
}

// IPIsVMAdapter check if string is in Virtual Machines Network Adapter slice
func isVMAdapter(adapterName string) bool {
	vms := []string{
		"VMware Network Adapter",
		"VirtualBox",
	}
	for _, vm := range vms {
		if strings.Contains(adapterName, vm) {
			return true
		}
	}
	return false
}
