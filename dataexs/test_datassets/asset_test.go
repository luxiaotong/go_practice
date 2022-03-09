package testdatassets

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

const (
	preAudit    int32 = 10
	finalAudit  int32 = 20
	publicAudit int32 = 30
)

type Industry struct {
	ClassA string `json:"class_a"`
	ClassB string `json:"class_b"`
	ClassC string `json:"class_c"`
	ClassD string `json:"class_d"`
}

type Address struct {
	ProvinceID string `json:"province_id"`
	CityID     string `json:"city_id"`
	CountyID   string `json:"county_id"`
}

type AssetRequest struct {
	ID        int64      `json:"id,string"`
	Name      string     `json:"name"`
	Desc      string     `json:"description"`
	UnitPrice float32    `json:"unit_price"`
	Logo      string     `json:"logo"`
	Industry  *Industry  `json:"industry"`
	Keywords  []string   `json:"keywords"`
	Areas     []*Address `json:"areas"`
	StartDate string     `json:"start_date"`
	EndDate   string     `json:"end_date"`
	Sample    string     `json:"sample"`
	PDF       string     `json:"pdf"`
	Status    int32      `json:"status"`
	Public    bool       `json:"public"`
	ProductID int64      `json:"product_id"`
	UUID      string     `json:"uuid"`
	Reason    string     `json:"reason"`
	Document  string     `json:"document"`
}

type ProductRequest struct {
	ID        int64  `json:"id"`
	Status    int32  `json:"status"`
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
	Query     string `json:"q"`
}

type AuditRequest struct {
	ID        int64  `json:"id,string"`
	AuditType int32  `json:"audit_type"`
	Approved  bool   `json:"approved"`
	Refuse    string `json:"refuse"`
	Document  string `json:"document"`
}

type SampleRequest struct {
	ID    int64  `json:"id,string"`
	Table string `json:"table"`
}

type AssetsRequest struct {
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
	Query     string `json:"q"`
}

func testIssue(t *testing.T) {
	resp := ep.POST("/data/asset/issue").
		WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		Expect().Status(http.StatusOK)
	fmt.Println("data/asset/issue result: ", resp.Body())

	uuid = resp.JSON().Object().Value("data").Object().Value("uuid").String().Raw()
}

func testAddAsset(t *testing.T) {
	req := &AssetRequest{
		PDF:  applicationTmp,
		UUID: uuid,
	}
	resp := ep.POST("/data/asset").
		WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/asset add result: ", resp.Body())

	productID = int64(resp.JSON().Object().Value("data").Object().Value("product_id").Number().Raw())
}

func testPreAuditAsset(t *testing.T) {
	req := &AuditRequest{
		ID:        productID,
		AuditType: preAudit,
		Approved:  true,
	}
	resp := eb.POST("/data/asset/status").
		WithCookie(backCookie, provUserToken).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/asset/status pre audit result: ", resp.Body())
}

func testFinalAuditAsset(t *testing.T) {
	// productID = 9
	req := &AuditRequest{
		ID:        productID,
		AuditType: finalAudit,
		Approved:  true,
	}
	resp := eb.POST("/data/asset/status").
		WithCookie(backCookie, provUserToken).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/asset/status final audit result: ", resp.Body())
}

func testGetAsset_ProvLevel(t *testing.T) {
	resp := eb.GET("/data/asset/9").
		WithCookie(backCookie, provUserToken).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/asset/9 result: ", resp.Body())
}

func testGetAssets_CityLevel(t *testing.T) {
	req := &ProductRequest{PageIndex: 1, PageSize: 10}
	resp := eb.POST("/data/assets").
		WithCookie(backCookie, cityUserToken).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/assets in city level result: ", resp.Body())
}

func testEditAsset(t *testing.T) {
	// lg := "logo.jpg"
	lg := ""
	req := &AssetRequest{
		ID:        productID,
		Desc:      "req.Desc",
		UnitPrice: 1.00,
		Industry:  &Industry{ClassA: "农、林、牧、渔业", ClassB: "农业", ClassC: "谷物种植", ClassD: "稻谷种植"},
		Keywords:  []string{"1", "2", "3"},
		Areas: []*Address{
			{ProvinceID: "420000000000", CityID: "420100000000", CountyID: "420102000000"}, // 湖北省武汉市江岸区
			{ProvinceID: "410000000000", CityID: "410000000000", CountyID: "410000000000"}, // 河南省
		},
		StartDate: "2020-05-06T00:08:50.00+08:00",
		EndDate:   "2020-05-20T00:08:50.00+08:00",
		Logo:      lg,
		Sample:    sampleTmp,
		Public:    true,
	}
	resp := ep.PUT("/data/asset").WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/asset edit result: ", resp.Body())
}

func testPublicAudit(t *testing.T) {
	req := &AuditRequest{
		ID:        productID,
		AuditType: publicAudit,
		Approved:  true,
	}
	resp := eb.POST("/data/asset/status").
		WithCookie(backCookie, provUserToken).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/asset/status final audit result: ", resp.Body())
}

func testGetSample(t *testing.T) {
	req := &SampleRequest{
		ID:    productID,
		Table: "JG_JGSX_CONVERGE_CLAIM",
	}
	resp := ep.POST("/data/product/sample").WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/product/sample result: ", resp.Body())
}

func testGetAssets_Seller(t *testing.T) {
	req := &AssetsRequest{
		Query:     "河南省新乡市统计局数据2",
		PageIndex: 1,
		PageSize:  10,
	}
	resp := ep.POST("/data/assets").
		WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/assets seller result: ", resp.Body())
	pid := resp.JSON().Object().Value("data").Object().Value("list").Array().First().Object().Value("id").String().Raw()
	productID, _ = strconv.ParseInt(pid, 10, 64)
	fmt.Println("latest product id: ", productID)
}
