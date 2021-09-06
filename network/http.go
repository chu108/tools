package network

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/chu108/tool2/files"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	POST = "POST"
	GET  = "GET"
)

type Http struct {
}

func NewHttp() *Http {
	return &Http{}
}

//获取client
func (*Http) getClient() *http.Client {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	t.MaxIdleConns = 100 //连接池最大连接数量
	// t.MaxConnsPerHost = 50     //每个host的最大连接数量，0表示不限制
	t.MaxIdleConnsPerHost = 10 //每个host的连接池最大空闲连接数,默认2
	return &http.Client{
		Timeout:   3 * time.Second, //超时为3秒
		Transport: t,
	}
}

//http get请求
func (obj *Http) Get(url string, data map[string]interface{}, header map[string]string) ([]byte, error) {
	return obj.Request(GET, url, data, header)
}

//http post请求
func (obj *Http) Post(url string, data map[string]interface{}, header map[string]string) ([]byte, error) {
	return obj.Request(POST, url, data, header)
}

//http 通用请求
func (obj *Http) Request(method, requestUrl string, data map[string]interface{}, header map[string]string) ([]byte, error) {
	method = strings.ToUpper(method)
	client := obj.getClient()
	var (
		req  *http.Request
		err  error
		body io.Reader = nil
	)

	dataLen := len(data)
	switch method {
	case POST:
		if dataLen > 0 {
			bytesData, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = bytes.NewReader(bytesData)
		}
		req, err = http.NewRequest(POST, requestUrl, body)
	case GET:
		if dataLen > 0 {
			params := url.Values{}
			for key, val := range data {
				if value, ok := val.(string); ok {
					params.Add(key, value)
				}
				if value, ok := val.(int); ok {
					params.Add(key, strconv.Itoa(value))
				}
			}
			URL, err := url.Parse(requestUrl)
			if err != nil {
				return nil, err
			}
			URL.RawQuery = params.Encode()
			requestUrl = URL.String()
		}
		req, err = http.NewRequest(GET, requestUrl, nil)
	}
	if err != nil {
		return nil, err
	}

	if len(header) > 0 {
		for key, val := range header {
			if key == "cookie" {
				obj.setCookie(req, val)
			}
			req.Header.Add(key, val)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

//设置cookie
func (*Http) setCookie(req *http.Request, cookie string) {
	cookie = strings.TrimSpace(cookie)
	cks := strings.Split(cookie, ";")
	for _, v := range cks {
		item := strings.Split(v, "=")
		cookieItem := &http.Cookie{Name: item[0], Value: url.QueryEscape(item[1])}
		req.AddCookie(cookieItem)
	}
}

/**
获取Url地址重写向地址
*/
func (*Http) GetUrlRedirect(url string) (*url.URL, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("redirect")
	}

	response, err := client.Do(req)
	if err != nil {
		if response != nil && response.StatusCode == http.StatusFound { //status code 302
			return response.Location()
		} else {
			return nil, err
		}
	}
	return nil, nil
}

/**
下载文件
url：下载地址
savePath：保存路径(包含文件名)
*/
func (obj *Http) DownloadFile(url, savePath string) (string, error) {
	if url == "" || savePath == "" {
		return "", errors.New("下载url或保存地址错误")
	}

	//判断文件是否存在，存在则返回
	if files.NewFile().IsExist(savePath) && files.NewFile().FileSize(savePath) > 0 {
		return "", nil
	}

	body, err := obj.GetUrlBody(url)
	if err != nil {
		return "", err
	}
	if len(body) == 0 {
		return "", err
	}

	file, err := os.Create(savePath)
	if err != nil {
		return "", err
	}

	writerLen, err := io.Copy(file, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	//如果写入内容为空，则删除空文件
	if writerLen == 0 {
		err = os.Remove(savePath)
		if err != nil {
			return "", err
		}
	}

	return savePath, nil
}

/**
获取url返回的内容
*/
func (*Http) GetUrlBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
