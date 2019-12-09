package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// get请求
func Get(url string) (res interface{}, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	str := string(body)
	fmt.Println(string(body))

	err = json.Unmarshal([]byte(str), &res)
	return res, err
}

// post请求
func Post(url string, data interface{}, contentType string) (res interface{}, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	content := string(result)
	err = json.Unmarshal([]byte(content), &res)
	return res, err
}
