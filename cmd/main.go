package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	ipString := func (ip1, ip2 net.IP) string{
		str1 := ip1.String()
		str2 := ip2.String()
		if str1 < str2 {
			return fmt.Sprintf("%s<->%s", str1, str2)
		}
		return fmt.Sprintf("%s<->%s", str2, str1)
	}

	fmt.Println(ipString(net.ParseIP("192.168.1.1"), net.ParseIP("192.168.1.2")))
	fmt.Println(ipString(net.ParseIP("192.168.1.4"), net.ParseIP("192.168.1.2")))
	fmt.Println(ipString(net.ParseIP("192.168.3.1"), net.ParseIP("192.168.1.2")))

	_, n, _ := net.ParseCIDR("99.86.119.0/24")

	mask := net.ParseIP("99.86.119.0")
	filter := func(ip, mask net.IP) bool {
		m := net.IPMask(mask.To4())
		result := ip.Mask(m)

		fmt.Printf("result: %s, mask: %s \n", result.String(), mask.String())
		return strings.EqualFold(result.String(), mask.String())
	}

	s := "99.86.119.31"
	result := filter(net.ParseIP(s), mask)
	fmt.Printf("ip: %s match: %v\n", s, result)
	fmt.Printf("parsed: %v, result: %v\n", net.ParseIP(s), n.Contains(net.ParseIP(s)))

	s = "99.88.119.31"
	result = filter(net.ParseIP(s), mask)
	fmt.Printf("ip: %s match: %v\n", s, result)
	fmt.Printf("parsed: %v, result: %v\n", net.ParseIP(s), n.Contains(net.ParseIP(s)))


	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}