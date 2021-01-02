package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func walkRecursive(dict interface{}, isP2 bool) int {
	ans := 0

	if num, ok := dict.(float64); ok {
		ans += int(num)
	} else if arr, ok := dict.([]interface{}); ok {
		for _, v := range arr {
			ans += walkRecursive(v, isP2)
		}
	} else if subDict, ok := dict.(map[string]interface{}); ok {
		for _, v := range subDict {
			if str, ok := v.(string); isP2 && ok && str == "red" {
				return 0
			}
			ans += walkRecursive(v, isP2)
		}
	}

	return ans
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	var dict interface{}
	json.Unmarshal(input, &dict)

	ansP1 := walkRecursive(dict, false)
	ansP2 := walkRecursive(dict, true)

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
