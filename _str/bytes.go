package _str

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"unsafe"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

type Bytes struct {
}

func NewBytes() *Bytes {
	return &Bytes{}
}

//byte转字符串
func (*Bytes) ToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//byte转字符串指定字符集
func (*Bytes) ToStrByCharset(byte []byte, charset Charset) (ret string) {
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		ret = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		ret = string(byte)
	}
	return
}
