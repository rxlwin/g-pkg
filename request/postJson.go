package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func PostJson(reqUrl string, param map[string]interface{}) ([]byte, error) {
	//return nil, nil
	client := &http.Client{Timeout: 10 * time.Second}
	data, _ := json.Marshal(param)
	fmt.Println("postJson请求参数：", string(data))
	req, err0 := http.NewRequest("POST", reqUrl, bytes.NewReader(data))
	if err0 != nil {
		return nil, err0
	}
	req.Header.Add("Content-Type", "application/json")
	res0, err0 := client.Do(req)
	if err0 != nil {
		return nil, err0
	}
	defer res0.Body.Close()
	result, err0 := ioutil.ReadAll(res0.Body)
	if err0 != nil {
		return nil, err0
	}
	fmt.Println("postJson返回结果：", string(result))
	return result, nil
}

func PostJsonWithHeader(reqUrl string, headers map[string]string, param map[string]interface{}) ([]byte, error) {
	//return nil, nil
	client := &http.Client{Timeout: 10 * time.Second}
	data, _ := json.Marshal(param)
	fmt.Println("postJson请求参数：", string(data))
	fmt.Println("postJson Header参数：", headers)
	req, err0 := http.NewRequest("POST", reqUrl, bytes.NewReader(data))
	if err0 != nil {
		return nil, err0
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res0, err0 := client.Do(req)
	if err0 != nil {
		return nil, err0
	}
	defer res0.Body.Close()
	result, err0 := ioutil.ReadAll(res0.Body)
	if err0 != nil {
		return nil, err0
	}
	fmt.Println("postJson返回结果：", string(result))
	return result, nil
}
