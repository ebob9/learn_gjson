package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"strings"
)

var json = `{
  "name": {"first": "Tom", "last": "Anderson", "dead": null, "testfalse": false, "testtrue": true},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

func isArray(chunk gjson.Result) bool {
	_, chunkResult := chunk.Value().([]interface{})
	return chunkResult
}

func isMap(chunk gjson.Result) bool {
	_, chunkResult := chunk.Value().(map[string]interface{})
	return chunkResult
}

func testParse(chunk gjson.Result, indent int) {
	indent_str := strings.Repeat(" ", indent)
	print(indent_str, "** Start for : ", chunk.String(), "\n")
	print(indent_str, "   Type(int) : ", chunk.Type, "\n")
	fmt.Printf("%s   Type(fmt) : %T\n", indent_str, chunk.Value())
	print(indent_str, "   Type(str) : ", chunk.Type.String(), "\n")
	if chunk.Type.String() == "JSON" {
		print(indent_str, "     IsArray : ", isArray(chunk), "\n")
		print(indent_str, "     IsMap   : ", isMap(chunk), "\n")
	}
	print("\n")
}

func main() {
	// Read playground
	testParse(gjson.Get(json, "friends"), 0)
	testParse(gjson.Get(json, "age"), 0)
	testParse(gjson.Get(json, "children"), 0)
	value := gjson.Get(json, "name")
	testParse(value, 0)
	subvalue := gjson.Get(value.Raw, "first")
	testParse(subvalue, 0)

	// modify playground the hard way - not using query but rebuild
	println("")
	newSubValue, err := sjson.Set(value.Raw, "first", "Harry")
	println("** Set Result : ", newSubValue, " Error : ", err)
	getNewSubValue := gjson.Parse(newSubValue)
	testParse(getNewSubValue, 2)
	println("** End Set Result \n")
	json, err = sjson.Set(json, "name", getNewSubValue.Value())
	println("** Set Result : ", json, " Error : ", err)
	testParse(gjson.Parse(json), 2)
	println("** End Set Result \n")

	// reparse
	value = gjson.Get(json, "name")
	testParse(value, 0)
	subvalue = gjson.Get(value.Raw, "first")
	testParse(subvalue, 0)
}
