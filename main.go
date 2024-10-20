package main

import (
	"fmt"
	jsonserialize "json-serialize/json_serialize"
)

type Name struct {
	Dyna     map[string]any
	LastName string
}

func main() {
	fmt.Println(jsonserialize.Serialize(nil))
	fmt.Println(jsonserialize.Serialize("woow"))
	fmt.Println(jsonserialize.Serialize(123))
	fmt.Println(jsonserialize.Serialize(int64(32)))
	fmt.Println(jsonserialize.Serialize(uint(32)))
	fmt.Println(jsonserialize.Serialize(10e5))
	fmt.Println(jsonserialize.Serialize([]int{1, 2, 3, 4}))
	fmt.Println(jsonserialize.Serialize([]string{"1", "123"}))
	fmt.Println(jsonserialize.Serialize([][]int{{1, 2, 3}, {12, 3, 4}}))
	fmt.Println(jsonserialize.Serialize([]any{1, "woow"}))
	fmt.Println(jsonserialize.Serialize("WooW\"WooW"))
	fmt.Println(jsonserialize.Serialize(map[string]any{"1": 1, "2": "string", "3": []int{1, 2, 3}, "bool": true, "dynamic": []any{"1", 2, 3, true, nil, false}}))

	n := Name{
		Dyna:     map[string]any{"1": 1, "2": "string", "3": []int{1, 2, 3}, "bool": true, "dynamic": []any{"1", 2, 3, true, nil, false}},
		LastName: "Not \"WooW",
	}

	v := jsonserialize.Serialize(n)
	fmt.Println(v)
}
