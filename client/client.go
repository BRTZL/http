package main

import (
	"fmt"
)

func main() {
	getConfig()
	getToken()

	if makeRequest() != true {
		logIn()
	}
	for {
		fmt.Scanln(&query)
		fmt.Println()
		makeRequest()
	}
}
