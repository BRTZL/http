package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func makeRequest() (makeit bool) {

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

	req, err := http.NewRequest("POST", "https://localhost:"+config.Port, bytes.NewBuffer([]byte("")))
	check(err)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	//fmt.Println("request Headers:", req.Header)

	fmt.Println("time        :", time.Now())
	fmt.Println("cookie      :", &http.Cookie{Name: "token", Value: token})

	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	//fmt.Println("respond Headers:", resp.Header)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
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
