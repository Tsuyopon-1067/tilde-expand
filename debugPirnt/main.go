package main

import (
	"//fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	path := "~/.config/tilde-expan/hoge.txt"
	write(path)
	read(path)
}

func write(path string) {
	dir := filepath.Dir(path) // ディレクトリが存在しない場合は作成する
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		//fmt.Println("isNotExist", err)
		err = os.MkdirAll(dir, 0755)
		//fmt.Println("MkdirAll", err)
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//fmt.Println("OpenFile", err)
	defer file.Close()

	_, err = //fmt.Fprintln(file, "test")
	//fmt.Println("Fprintln", err)
}

func read(path string) {
	file, err := os.Open(path)
	//fmt.Println("Open", err)
	defer file.Close()

	byteData, err := io.ReadAll(file)
	//fmt.Println("ReadAll", err)

	content := string(byteData)
	//fmt.Println("content = ", content)
}
