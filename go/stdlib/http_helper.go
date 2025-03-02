package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Handy struct {
	client *http.Client
	req    *http.Request
	data   any
	err    error
}

func NewHandy() *Handy {
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		panic(err)
	}
	return &Handy{
		client: http.DefaultClient,
		req:    req,
	}
}

func (h *Handy) URL(uri string) *Handy {
	u, err := url.Parse(uri)
	if err != nil {
		h.err = err
	}
	h.req.URL = u
	return h
}

func (h *Handy) Client(client *http.Client) *Handy {
	h.client = client
	return h
}

func (h *Handy) Header(key, value string) *Handy {
	h.req.Header.Add(key, value)
	return h
}

func (h *Handy) Param(key, value string) *Handy {
	query := h.req.URL.Query()
	query.Add(key, value)
	h.req.URL.RawQuery = query.Encode()
	return h
}

func (h *Handy) Form(form map[string]string) *Handy {
	h.req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	data := url.Values{}
	for key, value := range form {
		data.Add(key, value)
	}
	h.req.Body = io.NopCloser(bytes.NewReader([]byte(data.Encode())))
	h.req.ContentLength = int64(len(data.Encode()))
	return h
}

func (h *Handy) JSON(v any) *Handy {
	h.req.Header.Set("Content-Type", "application/json")
	h.req.Header.Add("Accept", "application/json")
	data, err := json.Marshal(v)
	if err != nil {
		h.err = err
	}
	h.req.Body = io.NopCloser(bytes.NewReader(data))
	h.req.ContentLength = int64(len(data))
	return h
}

func (h *Handy) Get() *HandyResponse {
	if h.err != nil {
		return &HandyResponse{err: h.err}
	}
	h.req.Method = http.MethodGet
	resp, err := h.client.Do(h.req)
	if err != nil {
		return &HandyResponse{err: err}
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &HandyResponse{err: err, StatusCode: resp.StatusCode}
	}
	return &HandyResponse{StatusCode: resp.StatusCode, data: data}
}

func (h *Handy) Post() *HandyResponse {
	if h.err != nil {
		return &HandyResponse{err: h.err}
	}
	h.req.Method = http.MethodPost
	resp, err := h.client.Do(h.req)
	if err != nil {
		return &HandyResponse{err: err}
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &HandyResponse{StatusCode: resp.StatusCode, err: err}
	}
	return &HandyResponse{StatusCode: resp.StatusCode, data: data}
}

type HandyResponse struct {
	StatusCode int
	err        error
	data       []byte
}

func (r *HandyResponse) OK() bool {
	return r.StatusCode == http.StatusOK
}

func (r *HandyResponse) Bytes() []byte {
	return r.data
}

// String возвращает тело ответа как строку
func (r *HandyResponse) String() string {
	return string(r.data)
}

func (r *HandyResponse) JSON(v any) {
	err := json.Unmarshal(r.data, v)
	if err != nil {
		r.err = err
	}
}

func (r *HandyResponse) Err() error {
	return r.err
}

func main() {
	{
		// примеры запросов

		// GET-запрос с параметрами
		NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()

		// HTTP-заголовки
		NewHandy().
			URL("https://httpbingo.org/get").
			Header("Accept", "text/html").
			Header("Authorization", "Bearer 1234567890").
			Get()

		// POST формы
		params := map[string]string{
			"brand":    "lg",
			"category": "tv",
		}
		NewHandy().URL("https://httpbingo.org/post").Form(params).Post()

		// POST JSON-документа
		NewHandy().URL("https://httpbingo.org/post").JSON(params).Post()
	}

	{
		// пример обработки ответа

		// отправляем GET-запрос с параметрами
		resp := NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()
		if !resp.OK() {
			panic(resp.String())
		}

		// декодируем ответ в JSON
		var data map[string]any
		resp.JSON(&data)

		fmt.Println(data["url"])
		// "https://httpbingo.org/get"
		fmt.Println(data["args"])
		// map[id:[42]]
	}
}
