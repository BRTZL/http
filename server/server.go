package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	debug = true
	getConfig()

	fmt.Println("Started")

	secureHandler := secureMiddleware.Handler(requestHandler)

	go func() {
		log.Fatal(fasthttp.ListenAndServe(":8080", secureHandler))
	}()

	if err := fasthttp.ListenAndServeTLS(":"+configs.Port, "cert.pem", "key.pem", secureHandler); err != nil {
		// if err := fasthttp.ListenAndServe(":"+configs.Port, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
