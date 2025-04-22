package chapter12

// 演習12-3
// 演習12-files.txtのデータを出力する。
// Reader構造体のReadLine()メソッドを利用する。

import (
	"fmt"
	"io"
	"os"
)

func ReadFile() {

	// ファイルを開く
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("ファイルを開けませんでした:", err)
		return
	}
	defer func() {
		if file != nil {
			file.Close() // ファイルが開かれていれば閉じる
		}
	}()

	// ファイルの先頭に移動
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Seekに失敗しました:", err)
		return
	}

	// ファイルの内容を読み取る
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("読み込みに失敗しました:", err)
		return
	}
	fmt.Println("ファイルの読み込みに成功しました。\nfile.txtの内容は以下")
	fmt.Println("--------------------------------")
	fmt.Println(string(buf[:n]))
	fmt.Println("--------------------------------\n以上")
}

func Exe12_3() {
	ReadFile()
}
