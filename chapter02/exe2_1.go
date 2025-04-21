package chapter02

// 練習問題2-1
import (
	"fmt"
	"strconv"
)

func Exe2_1() {
	var input string
	fmt.Print("値1(整数)を入力してください->")
	fmt.Scanln(&input)
	// 文字列を整数に変換
	value1, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("エラー: 整数を入力してください")
		return
	}

	fmt.Print("値2(整数)を入力してください->")
	fmt.Scanln(&input)
	// 文字列を整数に変換
	value2, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("エラー: 整数を入力してください")
		return
	}

	fmt.Println("値1+値2 = ", value1+value2)
	fmt.Println("値1-値2 = ", value1-value2)
	fmt.Println("値1*値2 = ", value1*value2)

	if value2 == 0 {
		fmt.Println("エラー: 値2が0のため、除算と剰余演算はできません")
		return
	}

	fmt.Println("値1/値2 = ", value1/value2)
	fmt.Println("値1%値2 = ", value1%value2)
}
