package testdatassets

import (
	"fmt"
	"net/http"
	"testing"
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

type SampleRequest struct {
	ID    int64  `json:"id,string"`
	Table string `json:"table"`
}

func testAddAsset(t *testing.T) {
	req := &AssetRequest{
		PDF:       applicationTmp,
		ProductID: productID,
	}
	resp := ep.POST("/data/asset").WithHeader("Authorization", "Bearer "+tokenVal).
		WithCookie(jwtCookieSecret, tokenKey).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/asset add result: ", resp.Body())
}

func testEditAsset(t *testing.T) {
	// lg := "logo.jpg"
	lg := ""
	req := &AssetRequest{
		ID:        1353639297386811392,
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
	resp := ep.PUT("/data/asset").WithHeader("Authorization", "Bearer "+tokenVal).
		WithCookie(jwtCookieSecret, tokenKey).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/asset edit result: ", resp.Body())
}

func testGetSample(t *testing.T) {
	req := &SampleRequest{
		ID:    1353639297386811392,
		Table: "area",
	}
	resp := ep.POST("/data/product/sample").WithHeader("Authorization", "Bearer "+tokenVal).
		WithCookie(jwtCookieSecret, tokenKey).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/product/sample result: ", resp.Body())
}
