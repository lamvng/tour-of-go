package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte // Array (Fixed size)

func (ip IPAddr) String() string {
	var ipString [4]string
	for i := 0; i < len(ip); i++ {
		ipString[i] = strconv.Itoa(int(ip[i]))
	}
	return strings.Join(ipString[:], ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
