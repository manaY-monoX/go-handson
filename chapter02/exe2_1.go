package chapter02

// 練習問題2-1
import (
	"fmt"
	"strconv"
)

func Exe2_1() {
	var input string
	fmt.Print("値1(整数)を入力してください->")
	fmt.Scanln("Input: ", &input)
	// 文字列を整数に変換
	value1, _ := strconv.Atoi(input)

	fmt.Print("値2(整数)を入力してください->")
	fmt.Scanln("Input: ", &input)
	// 文字列を整数に変換
	value2, _ := strconv.Atoi(input)

	fmt.Println("値1+値2 = ", value1+value2)
	fmt.Println("値1-値2 = ", value1-value2)
	fmt.Println("値1*値2 = ", value1*value2)
	fmt.Println("値1/値2 = ", value1/value2)
	fmt.Println("値1%値2 = ", value1%value2)
}
