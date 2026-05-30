// Package netx 提供网络相关的实用功能
// 功能特性：
//   - IP 地址提取
//   - 端口解析
//   - 本地 IP 获取
// 使用场景：
//   - 服务发现
//   - 端点配置
//   - 网络地址解析
package netx

import (
	"fmt"
	"net"
	"strconv"
)

// ExtractAddress 从 host:port 格式中提取有效的地址
// 功能说明：
//   - 解析 host:port 格式的地址
//   - 如果 host 是通配地址（0.0.0.0、[::]），则尝试获取本机有效 IP
//   - 支持从 Listener 中获取实际端口
//
// 参数：
//   - hostPort: host:port 格式的地址字符串
//   - lis: 网络监听器（可选，用于获取实际端口）
//
// 返回值：
//   - string: 有效地址（格式：ip:port）
//   - error: 解析失败时的错误信息
//
// 使用示例：
//   addr, err := netx.ExtractAddress("0.0.0.0:8080", nil)
//   // 返回本机 IP:8080
func ExtractAddress(hostPort string, lis net.Listener) (string, error) {
	addr, port, err := net.SplitHostPort(hostPort)
	if err != nil {
		return "", err
	}
	if lis != nil {
		if p, ok := Port(lis); ok {
			port = strconv.Itoa(p)
		} else {
			return "", fmt.Errorf("failed to extract port: %v", lis.Addr())
		}
	}
	if len(addr) > 0 && (addr != "0.0.0.0" && addr != "[::]" && addr != "::") {
		return net.JoinHostPort(addr, port), nil
	}
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, rawAddr := range addrs {
			var ip net.IP
			switch addr := rawAddr.(type) {
			case *net.IPAddr:
				ip = addr.IP
			case *net.IPNet:
				ip = addr.IP
			default:
				continue
			}
			if isValidIP(ip.String()) {
				return net.JoinHostPort(ip.String(), port), nil
			}
		}
	}
	return "", nil
}

// ExtractHostPort 从地址中提取 host 和 port
// 功能说明：
//   - 解析地址字符串，分离 host 和 port
//
// 参数：
//   - addr: 地址字符串（格式：host:port）
//
// 返回值：
//   - host: 主机地址
//   - port: 端口号
//   - error: 解析失败时的错误信息
//
// 使用示例：
//   host, port, err := netx.ExtractHostPort("localhost:8080")
func ExtractHostPort(addr string) (host string, port uint64, err error) {
	var (
		ports string
	)
	host, ports, err = net.SplitHostPort(addr)
	if err != nil {
		return
	}
	port, err = strconv.ParseUint(ports, 10, 16)
	if err != nil {
		return
	}
	return
}

// isValidIP 检查 IP 地址是否有效
// 功能说明：
//   - 检查 IP 是否为全局单播地址
//   - 排除接口本地多播地址
func isValidIP(addr string) bool {
	ip := net.ParseIP(addr)
	return ip.IsGlobalUnicast() && !ip.IsInterfaceLocalMulticast()
}

// Port 从 Listener 中获取实际端口
// 功能说明：
//   - 从网络监听器中获取监听的端口号
//
// 参数：
//   - lis: 网络监听器
//
// 返回值：
//   - int: 端口号
//   - bool: 是否成功获取端口
//
// 使用示例：
//   port, ok := netx.Port(listener)
func Port(lis net.Listener) (int, bool) {
	if addr, ok := lis.Addr().(*net.TCPAddr); ok {
		return addr.Port, true
	}
	return 0, false
}

// GetLocalServerIp 获取本地服务器 IP
// 功能说明：
//   - 获取本机非回环地址的 IPv4 地址
//   - 用于服务注册、服务发现等场景
//
// 返回值：
//   - string: 本机 IPv4 地址
//
// 使用示例：
//   ip := netx.GetLocalServerIp()
func GetLocalServerIp() (ip string) {
	manyAddress, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range manyAddress {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
			}
		}
	}

	return
}