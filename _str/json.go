package _str

import jsoniter "github.com/json-iterator/go"

type Json struct {
}

func NewJson() *Json {
	return &Json{}
}

func (*Json) UnmarshalJSON(data []byte, v interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}

func (*Json) MarshalJSON(v interface{}) ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

func (*Json) MarshalIndentJSON(v interface{}, prefix, indent string) ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.MarshalIndent(v, prefix, indent)
}
