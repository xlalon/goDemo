package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func JParse(json string) gjson.Result {
	return gjson.Parse(json)
}

func JGet(json, tag string) gjson.Result {
	return gjson.Get(json, tag)
}

// JPrint beautiful print struct
func JPrint(prefix string, v interface{}) {
	vs, _ := json.Marshal(v)
	var out bytes.Buffer
	_ = json.Indent(&out, vs, "", "\t")
	fmt.Printf("%s:%v\n", prefix, out.String())
}
