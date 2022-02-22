package testdatassets

import (
	"context"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
)

const (
	jwtCookieSecret = "crc_key"
	backCookie      = "back_token"
)

var ctx context.Context

var (
	platformURL string
	backendURL  string
	uploadURL   string
)

var (
	// platform
	ep *httpexpect.Expect
	// backend
	eb *httpexpect.Expect
	// upload
	eu *httpexpect.Expect
)

var (
	sampleTmp      string
	applicationTmp string
)

var (
	uuid      string
	productID int64
)

func initEnv() {
	// platformURL = "http://127.0.0.1:8080"
	// backendURL = "http://127.0.0.1:9091"
	// uploadURL = "http://127.0.0.1:8085"

	platformURL = "http://139.9.119.21:58099"
	backendURL = "http://139.9.119.21:58110"
	uploadURL = "http://139.9.119.21:58098"
}

func initData() {
	// productID = GenID()
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
	eb = httpexpect.New(t, backendURL)
	eu = httpexpect.New(t, uploadURL)
	t.Run("testSignInPlatform", testSignInPlatform)
	t.Run("testLoginBackend", testLoginBackend)
	t.Run("testUploadApplication", testUploadApplication)
	t.Run("testIssue", testIssue)
	t.Run("testAddAsset", testAddAsset)
	t.Run("testAuditAsset", testAuditAsset)
	t.Run("testUploadSample", testUploadSample)
	t.Run("testEditAsset", testEditAsset)
	t.Run("testGetSample", testGetSample)
}
