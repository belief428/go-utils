package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type (
	// Headers 消息头
	Headers struct {
		UserAgent   string
		ContentType string
		Cookies     map[string]string
		Others      map[string]string
	}

	// Method 请求方式
	Method string

	// Client
	Client struct {
		Url    string
		Method Method
		Params map[string]interface{}
	}
)

const (
	MethodForGet  Method = "GET"
	MethodForPost Method = "POST"

	DefaultUserAgent string = "Mozilla/5.0 (SF) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.6.6666.66 Safari/537.36"
)

// RequestBodyFormat 请求消息内容格式
type RequestBodyFormat int

const (
	RequestBodyFormatForFormData RequestBodyFormat = iota + 1
	RequestBodyFormatForXWWWFormUrlencoded
	RequestBodyFormatForRaw
)

const (
	RequestContentTypeForFormData           string = "application/form-data"
	RequestContentTypeForXWWWFormUrlencoded string = "application/x-www-form-urlencoded"
)

// Request 发起请求
func (this *Client) Request(format RequestBodyFormat, headers ...Headers) ([]byte, error) {
	client := new(http.Client)

	var reqBody io.Reader

	if this.Method == MethodForGet {
		_params := make([]string, 0)

		for k, v := range this.Params {
			_params = append(_params, fmt.Sprintf("%s=%v", k, v))
		}
		this.Url += "?" + strings.Join(_params, "&")
	} else {
		if format == RequestBodyFormatForFormData || format == RequestBodyFormatForXWWWFormUrlencoded {
			_params := make([]string, 0)
			for k, v := range this.Params {
				_params = append(_params, fmt.Sprintf("%s=%v", k, v))
			}
			reqBody = strings.NewReader(strings.Join(_params, "&"))
		} else if format == RequestBodyFormatForRaw {
			_bytes, _ := json.Marshal(this.Params)
			reqBody = bytes.NewReader(_bytes)
		}
	}
	req, err := http.NewRequest(string(this.Method), this.Url, reqBody)

	if err != nil {
		return nil, err
	}
	for _, v := range headers {
		if v.UserAgent != "" {
			req.Header.Add("User-Agent", v.UserAgent)
		}
		if v.ContentType != "" {
			req.Header.Add("Content-Type", v.ContentType)
		}
		if len(v.Cookies) > 0 {
			for key, val := range v.Cookies {
				req.AddCookie(&http.Cookie{Name: key, Value: val})
			}
		}
		if len(v.Others) > 0 {
			for key, val := range v.Others {
				req.Header.Add(key, val)
			}
		}
	}
	resp := new(http.Response)

	if resp, err = client.Do(req); err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return bytes, err
}

// NewClient
func NewClient(url string, method Method, params map[string]interface{}) *Client {
	return &Client{
		Url: url, Method: method, Params: params,
	}
}
