package chapter02

// 練習問題2-4
// 入力された数値に対して、符号(正/0/負)を判定した結果を表示する。

import (
	"fmt"
	"strconv"
)

func Exe2_4() {
	var input string
	fmt.Print("数値を入力してください->")
	fmt.Scanln(&input)

	value, _ := strconv.Atoi(input)

	switch {
	case value > 0:
		fmt.Println("正の数です。")
	case value < 0:
		fmt.Println("負の数です。")
	default:
		fmt.Println("0です。")
	}

}
