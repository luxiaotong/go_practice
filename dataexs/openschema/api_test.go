package main

import (
	"testing"

	"github.com/gavv/httpexpect"
)

var e *httpexpect.Expect

func init() {
}

func TestAll(t *testing.T) {
	e = httpexpect.New(t, "http://127.0.0.1:8096")
	t.Run("testSignUp", testSignUp)
	t.Run("testSignIn", testSignIn)
	t.Run("testUserPass", testUserPass)
	t.Run("testUploadFile", testUploadFile)
	t.Run("testUpdateUser", testUpdateUser)
	t.Run("testUpdateUser_Logo", testUpdateUser_Logo)
	defer clearUser()
}
