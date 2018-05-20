package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// writeFileDemo()
	// writeAt()
	bufferWriterDemo()
}

func writeFileDemo() {
	fmt.Println("ioutil writeFile example")

	path := "test.txt"
	content := "hello world"

	err := ioutil.WriteFile(path, []byte(content), 0755)
	errCheck(err)

	data, err := ioutil.ReadFile(path)
	errCheck(err)

	fmt.Println("file content is: ", string(data))
}

func writeAt() {
	fmt.Println("write file at")

	path := "test2.txt"
	content := "hello"
	newContent := "world"

	newFile, err := os.Create(path)
	errCheck(err)

	n, err := newFile.WriteAt([]byte(content), 0)
	errCheck(err)

	m, err := newFile.WriteAt([]byte(newContent), 10)
	errCheck(err)

	data, err := ioutil.ReadFile(path)
	errCheck(err)

	fmt.Println(string(data))

	fmt.Println(n, m)

}

func bufferWriterDemo() {
	fmt.Println("buffer writer file at")

	path := "test2.txt"
	content := "hello"

	newFile, err := os.Create(path)
	errCheck(err)

	bufferWriter := bufio.NewWriter(newFile)
	for _, v := range content {
		n, err := bufferWriter.WriteString(string(v))
		errCheck(err)
		fmt.Println(n)
	}
	bufferWriter.Flush()

	data, err := ioutil.ReadFile(path)
	errCheck(err)

	fmt.Println(string(data))
}
