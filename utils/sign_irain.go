package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"encoding/base64"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"crypto"
)


// IrainSign 第三方无感支付 API 签名
func IrainSign(data interface{}, key string) string {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}
	v := reflect.ValueOf(data)
	fieldNum := t.NumField()
	fieldKeys := []string{}
	fieldMap := make(map[string]interface{})
	for i := 0; i < fieldNum; i++ {
		signTag := t.Field(i).Tag.Get("sign")
		if signTag != "true" {
			continue
		}
		url := t.Field(i).Tag.Get("url")
		value := v.Field(i).Interface()
		fieldKeys = append(fieldKeys, url)
		fieldMap[url] = value
	}
	sort.Strings(fieldKeys)

	preRet := ""
	for _, fieldKey := range fieldKeys {
		preRet = fmt.Sprintf("%v%v", preRet, fieldMap[fieldKey])
	}
	preRet += key
	h := md5.New()
	h.Write([]byte(preRet))

	return hex.EncodeToString(h.Sum(nil))
}

// 创泰签名算法
func Sign(privateKey string, data []byte) (string, error) {
	kData, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	pkCs8PrivateKey, err := x509.ParsePKCS8PrivateKey(kData)
	if err != nil {
		return "", err
	}
	md5 := md5.New()
	md5.Write(data)
	h := md5.Sum(nil)

	buf, err := rsa.SignPKCS1v15(rand.Reader, pkCs8PrivateKey.(*rsa.PrivateKey), crypto.MD5, h)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf),nil
}
