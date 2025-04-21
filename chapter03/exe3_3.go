package chapter03

import "fmt"

// 練習問題3-3
// 「補足」の月名の配列からスライスを作成し、append()関数を使って"October"を追加したら内容を表示する。

func Exe3_3() {
	var months = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September"}
	months = append(months, "October")

	for i, month := range months {
		fmt.Println(i+1, month)
	}
}
