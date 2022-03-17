package gek_net

//
//import (
//	"fmt"
//	"net"
//	"net/http"
//	"strings"
//)
//
//// NetInterfaces 获取网络接口
//func NetInterfaces() ([]NetInterface, error) {
//
//	var netInterfaces []NetInterface
//
//	ifaces, err := net.Interfaces()
//	if err != nil {
//		return nil, err
//	}
//
//	for _, iface := range ifaces {
//		var netInterface NetInterface
//		Addrs, _ := iface.Addrs()
//		netInterface.Name = iface.Name
//		netInterface.MAC = iface.HardwareAddr.String()
//		netInterface.Flags.Up = iface.Flags&net.FlagUp != 0
//		netInterface.Flags.Unicast = iface.Flags&net.FlagPointToPoint != 0
//		netInterface.Flags.Broadcast = iface.Flags&net.FlagBroadcast != 0
//		netInterface.Flags.Multicast = iface.Flags&net.FlagMulticast != 0
//
//		for _, addr := range Addrs {
//			ip := addr.(*net.IPNet).IP
//			switch {
//			// loopback地址
//			case ip.IsLoopback():
//				if ip.To4() != nil {
//					netInterface.Addr.IPv4Loopback = ip
//				} else if ip.To16() != nil {
//					netInterface.Addr.IPv6Loopback = ip
//				}
//			// 本地私有网络地址
//			case isPrivateIP(ip):
//				if ip.To4() != nil {
//					netInterface.Addr.IPv4Private = ip
//				} else if ip.To16() != nil {
//					netInterface.Addr.IPv6UniqueLocal = ip
//				}
//			// 链路本地地址
//			case ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast():
//				if ip.To4() != nil {
//					netInterface.Addr.IPv4LinkLocal = ip
//				} else if ip.To16() != nil {
//					netInterface.Addr.IPv6LinkLocal = ip
//				}
//			// 全局单播地址
//			case ip.IsGlobalUnicast():
//				if ip.To4() != nil {
//					netInterface.Addr.IPv4Public = ip
//				} else if ip.To16() != nil {
//					netInterface.Addr.IPv6GlobalUnicast = ip
//				}
//			}
//		}
//		netInterfaces = append(netInterfaces, netInterface)
//	}
//	return netInterfaces, nil
//}
//
//// isPrivateIP IP是否属于私有IP块
//func isPrivateIP(ip net.IP) bool {
//	var privateIPBlocks []*net.IPNet
//	for _, cidr := range []string{
//		"127.0.0.0/8",    // IPv4 loopback
//		"10.0.0.0/8",     // RFC1918
//		"172.16.0.0/12",  // RFC1918
//		"192.168.0.0/16", // RFC1918
//		"169.254.0.0/16", // RFC3927 link-local
//		"::1/128",        // IPv6 loopback
//		"fe80::/10",      // IPv6 link-local
//		"fc00::/7",       // IPv6 unique local addr
//	} {
//		_, block, err := net.ParseCIDR(cidr)
//		if err != nil {
//			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
//		}
//		privateIPBlocks = append(privateIPBlocks, block)
//	}
//
//	// 去除链路本地地址
//	// ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast()
//	if ip.IsLoopback() {
//		return true
//	}
//
//	for _, block := range privateIPBlocks {
//		if block.Contains(ip) {
//			return true
//		}
//	}
//	return false
//}
//
//// GetPublicIP NetInterface方法获取公网地址
//func (n NetInterface) GetPublicIP(ipType string) net.IP {
//	// 如果 NetInterface 未启用,则直接返回空
//	if !n.Flags.Up {
//		return nil
//	}
//
//	switch ipType {
//	case "ipv4":
//		if n.Addr.IPv4Public.String() != "<nil>" && !isPrivateIP(n.Addr.IPv4Public) {
//			return n.Addr.IPv4Public
//		}
//	case "ipv6":
//		if n.Addr.IPv6GlobalUnicast.String() != "<nil>" && !isPrivateIP(n.Addr.IPv6GlobalUnicast) {
//			return n.Addr.IPv6GlobalUnicast
//		}
//	}
//	return nil
//}
//
//// GetPrivateIP NetInterface方法获取私网地址
//func (n NetInterface) GetPrivateIP(ipType string) net.IP {
//	// 如果 NetInterface 未启用,则直接返回空
//	if !n.Flags.Up {
//		return nil
//	}
//
//	switch ipType {
//	case "ipv4":
//		if n.Addr.IPv4Private.String() != "<nil>" && isPrivateIP(n.Addr.IPv4Private) {
//			return n.Addr.IPv4Private
//		}
//	case "ipv6":
//		if n.Addr.IPv6UniqueLocal.String() != "<nil>" && isPrivateIP(n.Addr.IPv6UniqueLocal) {
//			return n.Addr.IPv6UniqueLocal
//		}
//	}
//	return nil
//}
//
//// GetClientIP 获取客户端IP
//func GetClientIP(req *http.Request) string {
//	// 查找 X-Real-Ip
//	IPAddress := req.Header.Get("X-Real-Ip")
//	// 查找 X-Forwarded-For
//	if IPAddress == "" {
//		IPS := strings.Split(req.Header.Get("X-Forwarded-For"), ",")
//		IPAddress = IPS[0]
//	}
//	// 查找 RemoteAddr
//	if IPAddress == "" {
//		IPAddress = req.RemoteAddr
//		IPAddress, _, _ = net.SplitHostPort(IPAddress)
//	}
//	return IPAddress
//}
