package jsons

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.Config{
	EscapeHTML:  true,
	SortMapKeys: true,
}.Froze()

func ToString(v any) string {
	r, _ := json.MarshalToString(v)
	return r
}

func ToStringPretty(v any) string {
	r, _ := json.MarshalIndent(v, "", "  ")
	return string(r)
}

func DeepCopy(v any, dist any) error {
	vv, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(vv, dist)
}
