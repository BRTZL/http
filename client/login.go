package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func logIn() {

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a HTTPS client and supply the created CA pool
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	jsonBody := "{\"username\":\"" + config.Username + "\",\"password\":\"" + config.Password + "\"}"
	jsonBody = security(jsonBody, false)
	// TODO: secure here

	var jsonStr = []byte(jsonBody)
	req, err := http.NewRequest("POST", "https://localhost:"+config.Port, bytes.NewBuffer(jsonStr))
	check(err)
	fmt.Println("request Headers:", req.Header)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "token" {
			token = cookie.Value
			saveToken(token)
		}
	}

	if resp.StatusCode == 401 {
		fmt.Println("username or password might wrong check!")
		fmt.Print("username: ")
		fmt.Scanln(&config.Username)
		fmt.Print("password: ")
		fmt.Scanln(&config.Password)
		logIn()
	} else if resp.StatusCode == 200 {
		makeRequest()
	}
}
