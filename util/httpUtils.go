package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

func Request(req *http.Request, timeout time.Duration) (body []byte, err error) {
	client := &http.Client{
		Timeout: timeout,
	}
	beego.Debug("Request URL : ", req.URL)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func RequestJSON(req *http.Request, timeout time.Duration) (result map[string]interface{}, err error) {
	body, err := Request(req, timeout)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetJSON(url string, timeout time.Duration) (result map[string]interface{}, err error) {
	req, err := http.NewRequest("get", url, nil)
	if err != nil {
		return nil, err
	}
	return RequestJSON(req, timeout)
}

func main() {
	resp, err := GetJSON("https://api.weixin.qq.com/sns/", 60*time.Second)
	if err != nil {
		fmt.Printf("err : s%", err)
	}
	fmt.Println("resp code :", resp["errcode"])
}
