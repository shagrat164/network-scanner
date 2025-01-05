package ping

import (
	"fmt"
	"net"
)

// Генерация списка IP из диапазона
func GenerateIPs(startIP, endIP string) ([]string, error) {
	start := net.ParseIP(startIP)
	end := net.ParseIP(endIP)
	if start == nil || end == nil {
		return nil, fmt.Errorf("некорректный IP-адрес")
	}

	var ips []string
	for ip := start; !ip.Equal(end); incIP(ip) {
		ips = append(ips, ip.String())
	}
	ips = append(ips, end.String())
	return ips, nil
}

// Увеличение IP-адреса
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
