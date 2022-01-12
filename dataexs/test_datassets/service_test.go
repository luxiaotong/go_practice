package testdatassets

import (
	"context"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
)

const jwtCookieSecret = "crc_key"

var ctx context.Context

var (
	platformURL string
	uploadURL   string
)

var (
	// platform
	ep *httpexpect.Expect
	// upload
	eu *httpexpect.Expect
)

var (
	sampleTmp      string
	applicationTmp string
)

var (
	productID int64
)

func initEnv() {
	// platformURL = "http://127.0.0.1:8080"
	// uploadURL = "http://127.0.0.1:8085"
	platformURL = "http://139.9.119.21:58099"
	uploadURL = "http://139.9.119.21:58098"
}

func initData() {
	productID = GenID()
}

func TestMain(m *testing.M) {
	ctx = context.Background()
	initEnv()
	initData()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestAll(t *testing.T) {
	ep = httpexpect.New(t, platformURL)
	eu = httpexpect.New(t, uploadURL)
	t.Run("testSignIn", testSignIn)
	t.Run("testUploadApplication", testUploadApplication)
	t.Run("testAddAsset", testAddAsset)
	t.Run("testUploadSample", testUploadSample)
	t.Run("testEditAsset", testEditAsset)
	t.Run("testGetSample", testGetSample)
}
