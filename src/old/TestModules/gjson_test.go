package TestModules

import (
	"github.com/tidwall/gjson"
	"testing"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func TestGjson(t *testing.T) {
	value := gjson.Get(json, "name.last")
	value1 := gjson.Get(json,"name.first")
	t.Log(value1)
	println(value.String())
}
