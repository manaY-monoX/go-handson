package chapter07

import (
	"fmt"
	"strconv"
)

func Sum7_1(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Exe7_1() {
	var inputs []int
	for {
		var input string
		fmt.Print("整数を入力(0で終了): ")
		fmt.Scan(&input)

		// 先にエラーチェックを行う
		value, err := strconv.Atoi(input)
		if err != nil {
			panic("整数以外が入力されました") // 整数以外が入力された場合はパニック
		}

		if value == 0 {
			break // 0が入力された場合はループを抜ける
		}

		inputs = append(inputs, value)
	}
	fmt.Println("合計:", Sum7_1(inputs...))
}
