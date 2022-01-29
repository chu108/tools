package web

import "github.com/PuerkitoBio/goquery"

type (
	FO struct {
		Selector string             `json:"selector"` //选择器
		AttrType `json:"attr_type"` //属性类型
		AttrName string             `json:"attr_name"` //属性值
	}

	FM struct {
		Key string `json:"key"`
		FO  `json:"eo"`
	}

	KV map[string]string

	Doc struct {
		*goquery.Document
		Err error
	}
)

func NewDoc(doc *goquery.Document) *Doc {
	return &Doc{doc, nil}
}

// Find 根据规则查找
func (obj *Doc) Find(selector string) *Sel {
	return NewSel(obj.Document.Find(selector))
}

// FindOneValues 获取匹配的值数组
func (obj *Doc) FindOneValues(val FO) (ret []string, err error) {
	if obj.Err != nil {
		return nil, err
	}
	defer obj.Clear()
	return obj.Find(val.Selector).List(AR{val.AttrType, val.AttrName})
}

// FindMany 获取多个匹配的值，返回键值对
func (obj *Doc) FindMany(val []FO) (retMap KV, err error) {
	if obj.Err != nil {
		return nil, err
	}
	defer obj.Clear()
	retMap = make(KV, 0)
	for _, v := range val {
		switch v.AttrType {
		case Attr:
			retMap[v.AttrName], _ = obj.Find(v.Selector).Attr(v.AttrName)
		case Html:
			retMap[v.AttrName], err = obj.Find(v.Selector).Html()
		case Text:
			retMap[v.AttrName] = obj.Find(v.Selector).Text()
		}
		if err != nil {
			return nil, err
		}
	}
	return
}

// FindManyValues 获取多个匹配的值，返回键与值数组
//func (obj *Doc) FindManyValues(val []FM) (ret map[string][]string, err error) {
//	if obj.Err != nil {
//		return nil, err
//	}
//	defer obj.Clear()
//	ret = make(map[string][]string, 0)
//	for _, v := range val {
//		ret[v.Key], err = obj.FindOneValues(FO{v.Selector, v.AttrType, v.AttrName})
//		if err != nil {
//			return nil, err
//		}
//	}
//	return
//}

func (obj *Doc) Clear() {
	obj = nil
}
