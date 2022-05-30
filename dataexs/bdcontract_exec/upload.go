package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

/*
curl -H "Content-Type:multipart/form-data" -F "file=@/Users/luxiaotong/code/go_practice/dataexs/bdcontract_exec/postgresql-9.1-901-1.jdbc4.jar" \
"http://127.0.0.1:58150/Upload?path=/Datassets/Datassets.yjs&fileName=postgresql-9.1-901-1.jdbc4.jar&isPrivate=false&order=10&count=11&pubKey=04395d3c655637c4f06cbeabc1f061815a1e85fedd2c6f6d23afba9412e2bd8bf387786cbd659b588bca646ccc7076caf31923749db51a3f966c4e28594a9be0e5&sign=3046022100e908b1e87d1b927098ab54c4ba3d23bd158d0f5c5545094b1efca3d7c5c604f0022100bf7a74d8b8990e2bbeffe6101608e9dcedc4501858d5d45a3e35ef4fba983e20"
need timeout
*/

func UploadFile(privKey, pubKey, name, path string) error {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile("file", "blob")
	if err != nil {
		log.Fatalf("create form file failed: %s", err)
		return err
	}
	srcFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file failed: %s", err)
		return err
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Fatalf("write to form file falied: %s", err)
		return err
	}
	writer.Close()

	contentType := writer.FormDataContentType()
	url := "http://127.0.0.1:58150/Upload"
	param := "path=/Datassets/Datassets.yjs&fileName=" + name + "&isPrivate=false&order=0&count=1" + "&pubKey=" + pubKey
	signature, err := SignSM2KeyPairHex(privKey, param)
	if err != nil {
		log.Fatalf("sign upload file falied: %s", err)
		return err
	}
	url += "?" + param + "&sign=" + signature
	log.Printf("url: %v", url)
	req, _ := http.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", contentType)
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("upload file failed: %s", err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read resp failed: %s", err)
		return err
	}
	log.Printf("upload file response, status: %v, body %v", resp.StatusCode, string(body))
	return nil
}
