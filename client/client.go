package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var index int

func main() {
	getConfig()
	getToken()

	past := ""

	if makeRequest() != true {
		logIn()
	}

	for index = 0; index < 20; index++ {
		// fmt.Scanln(&query)
		// fmt.Println()
		start := time.Now()
		makeRequest()
		t := time.Now()
		elapsed := t.Sub(start)
		past += strconv.Itoa(index) + "    " + elapsed.String() + "\n"

		fmt.Println(elapsed)
	}
	data := []byte(past)
	dir, err := os.Getwd()
	fmt.Println(dir)
	check(err)
	err = ioutil.WriteFile(dir+"/dat", data, 0644)
	check(err)
}
