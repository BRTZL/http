package main

import (
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/blowfish"
)

const _CipherKey string = "parta network"

// decryptOnly true for decryption
// decryptOnly false for encryption
func security(strToEncrDcr string, decryptOnly bool) string {
	if decryptOnly {
		decr, err := DecryptToString(strToEncrDcr)
		if err != nil {
			fmt.Println("Error Decrypting!")
		} else {
			//fmt.Println("BlowFish Decrypted String: " + decr)
			return decr
		}
	} else {
		encr, err := EncryptToString(strToEncrDcr)
		if err != nil {
			fmt.Println("Error Encrypting!")
		} else {
			//fmt.Println("BlowFish Encrypted String: " + encr)
			return encr
		}
	}
	return ""
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func blowfishChecksizeAndPad(value []byte) []byte {
	modulus := len(value) % blowfish.BlockSize
	if modulus != 0 {
		padnglen := blowfish.BlockSize - modulus
		for i := 0; i < padnglen; i++ {
			value = append(value, 0)
		}
	}
	return value
}

func blowfishEncrypt(value, key []byte) ([]byte, error) {
	bcipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	returnMe := make([]byte, blowfish.BlockSize+len(value))
	eiv := returnMe[:blowfish.BlockSize]
	ecbc := cipher.NewCBCEncrypter(bcipher, eiv)
	ecbc.CryptBlocks(returnMe[blowfish.BlockSize:], value)
	return returnMe, nil
}

func EncryptToByte(value string) ([]byte, error) {
	var returnMe, valueInByteArr, paddedByteArr, keyByteArr []byte
	valueInByteArr = []byte(value)
	keyByteArr = []byte(_CipherKey)
	paddedByteArr = blowfishChecksizeAndPad(valueInByteArr)
	returnMe, err := blowfishEncrypt(paddedByteArr, keyByteArr)
	if err != nil {
		return nil, err
	}
	return returnMe, nil
}

func EncryptToString(value string) (string, error) {
	encryptedByteArr, err := EncryptToByte(value)
	if err != nil {
		return "", err
	}
	returnMe := encodeBase64(encryptedByteArr)
	return returnMe, nil
}

func decodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func blowfishDecrypt(value, key []byte) ([]byte, error) {
	dcipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	div := value[:blowfish.BlockSize]
	decrypted := value[blowfish.BlockSize:]
	if len(decrypted)%blowfish.BlockSize != 0 {
		return nil, err
	}
	dcbc := cipher.NewCBCDecrypter(dcipher, div)
	dcbc.CryptBlocks(decrypted, decrypted)
	return decrypted, nil
}

func DecryptToByte(value string) ([]byte, error) {
	var returnMe, keyByteArr []byte
	keyByteArr = []byte(_CipherKey)
	decodeB64, err1 := decodeBase64(value)
	if err1 != nil {
		return nil, err1
	}
	returnMe, err2 := blowfishDecrypt(decodeB64, keyByteArr)
	if err2 != nil {
		return nil, err2
	}
	return returnMe, nil
}

func DecryptToString(value string) (string, error) {
	decryptedByteArr, err := DecryptToByte(value)
	if decryptedByteArr == nil {
		return "", err
	}
	var returnMe string = string(decryptedByteArr[:])
	return returnMe, nil
}
