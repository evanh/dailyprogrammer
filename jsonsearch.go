package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func searchVal(path string, interVal interface{}) string {
	switch val := interVal.(type) {
	case string:
		if val == "dailyprogrammer" {
			return path
		}
	case []interface{}:
		if found := SearchList(val); found != "" {
			return path + " -> " + found
		}
	case map[string]interface{}:
		if found := SearchObject(val); found != "" {
			return path + " -> " + found
		}
	default:
		return ""
	}
	return ""
}

func SearchList(ts []interface{}) string {
	for i := range ts {
		if val := searchVal(strconv.Itoa(i), ts[i]); val != "" {
			return val
		}
	}
	return ""
}

func SearchObject(obj map[string]interface{}) string {
	for k := range obj {
		if val := searchVal(k, obj[k]); val != "" {
			return val
		}
	}
	return ""
}

func main() {
	data, err := ioutil.ReadFile("challenge2.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	obj := make(map[string]interface{})

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("Error decoding json:", err)
		return
	}

	path := SearchObject(obj)
	if path != "" {
		fmt.Println(path)
	}
}
