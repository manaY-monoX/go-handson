package chapter02

// // 練習問題2-3
// 入力された数値が4の倍数であるかを判定し、期待結果のようにメッセージを出力する。

import (
	"fmt"
	"strconv"
)

func Exe2_3() {
	var input string
	fmt.Print("数値を入力してください->")
	fmt.Scanln(&input)

	value, _ := strconv.Atoi(input)

	if value%4 == 0 {
		fmt.Println(input, "は4の倍数です。")
	} else {
		fmt.Println(input, "は4の倍数ではありません。")
	}
}
