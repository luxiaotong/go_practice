package main

import (
	"testing"

	"github.com/gavv/httpexpect"
)

var e *httpexpect.Expect
var u *httpexpect.Expect

func init() {
}

func TestAll(t *testing.T) {
	e = httpexpect.New(t, "http://139.9.119.21:58099")
	u = httpexpect.New(t, "http://139.9.119.21:58098")
	t.Run("testSignIn", testSignIn)
	t.Run("testUploadFile_Sample", testUploadFile_Sample)
}
