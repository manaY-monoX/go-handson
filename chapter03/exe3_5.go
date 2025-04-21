package chapter03

// 練習問題3-5
// 「補足」の月名のマップからスライスを作成し、入力された月名が存在すれば値を表示する。
// 存在しなければ「存在しません」を表示する。

import "fmt"

func Exe3_5() {
	months_map := map[string]int{"April": 30, "May": 31, "June": 30, "July": 31, "August": 31, "September": 30}

	fmt.Print("月名を入力してください:")
	var input string
	fmt.Scan(&input)

	for month, days := range months_map {
		if month == input {
			fmt.Println(month, "は", days, "日です。")
			break
		} else {
			fmt.Println("存在しません")
		}
	}
}
