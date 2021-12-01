package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func encrypt(clientID, secret string) []byte {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	origData := []byte(clientID + ":" + secret + ":" + ts + "000")
	fmt.Println("info: ", string(origData))
	k, _ := hex.DecodeString(secret)
	cipher, _ := aes.NewCipher(k)
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encrypted
}

func main() {
	host := "https://t200renzheng.zhengtoon.com"
	clientID := "20211117001"
	code := "d2523449-33eb-38ab-adcc-d5313bbf7a7e"
	enc := encrypt(clientID, "31c241309a9231f585bca20c9873b49a")
	fmt.Println("enc: ", hex.EncodeToString(enc))
	v := url.Values{}
	v.Set("client_id", clientID)
	v.Set("grant_type", "authorization_code")
	v.Set("grant_code", code)
	v.Set("auth_token", hex.EncodeToString(enc))
	client := &http.Client{}
	resp, err := client.PostForm(host+"/api/oauth/getAccessToken", v)
	if err != nil {
		fmt.Printf("natural access token request error: %v\n", err)
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(errors.Errorf("natural access token request failed: %s", resp.Status))
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if len(respBody) == 0 {
		panic(errors.Errorf("natural access token request invalid code: %s", code))
	}
	fmt.Printf("natural access token resp body: %s\n", string(respBody))
}

// this 1f7ffc274d02707a3a47f459d85c6af87ec83dc21512913720f7089f7a00fa44f2ba09284c05b0a8e23c38b6a77a467e67a181b97eabafc25d8006d33142a45b
// that 1f7ffc274d02707a3a47f459d85c6af87ec83dc21512913720f7089f7a00fa44f2ba09284c05b0a8e23c38b6a77a467e67a181b97eabafc25d8006d33142a45b
