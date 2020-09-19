package sms

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	smsURL     = "https://rtcsms.cn-north-1.myhuaweicloud.com:10743/sms/batchSendSms/v1"
	appKey     = "E6q93B3gv818f5Z6vsW1Pid6sraK"
	appSecret  = "734WE0h031mgJa402nsQaHFxP194"
	sender     = "10690400999305346"
	templateID = "b146927de4f947879f586db0d88ed275"
	signature  = "华为云短信测试"
)

func buildWSSE() (string, error) {
	now := time.Now().Format("2006-01-02T15:04:05Z")
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	nonce := strings.ReplaceAll(u.String(), "-", "")
	sum := sha256.Sum256([]byte(nonce + now + appSecret))
	digest := base64.StdEncoding.EncodeToString(sum[:])
	wwse := fmt.Sprintf("UsernameToken Username=\"%s\",PasswordDigest=\"%s\",Nonce=\"%s\",Created=\"%s\"", appKey, digest, nonce, now)
	return wwse, nil
}

func GenCode() int32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
}

// SendCode 发送短信验证码
// param: string receiver 接收方手机号, 格式:+86XXXXXXXXXXX
// param: string code 6位验证码
// return:
func SendCode(receiver string, vcode int32) error {
	wsse, err := buildWSSE()
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("from", sender)
	data.Set("to", receiver)
	data.Set("templateId", templateID)
	data.Set("templateParas", fmt.Sprintf("[\"%06d\"]", vcode))
	request, err := http.NewRequest("POST", smsURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", `WSSE realm="SDP",profile="UsernameToken",type="Appkey"`)
	request.Header.Set("X-WSSE", wsse)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   time.Duration(5 * time.Second),
		Transport: tr,
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var r map[string]interface{}
	if err := json.Unmarshal(body, &r); err != nil {
		return err
	}
	if r["code"] != "000000" {
		return errors.Errorf("got an error for sending code: %v", r)
	}
	fmt.Println("response: ", r)
	return nil
}

// SetCode 存储短信验证码
// param: string receiver 接收方手机号, 格式:+86XXXXXXXXXXX
// param: string code 6位验证码
// return:
func SetCode(receiver string, vcode int32) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "139.9.119.21:56379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	k := fmt.Sprintf("smslogin:%s", receiver)
	if err := rdb.Set(context.Background(), k, vcode, 60*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

// CheckCode 校验短信验证码
// param: string receiver 接收方手机号, 格式:+86XXXXXXXXXXX
// return:
func CheckCode(receiver string, vcode int32) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "139.9.119.21:56379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	k := fmt.Sprintf("smslogin:%s", receiver)
	val, err := rdb.Get(context.Background(), k).Result()
	if err == redis.Nil {
		return errors.Errorf("key does not exist")
	} else if err != nil {
		return err
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	if int32(v) != vcode {
		return errors.Errorf("vcode does not match")
	}
	fmt.Println("vcode: ", v)
	return nil
}
