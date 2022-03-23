package _web

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

type Grab struct {
	Http  *http.Client
	Debug bool
}

func NewGrab() *Grab {
	return &Grab{&http.Client{}, false}
}

func NewGrabByHttp(http *http.Client) *Grab {
	return &Grab{http, false}
}

func (obj *Grab) SetHttpClient(http *http.Client) {
	obj.Http = http
}

// Fetch 请求url获取文档对象
func (obj *Grab) Fetch(url string) *Doc {
	res, err := obj.Http.Get(url)
	if obj.Debug == true {
		body, _ := io.ReadAll(res.Body)
		fmt.Println("----- req body ---------------------")
		fmt.Println(body)
		fmt.Println("----- req err ---------------------")
		fmt.Println(err)
		fmt.Println("----- end ---------------------")
	}
	if err != nil {
		log.Fatal("Fetch.get.err: ", err)
		return NewDocErr(nil, err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Fetch.StatusCode.err: ", res.StatusCode, err)
		return NewDocErr(nil, errors.New(fmt.Sprintf("Fetch StatusCode:%d, err:%v", res.StatusCode, err)))
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Fetch.NewDocumentFromReader.err: ", err)
		return NewDocErr(nil, err)
	}
	return NewDoc(doc)
}
