package main

import (
	"fmt"
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
		err = os.MkdirAll(dir, 0755)
	}

	file, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	fmt.Fprintln(file, "test")
}

func read(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	byteData, _ := io.ReadAll(file)
	content := string(byteData)
	fmt.Println("content = ", content)
}
