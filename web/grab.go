package web

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Grab struct {
}

func NewGrab() *Grab {
	return &Grab{}
}

//请求url获取文档对象
func (obj *Grab) Fetch(url string) (doc *goquery.Document, err error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	defer res.Body.Close()

	return goquery.NewDocumentFromReader(res.Body)
}

//获取url查找内容列表
func (obj *Grab) FindList(url, find string, callBack func(i int, list *goquery.Selection)) error {
	doc, err := obj.Fetch(url)
	if err != nil {
		return err
	}
	doc.Find(find).Each(func(i int, selection *goquery.Selection) {
		callBack(i, selection)
	})
	return nil
}

//获取url查找第一个内容
func (obj *Grab) FindOne(url, find string, callBack func(sel *goquery.Selection)) error {
	doc, err := obj.Fetch(url)
	if err != nil {
		return err
	}
	callBack(doc.Find(find))
	return nil
}
