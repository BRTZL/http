package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	secure "github.com/mithorium/secure-fasthttp"
)

var (
	debug   bool
	timeOut int
	configs Configs
)

var jwtKey = []byte("partanetwork")

// Local database
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Configs : config settings from conf.json
type Configs struct {
	Port string `json:"port"`
	Time string `json:"time"`
}

// Credentials : username and password
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims : claim token for specific username
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func check(err error) {
	if err == sql.ErrNoRows {
		err = nil // Ignore not-found rows
	}
	if err != nil {
		panic(err)
	}
}

var secureMiddleware = secure.New(secure.Options{
	AllowedHosts:          []string{"localhost", "https://localhost:8443", "https://127.0.0.1:8443"},
	SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	STSSeconds:            315360000,
	SSLRedirect:           true,
	SSLHost:               "localhost" + configs.Port,
	STSIncludeSubdomains:  true,
	STSPreload:            true,
	FrameDeny:             true,
	ContentTypeNosniff:    true,
	BrowserXssFilter:      true,
	ContentSecurityPolicy: "default-src 'self'",
	PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://localhost/hpkp-report"`,
	IsDevelopment:         true,
})

func getConfig() {
	jsonFile, err := os.Open("conf.json")
	if err != nil {
		check(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal json file to struct
	json.Unmarshal(byteValue, &configs)

	if configs.Port == "" || configs.Time == "" {
		log.Fatalf("Error in config file")
	}

	timeOut, _ = strconv.Atoi(configs.Time)
}
