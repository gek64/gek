package gNet

import (
	"net"
	"net/netip"
)

// NetInterface 网络接口
type NetInterface struct {
	Name string   `json:"name"`
	Ipv4 []string `json:"ipv4"`
	Ipv6 []string `json:"ipv6"`
	Flag Flag     `json:"flag"`
	Mac  string   `json:"mac"`
}
type Flag struct {
	Up        bool `json:"up"`
	Unicast   bool `json:"unicast"`
	Broadcast bool `json:"broadcast"`
	Multicast bool `json:"multicast"`
}

func GetNetInterfaces() (netInterface []NetInterface, err error) {
	// 获取系统中所有的网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// 对每一个系统网络接口循环
	for _, i := range interfaces {
		// 新建自定义网络接口
		var ni NetInterface

		// 获取系统网络接口地址
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			// 解析cidr,addr中ip是以带后缀的ip形式存在的例如 192.168.1.1/24
			parsePrefix, err := netip.ParsePrefix(addr.String())
			if err != nil {
				return nil, err
			}

			// 地址分ipv4或ipv6存入当前自定义网络接口
			if parsePrefix.Addr().Is4() {
				ni.Ipv4 = append(ni.Ipv4, parsePrefix.Addr().String())
			} else {
				ni.Ipv6 = append(ni.Ipv6, parsePrefix.Addr().String())
			}
		}
		// 填充当前自定义网络接口
		ni.Name = i.Name
		ni.Mac = i.HardwareAddr.String()
		ni.Flag.Up = i.Flags&net.FlagUp != 0
		ni.Flag.Unicast = i.Flags&net.FlagPointToPoint != 0
		ni.Flag.Broadcast = i.Flags&net.FlagBroadcast != 0
		ni.Flag.Multicast = i.Flags&net.FlagMulticast != 0

		// 填充自定义网络接口切片
		netInterface = append(netInterface, ni)
	}

	return netInterface, err
}
