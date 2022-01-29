package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Grab struct {
}

func NewGrab() *Grab {
	return &Grab{}
}

// Fetch 请求url获取文档对象
func (obj *Grab) Fetch(url string) *Doc {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic(fmt.Sprintf("Fetch StatusCode:%d, err:%v", res.StatusCode, err))
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	return NewDoc(doc)
}
