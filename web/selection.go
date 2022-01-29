package web

import "github.com/PuerkitoBio/goquery"

const (
	Attr AttrType = "attr" //属性
	Html AttrType = "html" //html
	Text AttrType = "text" //文本

	One  Multiple = "one"  //一个
	Many Multiple = "many" //多个
	Obj  Multiple = "grab" //新对象
)

type (
	AttrType string //属性类型
	Multiple string //数量类型：一个或多个
	// LV 用于列表页
	AR struct {
		AttrType AttrType `json:"attr_type"` //属性类型
		AttrName string   `json:"attr_name"` //属性值
	}

	Sel struct {
		*goquery.Selection
		Err error
	}
)

func NewSel(sel *goquery.Selection) *Sel {
	return &Sel{sel, nil}
}

// Foreach 循环
func (obj *Sel) Foreach(f func(int, *goquery.Selection)) *Sel {
	return NewSel(obj.Each(f))
}

// List 获取列表
// List(LV{Attr, "href"})
func (obj *Sel) List(v AR) (ret []string, err error) {
	if obj.Err != nil {
		return nil, err
	}
	defer obj.Clear()
	ret = make([]string, 0, 10)
	obj.Each(func(i int, selection *goquery.Selection) {
		var tmpVal string
		switch v.AttrType {
		case Attr:
			tmpVal, _ = selection.Attr(v.AttrName)
		case Html:
			tmpVal, err = selection.Html()
		case Text:
			tmpVal = selection.Text()
		}
		if err != nil {
			return
		}
		ret = append(ret, tmpVal)
	})
	return
}

// ListManyAttr 获取列表
// ListManyAttr([]AR{
//   {Attr, "href"}
//   {Attr, "href"}
//})
func (obj *Sel) ListManyAttr(val []AR) (ret []map[string]string, err error) {
	if obj.Err != nil {
		return nil, err
	}
	defer obj.Clear()
	ret = make([]map[string]string, 0, 10)
	obj.Each(func(i int, s *goquery.Selection) {
		tmp := make(map[string]string, 0)
		for _, v := range val {
			switch v.AttrType {
			case Attr:
				tmp[v.AttrName], _ = s.Attr(v.AttrName)
			case Html:
				tmp[v.AttrName], err = s.Html()
			case Text:
				tmp[v.AttrName] = s.Text()
			}
			if err != nil {
				return
			}
		}
		ret = append(ret, tmp)
	})
	return
}

func (obj *Sel) Clear() {
	obj = nil
}
