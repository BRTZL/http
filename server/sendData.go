package main

import (
	"crypto/rand"

	"github.com/valyala/fasthttp"
)

func sendData(ctx *fasthttp.RequestCtx, index int) {

	blk := make([]byte, index)
	_, err := rand.Read(blk)
	check(err)

	data := security(string(blk), false)
	ctx.SetBody([]byte(data))
}
