package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	str := "https://www.google.com?q=\"Lady Gaga\""

	sEnc := b64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(str))
	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
