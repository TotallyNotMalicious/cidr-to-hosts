package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

func main() {
	for _, hosts := range os.Args {
		display(hosts)
	}
}

func validate(cidr string) bool {
	regex := `\b(?:\d{1,3}\.){3}\d{1,3}\b$`
	match, _ := regexp.MatchString(regex, cidr)
	return match
}

func display(cidr string) {
	var hosts []string

	if validate(cidr) {
		fmt.Println(cidr)
		return
	}

	address, net, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	for host := address.Mask(net.Mask); net.Contains(host); inc(host) {
		hosts = append(hosts, host.String())
	}

	if len(hosts) <= 2 {
		return
	}

	for _, host := range hosts[1 : len(hosts)-1] {
		fmt.Println(host)
	}
}

func inc(host net.IP) {
	for i := len(host) - 1; i >= 0; i-- {
		host[i]++
		if host[i] != 0 {
			break
		}
	}
}
