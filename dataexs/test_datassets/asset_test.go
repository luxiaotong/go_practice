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
