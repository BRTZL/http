package main

import (
	"io/ioutil"
)

func saveToken(_token string) {
	data := []byte(_token)
	err := ioutil.WriteFile("/tmp/dat", data, 0644)
	check(err)
}

func getToken() {
	data, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	token = string(data)
}
