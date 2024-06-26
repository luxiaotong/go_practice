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
	clientURL   string
)

var (
	// platform
	ep *httpexpect.Expect
	// backend
	eb *httpexpect.Expect
	// upload
	eu *httpexpect.Expect
	// clinet
	ec *httpexpect.Expect
)

var (
	sampleTmp      string
	applicationTmp string
	voucherTmp     string
	logoTmp        string
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

	clientURL = "http://127.0.0.1:8081"
	// clientURL = "http://127.0.0.1:58081"
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
	ec = httpexpect.New(t, clientURL)

	// Login
	// t.Run("testSignInSeller", testSignInSeller)
	// t.Run("testSignInBuyer", testSignInBuyer)
	// t.Run("testLoginBackend", testLoginBackend)

	// User & Firm
	// t.Run("testGetFirms", testGetFirms)
	// t.Run("testAuditFirm", testAuditFirm)

	// Client Login
	// t.Run("testLoginClient", testLoginClient)
	// t.Run("testKeyPair", testKeyPair)
	// t.Run("testClientStatus", testClientStatus)

	// Client Generate Asset
	// t.Run("testNewDatassets", testNewDatassets)
	// t.Run("testSetDBConn", testSetDBConn)
	// t.Run("testGetTableList", testGetTableList)
	// t.Run("testGetFields", testGetFields)
	// t.Run("testSetMapping", testSetMapping)
	// t.Run("testSchematize", testSchematize)
	// t.Run("testSetDatassetsApply", testSetDatassetsApply)
	// t.Run("testGenerate", testGenerate)
	// t.Run("testGenerateDatassetsApplyPdf", testGenerateDatassetsApplyPdf)
	// t.Run("testListDatassets", testListDatassets)

	// Client Upload Asset
	// t.Run("testLoginClient", testLoginClient)
	// t.Run("testUploadDatassetsApplyPdf", testUploadDatassetsApplyPdf)

	// Local Create Asset
	// t.Run("testUploadApplication", testUploadApplication)
	// t.Run("testIssue", testIssue)
	// t.Run("testAddAsset", testAddAsset)
	// t.Run("testAddAsset_Up", testAddAsset_Up)
	// t.Run("testGetAssets_CityLevel", testGetAssets_CityLevel)
	// t.Run("testGetAsset_ProvLevel", testGetAsset_ProvLevel)

	// productID = 1549641115068534784

	// Audit & Public Asset
	// t.Run("testGetAssets_Seller", testGetAssets_Seller) // get & set productID
	// t.Run("testPreAuditAsset", testPreAuditAsset)
	// t.Run("testFinalAuditAsset", testFinalAuditAsset)
	// t.Run("testUploadLogo", testUploadLogo)
	// t.Run("testUploadSample", testUploadSample)
	// t.Run("testEditAsset", testEditAsset)
	// t.Run("testPublicAudit", testPublicAudit)
	// t.Run("testGetSample", testGetSample)

	// Order
	// t.Run("testAddOrder", testAddOrder)
	// t.Run("testGetOrders_WaitSeller", testGetOrders_WaitSeller) // get & set orderID
	// t.Run("testOpOrder_SellerConfirm", testOpOrder_SellerConfirm)
	// t.Run("testUploadVoucher", testUploadVoucher)
	// t.Run("testPayOrder", testPayOrder)
	// t.Run("testAuditOrder", testAuditOrder)
	// t.Run("testOpOrder_BuyerConfirm", testOpOrder_BuyerConfirm)

	// Client Execute Contract
	// t.Run("testGetOrders_WaitExec", testGetOrders_WaitExec) // get & set orderID
	// t.Run("testGetOrders_Client", testGetOrders_Client)
	// t.Run("testGetContract", testGetContract)
	// t.Run("testExecute", testExecute)
	// t.Run("testAudit", testAudit)
	// t.Run("testDistribute", testDistribute)

	// Client Delete Datassets
	// t.Run("testListDatassets", testListDatassets)
	// t.Run("testDelete", testDelete)
	// t.Run("testListDatassets", testListDatassets)
}
