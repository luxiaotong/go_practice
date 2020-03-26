package main

import (
    "github.com/tjfoc/gmsm/sm3"
    "github.com/tjfoc/gmsm/sm2"
    "github.com/tjfoc/gmsm/sm4"

    "fmt"
    "log"
    "bytes"
    "crypto/cipher"
)

func main() {
    //sm3
    data := "test"
    h := sm3.New()
    h.Write([]byte(data))
    sum := h.Sum(nil)
    fmt.Printf("origin value is: %s\n", data)
    fmt.Printf("digest value is: %x\n", sum)

    fmt.Println()

    //sm2
    priv, err := sm2.GenerateKey() // 生成密钥对
    if err != nil {
        log.Fatal(err)
    }
    msg := []byte("Tongji Fintech Research Institute")
    pub := &priv.PublicKey
    ciphertxt, err := pub.Encrypt(msg)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("加密结果:%x\n",ciphertxt)
    plaintxt,err :=  priv.Decrypt(ciphertxt)
    if err != nil {
        log.Fatal(err)
    }
    if !bytes.Equal(msg,plaintxt){
        log.Fatal("原文不匹配")
    }

    r,s,err := sm2.Sign(priv, msg)
    if err != nil {
        log.Fatal(err)
    }
    isok := sm2.Verify(pub,msg,r,s)
    fmt.Printf("Verified: %v\n", isok)

    fmt.Println()

    //sm4
    key := []byte("1234567890abcdef")
    // 128比特iv
    iv := make([]byte, sm4.BlockSize)
    data4 := []byte("Tongji Fintech Research Institute")
    fmt.Printf("加密前: %s\n", data4)
    ciphertxt4,err := sm4Encrypt(key,iv, data4)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("加密结果: %x\n", ciphertxt4)

    origintxt,err := sm4Decrypt(key, iv, ciphertxt4)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("解密结果: %s\n", origintxt)
}

func sm4Encrypt(key, iv, plainText []byte) ([]byte, error) {
    block, err := sm4.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    origData := pkcs5Padding(plainText, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, iv)
    cryted := make([]byte, len(origData))
    blockMode.CryptBlocks(cryted, origData)
    return cryted, nil
}

func sm4Decrypt(key, iv, cipherText []byte) ([]byte, error) {
    block, err := sm4.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockMode := cipher.NewCBCDecrypter(block, iv)
    origData := make([]byte, len(cipherText))
    blockMode.CryptBlocks(origData, cipherText)
    origData = pkcs5UnPadding(origData)
    return origData, nil
}
// pkcs5填充
func pkcs5Padding(src []byte, blockSize int) []byte {
    padding := blockSize - len(src)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(src, padtext...)
}

func pkcs5UnPadding(src []byte) []byte {
    length := len(src)
    unpadding := int(src[length-1])
    return src[:(length - unpadding)]
}
