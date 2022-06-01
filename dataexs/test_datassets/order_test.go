package testdatassets

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

const (
	statusWaitSeller = 10
	statusUnpaid     = 20
	statusWaitExec   = 45
)

const contract = `import "postgres.yjs";
contract Data{
  export function main(arg) {
    var args = JSON.parse(arg);
    var conn = PGConnect(args.host, args.port, args.db, args.user, args.pass);
    var  sql ="select count(1) from area_statistics where level=3";
    var statement = conn.createStatement();
    var resultSet = statement.executeQuery(sql);
    ret  = [];
    var meta = resultSet.getMetaData();
    for (;resultSet.next();){
      var line = {
      };
      for (var j=1;j<=meta.getColumnCount();j++){
        line[meta.getColumnName(j)] = resultSet.getString(j);
      }
      ret.push(line);
    }
    return JSON.stringify(ret);
  }
}`

type OrderRequest struct {
	ID        int64   `json:"id,string"`
	ProductID int64   `json:"product_id,string"`
	Contract  string  `json:"contract"`
	Node      string  `json:"node"`
	Usage     string  `json:"usage"`
	DRs       int64   `json:"drs"`
	Discount  float32 `json:"discount"`
	Role      int     `json:"role"`
	Status    int32   `json:"status"`
	Refuse    string  `json:"refuse"`
}

type SearchOrderRequest struct {
	Query     string `json:"q"`
	Role      int    `json:"role"`
	Recent    int    `json:"recent"`
	Over      int    `json:"over"`
	Status    int32  `json:"status"`
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
	ProductID int64  `json:"product_id,string"`
}

type Payment struct {
	No       string  `json:"no"`
	Account  string  `json:"account"`
	Amount   float32 `json:"amount"`
	BankAddr string  `json:"bank_addr"`
	Voucher  string  `json:"voucher"`
	PayTime  string  `json:"pay_time"`
}

type PayRequest struct {
	ID int64 `json:"id,string"`
	*Payment
}

type AuditOrderRequest struct {
	ID       int64  `json:"id,string"`
	Approved bool   `json:"approved"`
	Refuse   string `json:"refuse"`
}

func testAddOrder(t *testing.T) {
	// productID = 1499628438057652224
	req := &OrderRequest{
		ProductID: productID,
		Usage:     "test usage",
		Node:      trustURL,
		Contract:  contract,
		DRs:       100,
	}
	resp := ep.POST("/data/order").WithHeader("Authorization", "Bearer "+tokenValBuyer).
		WithCookie(jwtCookieSecret, tokenKeyBuyer).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/order add result: ", resp.Body())
}

func testGetOrders_WaitSeller(t *testing.T) {
	req := &SearchOrderRequest{
		Role:      2,
		ProductID: productID,
		Status:    statusWaitSeller,
	}
	resp := ep.POST("/data/orders").WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/orders result: ", resp.Body())
	oid := resp.JSON().Object().Value("data").Object().Value("list").Array().First().Object().Value("id").String().Raw()
	orderID, _ = strconv.ParseInt(oid, 10, 64)
	fmt.Println("latest order id: ", orderID)
}

func testGetOrders_WaitExec(t *testing.T) {
	// productID = 1531485725898313728
	req := &SearchOrderRequest{
		Role:      2,
		ProductID: productID,
		Status:    statusWaitExec,
	}
	resp := ep.POST("/data/orders").WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/orders result: ", resp.Body())
	oid := resp.JSON().Object().Value("data").Object().Value("list").Array().First().Object().Value("id").String().Raw()
	orderID, _ = strconv.ParseInt(oid, 10, 64)
	fmt.Println("latest order id: ", orderID)
}

func testOpOrder_SellerConfirm(t *testing.T) {
	req := &OrderRequest{
		ID:     orderID,
		Role:   2,
		Status: statusUnpaid,
	}
	resp := ep.POST("/data/order/status").WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("data/order/status result: ", resp.Body())
}

func testPayOrder(t *testing.T) {
	// orderID = 1501119917003378688
	req := &PayRequest{
		ID: orderID,
		Payment: &Payment{
			No:       "xxxxxxxxxxxx",
			Account:  "xxxxxxxxxxxxx",
			Amount:   100,
			BankAddr: "招商银行xxxxx分行",
			Voucher:  voucherTmp,
			PayTime:  "2020-05-06T00:08:50.00+08:00",
		},
	}
	resp := ep.POST("/data/pay").WithHeader("Authorization", "Bearer "+tokenValBuyer).
		WithCookie(jwtCookieSecret, tokenKeyBuyer).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/data/pay result: ", resp.Body())
}

func testAuditOrder(t *testing.T) {
	// orderID = 1501119917003378688
	req := &AuditOrderRequest{
		ID:       orderID,
		Approved: true,
	}
	resp := eb.POST("/data/order/status").
		WithCookie(backCookie, provUserToken).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	resp.JSON().Object().Value("err_code").Equal(0)
	fmt.Println("/data/order/status result: ", resp.Body())
}

func testOpOrder_BuyerConfirm(t *testing.T) {
	// orderID = 1501119917003378688
	req := &OrderRequest{
		ID:     orderID,
		Role:   1,
		Status: statusWaitExec,
	}
	resp := ep.POST("/data/order/status").WithHeader("Authorization", "Bearer "+tokenValBuyer).
		WithCookie(jwtCookieSecret, tokenKeyBuyer).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("data/order/status result: ", resp.Body())
}
