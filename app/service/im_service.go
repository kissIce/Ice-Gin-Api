package service

import (
	"bytes"
	"compress/zlib"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"ice/utils/response"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	baseUri = "https://console.tim.qq.com/v4"
	appId   = 0
	secret  = ""
)

type ImInfo struct {
	Identifier string
	Nick       string
	FaceUrl    string
}

func ImReg(info *ImInfo) (bool, int) {
	str, _ := json.Marshal(info)
	req, err := http.NewRequest("POST",
		getRequestUrl("im_open_login_svc/account_import"),
		bytes.NewBuffer(str),
	)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, response.ThirdError
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, response.ThirdError
	}
	fmt.Println(string(body))
	return true, 0
}

func getRequestUrl(uri string) string {
	rand.Seed(time.Now().UnixNano())
	return baseUri + "/" + uri + "?sdkappid=" + strconv.FormatInt(appId, 10) + "&identifier=Ice" + "&usersig=" + CreateSign("Ice") + "&random=" + strconv.Itoa(rand.Intn(4294967295)) + "&contenttype=json"
}

func CreateSign(id string) string {
	timestamp := time.Now().Unix()
	expire := 15552000
	var sigDoc map[string]interface{}
	sigDoc = make(map[string]interface{})
	sigDoc["TLS.ver"] = "2.0"
	sigDoc["TLS.identifier"] = id
	sigDoc["TLS.sdkappid"] = appId
	sigDoc["TLS.expire"] = expire
	sigDoc["TLS.time"] = timestamp
	sigDoc["TLS.sig"] = hmacsha256(id, timestamp, expire, "", false)
	data, _ := json.Marshal(sigDoc)
	fmt.Println(data)
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(data)
	w.Close()
	return base64urlEncode(b.Bytes())
}

func hmacsha256(id string, timestamp int64, expire int, userBuf string, userBufEnable bool) string {
	var str string
	str = "TLS.identifier:" + id + "\n"
	str += "TLS.sdkappid:" + strconv.Itoa(appId) + "\n"
	str += "TLS.time:" + strconv.FormatInt(timestamp, 10) + "\n"
	str += "TLS.expire:" + strconv.Itoa(expire) + "\n"
	if userBufEnable {
		str += "TLS.userbuf:" + userBuf + "\n"
	}
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func base64urlEncode(data []byte) string {
	str := base64.StdEncoding.EncodeToString(data)
	str = strings.Replace(str, "+", "*", -1)
	str = strings.Replace(str, "/", "-", -1)
	str = strings.Replace(str, "=", "_", -1)
	return str
}

func base64urlDecode(str string) ([]byte, error) {
	str = strings.Replace(str, "_", "=", -1)
	str = strings.Replace(str, "-", "/", -1)
	str = strings.Replace(str, "*", "+", -1)
	return base64.StdEncoding.DecodeString(str)
}
