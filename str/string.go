package str

import (
	"bytes"
	"github.com/chu108/tools/config"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

type Str struct {
}

func NewStr() *Str {
	return &Str{}
}

//字符串转byte
func (*Str) ToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//高效截取字符串
func (*Str) SubStr(str string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(str); i++ {
		_, size = utf8.DecodeRuneInString(str[n:])
		n += size
	}
	return str[:n]
}

//字符串替换
func (*Str) ReplaceEmpty(str string, rep ...string) string {
	for i := 0; i < len(rep); i++ {
		str = strings.Replace(str, rep[i], "", -1)
	}
	return str
}

//字符串正则替换
func (*Str) ReplaceRegexpStrEmpty(str string, math ...string) string {
	for i := 0; i < len(math); i++ {
		rep := regexp.MustCompile(math[i])
		str = strings.Replace(str, rep.FindString(str), "", -1)
	}
	return str
}

//字符串拆分，按字数
func (*Str) SplitByNum(txt string, length int) []string {
	txtRune := []rune(txt)
	txtLen := len(txtRune)
	retTxt := make([]string, 0, 5)
	if txtLen > length {
		for i := 0; i < txtLen; i += length {
			retTxt = append(retTxt, string(txtRune[i:i+length]))
		}
	} else {
		retTxt = append(retTxt, txt)
	}
	return retTxt
}

//字符串转int
func (*Str) ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

//字符串转int64
func (*Str) ToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

//字符串转float
func (*Str) ToFloat64(s string) float64 {
	i, _ := strconv.ParseFloat(s, 64)
	return i
}

//时间戳字符串转日期字符串
func (obj *Str) ToDateStr(t string) string {
	i := obj.ToInt64(t)
	if i != 0 {
		return time.Unix(i, 0).Format(config.LayoutDate)
	}
	return ""
}

//int字符串转时间字符串
func (obj *Str) ToTimeStr(s string) string {
	i := obj.ToInt64(s)
	if i != 0 {
		return time.Unix(i, 0).Format(config.LayoutTime)
	}
	return ""
}

//日期字符串转时间对象
func (*Str) ToTime(s string) (time.Time, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(config.LayoutTime, s, loc)
}

//中文拼音排序
func (*Str) PYSort(strArr []string) []string {
	enc := simplifiedchinese.GB18030.NewEncoder()
	sort.Slice(strArr, func(i, j int) bool {
		cnamei, _ := enc.String(strArr[i])
		cnamej, _ := enc.String(strArr[j])
		return strings.Compare(cnamei, cnamej) < 0
	})
	return strArr
}

// Contains determines whether the str is in the strs.
func (*Str) Contains(str string, strs []string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}

// ReplaceIgnoreCase replace searchStr with repl in the text, case-insensitively.
func (*Str) ReplaceIgnoreCase(text, searchStr, repl string) string {
	buf := &bytes.Buffer{}
	textLower := strings.ToLower(text)
	searchStrLower := strings.ToLower(searchStr)
	searchStrLen := len(searchStr)
	var end int
	for {
		idx := strings.Index(textLower, searchStrLower)
		if 0 > idx {
			break
		}
		buf.WriteString(text[:idx])
		buf.WriteString(repl)
		end = idx + searchStrLen
		textLower = textLower[end:]
	}
	buf.WriteString(text[end:])
	return buf.String()
}

// ReplacesIgnoreCase replace searchStr-repl pairs in the text, case-insensitively.
func (*Str) ReplacesIgnoreCase(text string, searchStrRepl ...string) string {
	if 1 == len(searchStrRepl)%2 {
		return text
	}

	buf := &bytes.Buffer{}
	textLower := strings.ToLower(text)
	for i := 0; i < len(textLower); i++ {
		sub := textLower[i:]
		var found bool
		for j := 0; j < len(searchStrRepl); j += 2 {
			idx := strings.Index(sub, strings.ToLower(searchStrRepl[j]))
			if 0 != idx {
				continue
			}
			buf.WriteString(searchStrRepl[j+1])
			i += len(searchStrRepl[j]) - 1
			found = true
			break
		}
		if !found {
			buf.WriteByte(text[i])
		}
	}
	return buf.String()
}

// Enclose encloses search strings with open and close, case-insensitively.
func (*Str) EncloseIgnoreCase(text, open, close string, searchStrs ...string) string {
	buf := &bytes.Buffer{}
	textLower := strings.ToLower(text)
	for i := 0; i < len(textLower); i++ {
		sub := textLower[i:]
		var found bool
		for j := 0; j < len(searchStrs); j++ {
			idx := strings.Index(sub, strings.ToLower(searchStrs[j]))
			if 0 != idx {
				continue
			}
			buf.WriteString(open)
			buf.WriteString(text[i : i+len(searchStrs[j])])
			buf.WriteString(close)
			i += len(searchStrs[j]) - 1
			found = true
			break
		}
		if !found {
			buf.WriteByte(text[i])
		}
	}
	return buf.String()
}
