package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("file example")

	path := "test.txt"
	content := "hello world"

	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("file content is: ", string(data))
}
