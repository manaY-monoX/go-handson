package chapter07

import (
	"fmt"
	"strconv"
)

func Sum7_2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 異常終了通知を出す関数
func result() {
	fmt.Println("プログラムが異常終了しました")
}

func Exe7_2() {
	// パニックをリカバリするための関数をdefer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nエラーが発生しました:", r)
			result() // 異常終了通知
			return
		}
	}()

	var inputs []int
	for {
		var input string
		fmt.Print("整数を入力(0で終了): ")
		fmt.Scan(&input)

		// 数値変換とエラーチェック
		value, err := strconv.Atoi(input)
		if err != nil {
			panic("整数以外が入力されました") // ここでパニックを発生させる
		}

		if value == 0 {
			break // 0が入力された場合はループを抜ける
		}

		inputs = append(inputs, value)
	}

	fmt.Println("\n合計:", Sum7_2(inputs...))
	fmt.Println("プログラムが正常終了しました")
}
