package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"` // user id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	// create header for the jwt and convert to base 64 data
	header := &Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	b64Header := Base64Encode(byteArrHeader)

	// convert payload data to base 64 data
	byteArrData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	b64Data := Base64Encode(byteArrData)
	message := b64Header + "." + b64Data

	// use the raw secret bytes for HMAC key
	byteArrSecret := []byte(secret)

	byteArrMessage := []byte(message)

	// use hmac package to hash the message and secret together, in their bytes form
	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)
	signature := h.Sum(nil)
	b64Signature := Base64Encode(signature)

	jwt := b64Header + "." + b64Data + "." + b64Signature

	return jwt, nil

}

func Base64Encode(b []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}
