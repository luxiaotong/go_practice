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
	trustURL    string
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
	voucherTmp     string
)

var (
	uuid      string
	productID int64
	orderID   int64
)

func initEnv() {
	platformURL = "http://139.9.119.21:58099"
	// platformURL = "http://127.0.0.1:8080"

	backendURL = "http://139.9.119.21:58110"
	// backendURL = "http://127.0.0.1:9091"

	uploadURL = "http://139.9.119.21:58098"
	// uploadURL = "http://127.0.0.1:8085"

	trustURL = "http://139.9.119.21:58300"
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
	t.Run("testSignInSeller", testSignInSeller)
	t.Run("testSignInBuyer", testSignInBuyer)
	t.Run("testLoginBackend", testLoginBackend)

	t.Run("testUploadApplication", testUploadApplication)
	t.Run("testIssue", testIssue)
	t.Run("testAddAsset", testAddAsset)
	t.Run("testGetAssets", testGetAssets)
	t.Run("testPreAuditAsset", testPreAuditAsset)
	t.Run("testFinalAuditAsset", testFinalAuditAsset)
	t.Run("testGetAsset", testGetAsset)
	t.Run("testUploadSample", testUploadSample)
	t.Run("testEditAsset", testEditAsset)
	t.Run("testPublicAudit", testPublicAudit)
	t.Run("testGetSample", testGetSample)

	t.Run("testAddOrder", testAddOrder)
	t.Run("testGetOrders", testGetOrders)
	t.Run("testOpOrder_SellerConfirm", testOpOrder_SellerConfirm)
	t.Run("testUploadVoucher", testUploadVoucher)
	t.Run("testPayOrder", testPayOrder)
	t.Run("testAuditOrder", testAuditOrder)
	t.Run("testOpOrder_BuyerConfirm", testOpOrder_BuyerConfirm)
}
