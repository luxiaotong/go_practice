package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const ssoURL = "http://127.0.0.1:9096"
const (
	tokenAPI = "/oauth2/token"
	userAPI  = "/oauth2/getuserinfo"
)

func main() {
	log.Println("test oauth client start")

	http.HandleFunc("/oauth2/callback", callbackHandler)

	log.Println("oauth server addr: http://127.0.0.1:9094")
	log.Fatal(http.ListenAndServe(":9094", nil))
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("go_session_id")
	if err != nil {
		fmt.Printf("get cookie error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("go_session_id=%v", ck.Value)
	code := r.URL.Query().Get("code")
	fmt.Println("code =>", code)
	params := url.Values{}
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", "http://127.0.0.1:9094/oauth2/callback")
	params.Add("client_id", "test")
	log.Printf("url: %v, params: %v\n", ssoURL+tokenAPI, params.Encode())

	req, err := http.NewRequest(http.MethodPost, ssoURL+tokenAPI, strings.NewReader(params.Encode()))
	if err != nil {
		log.Printf("new request error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "go_session_id="+ck.Value)
	// resp, err := http.PostForm(ssoURL+tokenAPI, params)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Printf("get token error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respBody := resp.Body
	defer respBody.Close()
	b, err := io.ReadAll(respBody)
	if err != nil {
		log.Printf("read resp body error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("resp body: ", string(b))
	type getTokenResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
	}
	var data getTokenResponse
	if err := json.Unmarshal(b, &data); err != nil {
		log.Printf("decode resp body error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("get token resp: %#v", data)
	// _, _ = w.Write(b)
	if len(data.AccessToken) == 0 {
		return
	}

	// get user info
	req, err = http.NewRequest(http.MethodGet, ssoURL+userAPI, nil)
	if err != nil {
		log.Printf("new request error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+data.AccessToken)
	cli = &http.Client{}
	resp, err = cli.Do(req)
	if err != nil {
		log.Printf("get user error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respBody = resp.Body
	defer respBody.Close()
	b, err = io.ReadAll(respBody)
	if err != nil {
		log.Printf("read resp body error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("resp body: ", string(b))
}
