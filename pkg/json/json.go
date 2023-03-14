package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func JGet(json, tag string) gjson.Result {
	return gjson.Get(json, tag)
}

// PPrint beautiful print struct
func PPrint(prefix string, v interface{}) {
	vs, _ := json.Marshal(v)
	var out bytes.Buffer
	_ = json.Indent(&out, vs, "", "\t")
	fmt.Printf("%s:%v\n", prefix, out.String())
}
