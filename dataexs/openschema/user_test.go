package main

import (
	"database/sql"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const CookieSecret = "crc_key"
const defaultSecret = "secret@datassets"

var (
	cookieVal string
	token     string
	uid       int64
)

type UserRequest struct {
	ID        int64  `json:"id,string"`
	Username  string `json:"username"`
	Mobile    string `json:"mobile"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	FirmName  string `json:"firm_name"`
	FirmAbbr  string `json:"firm_abbr"`
	Logo      string `json:"logo"`
	Specialty string `json:"specialty"`
	Award     string `json:"award"`
	Desc      string `json:"desc"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PassRequest struct {
	New string `json:"new"`
	Old string `json:"old"`
}

func testSignUp(t *testing.T) {
	req := UserRequest{
		Username: "shannon",
		Password: "123456",
		Mobile:   "18500022713",
		Name:     "shannon",
		Email:    "shannon.lu@datassets.cn",
		FirmName: "abcabcabc",
		FirmAbbr: "abc",
	}
	resp := e.POST("/user/signup").WithJSON(req).Expect().Status(http.StatusOK)
	cookieVal = resp.Cookie(CookieSecret).Value().Raw()
	token = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	assert.NotEmpty(t, cookieVal)
	assert.NotEmpty(t, token)
	uid = parse(token, cookieVal)
	fmt.Println("signup user id: ", uid)
}

func parse(tokenString, key string) int64 {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(defaultSecret + ":" + key), nil
	})
	if err != nil {
		fmt.Println("parse err: ", err)
		return 0
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int64(claims["v"].(map[string]interface{})["i"].(float64))
	} else {
		return 0
	}
}

func testSignIn(t *testing.T) {
	req := SignInRequest{
		Username: "shannon",
		Password: "123456",
	}
	resp := e.POST("/user/signin").WithJSON(req).Expect().Status(http.StatusOK)
	cookieVal = resp.Cookie(CookieSecret).Value().Raw()
	token = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Println("user cookie: ", cookieVal)
	fmt.Println("user token: ", token)
	assert.NotEmpty(t, cookieVal)
	assert.NotEmpty(t, token)
	uid = parse(token, cookieVal)
}

func testUserPass(t *testing.T) {
	req := PassRequest{
		Old: "123456",
		New: "123456",
	}
	resp := e.POST("/user/pass").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/pass response: %v\n", resp.Body())
}

func testUpdateUser(t *testing.T) {
	req := UserRequest{
		Mobile:   "18500022713",
		Name:     "shannon",
		Email:    "shannon@datassets.cn",
		FirmName: "firm_name_2",
		FirmAbbr: "firm_abbr_2",
	}
	resp := e.PUT("/user/info").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/info update response: %v\n", resp.Body())
}

func testUpdateUser_Logo(t *testing.T) {
	f, _ := ioutil.TempFile("", "*.jpg")
	fmt.Println("tmp refund voucher: ", f.Name())
	defer os.Remove(f.Name())
	defer f.Close()
	alpha := image.NewAlpha(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			alpha.Set(x, y, color.Alpha{uint8(x % 256)})
		}
	}
	_ = jpeg.Encode(f, alpha, nil)

	resp := e.POST("/upload/file").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithMultipart().
		WithFile("file", f.Name()).WithFormField("api_type", "logo").
		Expect().Status(http.StatusOK)
	name := resp.JSON().Object().Value("data").Object().Value("name").String().Raw()
	fmt.Printf("/upload/file result: %v\n", name[6:])
	logo := name[6:]

	req := UserRequest{
		Mobile:   "18500022713",
		Name:     "shannon",
		Email:    "shannon@datassets.cn",
		FirmName: "firm_name_2",
		FirmAbbr: "firm_abbr_2",
		Logo:     logo,
	}
	resp = e.PUT("/user/info").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/info with logo response: %v\n", resp.Body())
}

func clearUser() {
	dsn := "host=139.9.119.21 port=5432 user=auth password=authpass dbname=openschema sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("open error: ", err)
		return
	}
	defer db.Close()

	if _, err := db.Exec("delete from users where id=$1", uid); err != nil {
		fmt.Println("delete error: ", err)
		return
	}
	clearLogo(uid)
}

func testCertifyUser(t *testing.T) {
	req := UserRequest{
		Specialty: "specialty",
		Award:     "award",
		Desc:      "desc",
	}
	resp := e.PUT("/user/certify").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/certify response: %v\n", resp.Body())
}

func testGetUser(t *testing.T) {
	req := &UserRequest{}
	resp := e.POST("/user/info").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/info get response: %v\n", resp.Body())
}
