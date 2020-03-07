package main

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var index int = 0

func requestHandler(ctx *fasthttp.RequestCtx) {
	var creds Credentials

	fmt.Println(time.Now())

	// If there is no cookie in existence check if creds are correct for login otherwise getInterfaces
	if ctx.Request.Header.Cookie("token") == nil {
		checkCreds(creds, ctx)
	} else if tokenValidation(ctx) {
		fmt.Println("token is valid")
		index++
		//getInterfaces(ctx)
		sendData(ctx, index)
		fmt.Println(index)
	}

}
