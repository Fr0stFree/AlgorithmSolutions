package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type StatusErr struct {
	Code   int
	Status string
}

func (e StatusErr) Error() string {
	return "invalid response status: " + e.Status
}

func httpGet(uri string, headers map[string]string, params map[string]string, timeout int) (map[string]any, error) {
	client := http.Client{Timeout: time.Duration(timeout) * time.Millisecond}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, StatusErr{resp.StatusCode, resp.Status}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := make(map[string]any)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func main() {
	{
		// GET-запрос
		const uri = "https://httpbingo.org/json"
		data, err := httpGet(uri, nil, nil, 3000)
		fmt.Printf("GET %v\n", uri)
		fmt.Println(data, err)
		fmt.Println()
		// GET https://httpbingo.org/json
		// map[slideshow:map[author:Yours Truly date:date of publication slides:[map[title:Wake up to WonderWidgets! type:all] map[items:[Why <em>WonderWidgets</em> are great Who <em>buys</em> WonderWidgets] title:Overview type:all]] title:Sample Slide Show]] <nil>
	}

	{
		// 404 Not Found
		const uri = "https://httpbingo.org/whatever"
		data, err := httpGet(uri, nil, nil, 3000)
		fmt.Printf("GET %v\n", uri)
		fmt.Println(data, err)
		fmt.Println()
		// GET https://httpbingo.org/whatever
		// map[] invalid response status: 404 Not Found
	}

	{
		// С заголовками
		const uri = "https://httpbingo.org/headers"
		headers := map[string]string{
			"accept": "application/xml",
			"host":   "httpbingo.org",
		}
		data, err := httpGet(uri, headers, nil, 3000)
		fmt.Printf("GET %v\n", uri)
		respHeaders := data["headers"].(map[string]any)
		fmt.Println(respHeaders["Accept"], respHeaders["Host"], err)
		fmt.Println()
		// GET https://httpbingo.org/headers
		// [application/xml] [httpbingo.org] <nil>
	}

	{
		// С URL-параметрами
		const uri = "https://httpbingo.org/get"
		params := map[string]string{"id": "42"}
		data, err := httpGet(uri, nil, params, 3000)
		fmt.Printf("GET %v\n", uri)
		fmt.Println(data["args"], err)
		fmt.Println()
		// GET https://httpbingo.org/get
		// map[id:[42]] <nil>
	}
}
