// Exercise: Stringers

// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

// for instance, IPAddr{1, 3, 3, 4} should print as "1.2.3.4".

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
///
func (ip_Addr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip_Addr[0], ip_Addr[1], ip_Addr[2], ip_Addr[2])
}
///

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v:\t%v\n", name, ip)
	}
}
