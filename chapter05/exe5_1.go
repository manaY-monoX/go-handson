package chapter05

// 練習問題5-1
// 構造体のフィールド値を出力するメソッドを作成し、実行結果を確認する。

import "fmt"

func Exe5_1() {
	acc := new(Account) // 口座情報を管理する構造体のポインタを作成
	acc.no = "10710"
	acc.name = "Yamashita"
	acc.balance = 4000000
	acc.Print()
}

// 口座情報を管理する構造体
type Account struct {
	no      string // 口座番号
	name    string // 口座名義
	balance int    // 残高
}

// 口座情報を出力するメソッド（レシーバ）
func (acc Account) Print() {
	fmt.Println("口座番号:", acc.no)
	fmt.Println("口座名義:", acc.name)
	fmt.Println("残高:", acc.balance)
}
