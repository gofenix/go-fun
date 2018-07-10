package main

import "fmt"

func main() {
	m := map[string]interface{}{
		"name": "zhuzhenfeng",
		"age":  "18",
	}

	m2 := make(map[string]interface{})

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m2)
}
