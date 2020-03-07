package main

import (
	"crypto/x509"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	token      string
	query      string
	config     Configs
	caCertPool *x509.CertPool
)

// Configs : config settings from conf.json
type Configs struct {
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func check(err error) {
	if err == sql.ErrNoRows {
		err = nil // Ignore not-found rows
	}
	if err != nil {
		panic(err)
	}
}

func getConfig() {
	jsonFile, err := os.Open("conf.json")
	check(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal json file to struct
	json.Unmarshal(byteValue, &config)

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool = x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

}
