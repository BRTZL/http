package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func makeRequest() (makeit bool) {
	// Create a HTTPS client and supply the created CA pool
	client := &http.Client{
		// NOTE: https remove for ssl
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	req, err := http.NewRequest("POST", "https://localhost:"+config.Port, bytes.NewBuffer([]byte("")))
	check(err)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	//fmt.Println("request Headers:", req.Header)

	fmt.Println("time        :", time.Now())
	fmt.Println("cookie      :", &http.Cookie{Name: "token", Value: token})

	resp, err := client.Do(req)
	body := resp.Body
	resp.Body.Close()

	check(err)

	//fmt.Println("respond Headers:", resp.Header)
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	newStr := buf.String()
	// TODO: secure here
	if newStr != "" {
		newStr = security(newStr, true)
	}
	fmt.Println("respond Body:", newStr)

	if resp.StatusCode == 401 {
		fmt.Println("token not correct")
		logIn()
	} else if resp.StatusCode == 200 {
		return true
	} else if resp.StatusCode == 400 {
		fmt.Printf("error\n\n")
		logIn()
	}
	return false
}
