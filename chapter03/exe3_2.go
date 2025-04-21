package chapter03

import "fmt"

// 練習問題3-2
// 下記の「補足」にある駅名の配列を用意し、キーボードで入力された駅名と比較する。
// 駅名が見つかった場合は、何番目の駅か表示する。
// 見つからない場合は、見つからないことを表すメッセージを表示する。

func Exe3_2() {
	var yamanoteLine = []string{"品川", "高輪ゲートウェイ", "大崎", "五反田", "目黒", "恵比寿", "渋谷", "原宿", "代々木", "新宿", "新大久保", "高田馬場", "目白", "池袋", "大塚", "巣鴨", "駒込", "田端", "西日暮里", "日暮里", "鶯谷", "上野", "御徒町", "秋葉原", "千代田", "神田", "東京", "有楽町", "新橋", "浜松町", "田町"}

	var input string
	fmt.Print("駅名を入力してください: ")
	fmt.Scanln(&input)

	for i, station := range yamanoteLine {
		if station == input {
			fmt.Println(input, "は", i+1, "番目の駅です。")
			return
		}
	}
}
