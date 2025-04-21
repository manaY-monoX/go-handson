package chapter03

// 練習問題3-1
// 入力された値が数値変換できるかを判定し、結果を表示する。
import (
	"fmt"
	"strconv"
)

func Exe3_1() {
	var input string
	fmt.Print("値を入力: ")
	fmt.Scanln(&input)

	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error! 数値を入力してください。")
		return
	}

	fmt.Println("入力された値は", value, "です。")
}
