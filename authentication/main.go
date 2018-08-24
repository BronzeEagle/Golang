package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	message := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYTFiMmMzIiwidXNlcm5hbWUiOiJuaWtvbGEifQ=="
	sKey := "42isTheAnswer"

	key := []byte(sKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	b := base64.URLEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(string(b))
}
