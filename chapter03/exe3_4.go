package chapter03

import "fmt"

// 練習問題3-4
// 「補足」の月名のマップを作成し、すべてのキーと値を表示する。

func Exe3_4() {
	months_map := map[string]int{"April": 30, "May": 31, "June": 30, "July": 31, "August": 31, "September": 30}

	for key, value := range months_map {
		fmt.Println("key=", key, "value=", value)
	}
}
