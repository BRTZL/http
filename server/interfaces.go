package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/valyala/fasthttp"
)

func getInterfaces(ctx *fasthttp.RequestCtx) {

	data := "\n"

	fmt.Println("=== interfaces ===")

	// Get interfaces by net library
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {

		data += fmt.Sprintf("%v", iface) + "\n\t"
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			addrStr := addr.String()
			data += addr.Network() + " " + addrStr

			// Must drop the stuff after the slash in order to convert it to an IP instance
			split := strings.Split(addrStr, "/")
			addrStr0 := split[0]

			// Parse the string to an IP instance
			ip := net.ParseIP(addrStr0)
			if ip.To4() != nil {
				data += " - " + addrStr0 + " is ipv4\n"
			} else {
				data += " - " + addrStr0 + " is ipv6\n"
			}
		}
	}

	// Write to the body
	data = security(data, false)
	ctx.SetBody([]byte(data))
	// TODO: secure here

}
