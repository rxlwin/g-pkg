package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func PostUrlEncoded(reqUrl string, param map[string]string) (string, error) {
	client := &http.Client{Timeout: 2 * time.Second}
	data := url.Values{}
	for k, v := range param {
		data.Add(k, v)
	}
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
