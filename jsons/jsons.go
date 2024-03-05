package jsons

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/yhh-show/helpers/logger"
)

var json = jsoniter.Config{
	EscapeHTML:  true,
	SortMapKeys: true,
}.Froze()

func ToString(v any) string {
	r, e := json.MarshalToString(v)
	if e != nil {
		logger.L.Println("jsons.ToString error:", e, v)
	}
	return r
}

func ToStringPretty(v any) string {
	r, e := json.MarshalIndent(v, "", "  ")
	if e != nil {
		logger.L.Println("jsons.ToStringPretty error:", e, v)
	}
	return string(r)
}

func DeepCopy(v any, dist any) error {
	vv, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(vv, dist)
}
