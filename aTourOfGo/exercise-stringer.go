package main

import "fmt"

//TODO: Add a "String() string" method to IPAddr.

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDns": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("name:%v,ip:%v\n", name, ip)
	}
}
