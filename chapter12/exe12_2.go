package chapter12

// 演習12-2
// 入力された値をファイルに書き込み続け、End が入力されたらクローズして終了する。
// Writer構造体のWriteString()、Flash()メソッドを利用する。
// 「補足」のコードを利用してファイルを作成する。

import (
	"bufio"
	"fmt"
	"os"
)

func UseFileStructure() {
	file, err := os.Create("file.txt")
	defer func() {
		if file != nil {
			file.Close() // ファイルが開かれていれば閉じる
		}
	}()
	if err != nil { // ファイルが作成できなかったらerrを表示して終了
		fmt.Println(err)
		return
	}
	// ファイルのデータを書き込む
	writer := bufio.NewWriter(file)
	fmt.Println("ファイルに書き込むデータを入力してください")
	for {
		var input string
		fmt.Scan(&input)
		writer.WriteString(input + "\n")
		if input == "end" || input == "END" {
			break
		}
	}
	writer.Flush() // bufio.Writerのバッファをファイルに書き込む
}

func Exe12_2() {
	UseFileStructure()
	fmt.Println("ファイルに書き込みました。")

}
