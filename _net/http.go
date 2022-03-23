package _net

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chu108/tools/_file"
	"golang.org/x/net/proxy"
	"golang.org/x/net/publicsuffix"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
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
	Trans  *http.Transport
	Client *http.Client
	Errs   error
}

func NewHttp() *Http {
	return &Http{
		Trans: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			MaxIdleConns:        100,             //连接池最大连接数量
			MaxIdleConnsPerHost: 10,              //每个host的连接池最大空闲连接数,默认2
			TLSHandshakeTimeout: 5 * time.Second, //限制 TLS握手的时间
		},
		Client: &http.Client{
			Timeout: 3 * time.Second, //超时为3秒
		},
	}
}

//设置cookie
func (obj *Http) SetCookieJar(urlStr string, cookies []*http.Cookie) *Http {
	cookiejarOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, _ := cookiejar.New(&cookiejarOptions)
	if urlStr != "" && cookies != nil && len(cookies) > 0 {
		proxyUrl, err := url.Parse(urlStr)
		if err != nil {
			obj.Errs = err
			return nil
		}
		jar.SetCookies(proxyUrl, cookies)
	}
	obj.Client.Jar = jar
	return obj
}

//设置Http代理
//全局设置：
//os.Setenv("HTTP_PROXY", "http://127.0.0.1:9743")
//os.Setenv("HTTPS_PROXY", "https://127.0.0.1:9743")
func (obj *Http) SetProxyHttp(ip string, port int64) *Http {
	host := fmt.Sprintf("http://%s:%d", ip, port)
	proxyUrl, err := url.Parse(host)
	if err != nil {
		obj.Errs = err
		return obj
	}
	obj.Trans.Proxy = http.ProxyURL(proxyUrl)
	return obj
}

//设置Socks代理
func (obj *Http) SetProxySocks(ip string, port int64, name, password string) *Http {
	host := fmt.Sprintf("%s:%d", ip, port)
	var auth *proxy.Auth
	if name != "" {
		auth = &proxy.Auth{User: name, Password: password}
	}
	//forward := &net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}
	dialer, err := proxy.SOCKS5("tcp", host, auth, proxy.Direct)
	if err != nil {
		obj.Errs = err
		return obj
	}
	obj.Trans.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}
	return obj
}

//GetClient
func (obj *Http) GetClient() *http.Client {
	obj.Client.Transport = obj.Trans
	return obj.Client
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
	if obj.Errs != nil {
		return nil, obj.Errs
	}
	method = strings.ToUpper(method)
	client := obj.GetClient()
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

	//fmt.Println("req cookie:", obj.Client.Jar.Cookies(req.URL))
	//fmt.Println("res cookie:", resp.Cookies())

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
	if _file.NewFile().IsExist(savePath) && _file.NewFile().FileSize(savePath) > 0 {
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
