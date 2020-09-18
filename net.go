package gopkg

import (
	"log"
	"net"
	"strings"
)

// IPGet Get the IP address of the physical network adapter, ignore the virtual machine network adapter
func IPGet() map[string]string {
	IPMap := make(map[string]string)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, netInterface := range netInterfaces {
		addresses, _ := netInterface.Addrs()
		name := netInterface.Name
		up := netInterface.Flags&net.FlagUp != 0
		loopback := netInterface.Flags&net.FlagLoopback != 0
		// unicast := netInterface.Flags&net.FlagPointToPoint != 0
		// Multicast := netInterface.Flags&net.FlagMulticast != 0
		// broadcast := netInterface.Flags&net.FlagBroadcast != 0

		// Not a Loopback address
		if up && !loopback && !isVMAdapter(name) {
			for _, addr := range addresses {
				ipNet := addr.(*net.IPNet)

				// get net.IP in *Net.IPNet
				// ip := ipNet.IP

				// convert ip net.IP([]byte) to ip string
				// ipString := ipNet.IP.String()

				// convert ip string to ip net.IP([]byte)
				// ip = net.ParseIP(ipString)

				if ipNet.IP.To4() != nil {
					IPMap["ipv4"] = ipNet.IP.String()
				} else if ipNet.IP.To16() != nil {
					IPMap["ipv6"] = ipNet.IP.String()
				}
			}
		}
		// Loopback address
		if up && loopback && !isVMAdapter(name) {
			for _, addr := range addresses {
				ipNet := addr.(*net.IPNet)
				if ipNet.IP.To4() != nil {
					IPMap["loopbackv4"] = ipNet.IP.String()
				} else if ipNet.IP.To16() != nil {
					IPMap["loopbackv6"] = ipNet.IP.String()
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
