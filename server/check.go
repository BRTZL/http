package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

func checkCreds(creds Credentials, ctx *fasthttp.RequestCtx) {
	fmt.Println("works")
	// Json parser
	body := security(string(ctx.PostBody()), true)
	//body := string(ctx.PostBody())
	//fmt.Println(body)
	err := json.NewDecoder(bytes.NewReader([]byte(body))).Decode(&creds)
	// TODO: secure here
	// NOTE: return type of body is []byte
	// NOTE: I change here ctx.PostBody() -> ctx.Request.Body()

	if err != nil {
		fmt.Println(err)
		ctx.Response.SetStatusCode(400)
		return
	}

	// If creds are not empty
	if creds.Password != "" && creds.Username != "" {

		// Check username and password
		expectedPassword, ok := users[creds.Username]
		if !ok || expectedPassword != creds.Password {
			ctx.Response.SetStatusCode(401)
			fmt.Println("username or password wrong")
			return
		}

		// Add token with time out and ads to the claimed db
		expirationTime := time.Now().Add(time.Duration(timeOut) * time.Second)
		claims := &Claims{
			Username: creds.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			ctx.Response.SetStatusCode(500)
			return
		}

		// Add token to the cookie
		authCookie := fasthttp.Cookie{}
		authCookie.SetKey("token")
		authCookie.SetValue(tokenString)
		authCookie.SetExpire(expirationTime)
		ctx.Response.Header.SetCookie(&authCookie)

	} else {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		if debug {
			fmt.Println("no data entered")
		}
	}

}

func tokenValidation(ctx *fasthttp.RequestCtx) bool {
	// Get cookie named token
	tknStr := string(ctx.Request.Header.Cookie("token"))

	// Checks for validation
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return false
		}
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return false
	}

	if !tkn.Valid {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return false
	}

	return true

}
