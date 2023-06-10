package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Get(reqUrl string, param map[string]string) ([]byte, error) {
	reqUrl += "?"
	for k, v := range param {
		reqUrl += k + "=" + url.QueryEscape(v) + "&"
	}
	reqUrl = strings.Trim(reqUrl, "&")
	fmt.Println("GET请求参数：", reqUrl)
	resp, err0 := http.Get(reqUrl)
	if err0 != nil {
		return nil, err0
	}
	defer resp.Body.Close()
	result, err0 := ioutil.ReadAll(resp.Body)
	if err0 != nil {
		return nil, err0
	}
	fmt.Println("GET返回结果：", string(result))
	return result, nil
}
