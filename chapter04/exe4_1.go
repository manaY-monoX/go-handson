package chapter04

// 練習問題4-1
// 入力された値で四則演算した結果を返す関数を作成し、計算結果を出力する。

import (
	"fmt"
	"strconv"
	"strings"
)

func Exe4_1() {
	var input1, input2 string
	var num1, num2 int
	var err error

	fmt.Print("整数1を入力してください:")
	fmt.Scan(&input1)
	num1, err = strconv.Atoi(strings.TrimSpace(input1))
	if err != nil {
		fmt.Println("エラー: 整数を入力してください。")
		return
	}

	fmt.Print("整数2を入力してください:")
	fmt.Scan(&input2)
	num2, err = strconv.Atoi(strings.TrimSpace(input2))
	if err != nil {
		fmt.Println("エラー: 整数を入力してください。")
		return
	}

	if num2 == 0 {
		fmt.Println("エラー: 0で割ることはできません。")
		return
	}

	sum, diff, prod, quot := calc4_namedReturn(num1, num2)

	fmt.Println("和:", sum)
	fmt.Println("差:", diff)
	fmt.Println("積:", prod)
	fmt.Println("商:", quot)

	println("--- 名前付き戻り値を省略した場合 ---")
	sum, diff, prod, quot = calc4_unnamedReturn(num1, num2)

	fmt.Println("和:", sum)
	fmt.Println("差:", diff)
	fmt.Println("積:", prod)
	fmt.Println("商:", quot)
}

// 四則演算を行う関数（名前付き戻り値）
func calc4_namedReturn(num1 int, num2 int) (sum int, diff int, prod int, quot int) {
	sum = num1 + num2
	diff = num1 - num2
	prod = num1 * num2
	quot = num1 / num2

	return
}

// 四則演算を行う関数（戻り値を省略）
func calc4_unnamedReturn(num1 int, num2 int) (int, int, int, int) {
	sum := num1 + num2
	diff := num1 - num2
	prod := num1 * num2
	quot := num1 / num2

	return sum, diff, prod, quot
}
