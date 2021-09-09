package web

import (
	"bytes"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/chu108/tools/network"
	"github.com/chu108/tools/str"
)

//爬虫库

type reptile struct {
	ResBody    []byte
	Doc        *goquery.Document
	Selection  *goquery.Selection
	Err        error
	HttpHeader map[string]string
}

//创建爬虫对象
func NewReptile() *reptile {
	return &reptile{}
}

//设置请求头
func (obj *reptile) SetReqHeader(header map[string]string) *reptile {
	obj.HttpHeader = header
	return obj
}

//获取URL内容
func (obj *reptile) Fetch(url string, data map[string]interface{}) *reptile {
	if obj.Err != nil {
		return obj
	}
	obj.ResBody, obj.Err = network.NewHttp().Get(url, data, obj.HttpHeader)
	return obj
}

//将HTTP返回内容解码
func (obj *reptile) DecodeBodyBase64() *reptile {
	if obj.Err != nil {
		return obj
	}
	//buf := make([]byte, base64.StdEncoding.DecodedLen(len(obj.ResBody)))
	//_, obj.Err = base64.StdEncoding.Decode(buf, obj.ResBody)
	//obj.ResBody = buf

	obj.ResBody, obj.Err = base64.StdEncoding.DecodeString(str.NewBytes().ToStr(obj.ResBody))
	return obj
}

//将HTTP返回内容转换成文档对象
func (obj *reptile) ToDoc() *reptile {
	if obj.Err != nil {
		return obj
	}
	obj.Doc, obj.Err = goquery.NewDocumentFromReader(bytes.NewReader(obj.ResBody))
	return obj
}

//获取查找内容列表
func (obj *reptile) DocList(find string, callBack func(i int, list *goquery.Selection)) *reptile {
	if obj.Doc == nil {
		obj.ToDoc()
	}
	if obj.Err != nil {
		return obj
	}
	obj.Doc.Find(find).Each(func(i int, selection *goquery.Selection) {
		callBack(i, selection)
	})
	return obj
}

//获取查找第一个内容
func (obj *reptile) DocOne(url, find string, callBack func(sel *goquery.Selection)) *reptile {
	if obj.Doc == nil {
		obj.ToDoc()
	}
	if obj.Err != nil {
		return obj
	}
	callBack(obj.Doc.Find(find))
	return obj
}

//运行
func (obj *reptile) Run() error {
	return obj.Err
}
