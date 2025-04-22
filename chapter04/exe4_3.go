package chapter04

import (
	"fmt"
	"os"
)

func Exe4_3() {
	file := openFile()       // ファイルを開く
	defer closeFile(file)    // ファイルを閉じる
	fmt.Fprintf(file, "データ") // ファイルにデータを書き込む
}

// ファイルを開いてファイル構造体を返す関数
func openFile() *os.File {
	fmt.Println("ファイルを開く")
	// ファイルを開く
	file, err := os.OpenFile("./exe4_3.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file
}

// 引数指定されたファイルをクローズする
func closeFile(file *os.File) {
	if file != nil {
		fmt.Println("ファイルのクローズ")
		file.Close()
	}
}
