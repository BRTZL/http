package main

import (
    "fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func secure(strToEncrDcr string,decryptOnly bool) string{

	if decryptOnly {
		decr, err := DecryptToString(strToEncrDcr)
		if err != nil {
			fmt.Println("Error Decrypting!")
		} else {
			fmt.Println("BlowFish Decrypted String: " + decr)
      return decr
		}
	} else {
		encr, err := EncryptToString(strToEncrDcr)
		if err != nil {
			fmt.Println("Error Encrypting!")
		} else {
			fmt.Println("BlowFish Encrypted String: " + encr)
      return encr
		}
	}
  return ""
}

func main() {
	secure(secure("deneme deneme",false),true)

}
