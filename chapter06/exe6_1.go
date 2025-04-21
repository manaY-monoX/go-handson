package chapter06

// 練習問題6-1
// シリアライズ機能を表すインターフェイスを作成する。
// 構造体にインターフェイスを実装し、YAML形式にシリアライズする。
// YAML形式への変換処理は、「補足」のメソッドを利用する。

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func Exe6_1() {
	var acc Account

	fmt.Print("口座番号: ")
	fmt.Scan(&acc.No)
	fmt.Print("名前: ")
	fmt.Scan(&acc.Name)
	fmt.Print("残高: ")
	fmt.Scan(&acc.Balance)

	// シリアライズ
	yamlData, err := acc.Serialize()
	if err != nil {
		fmt.Println("\nシリアライズに失敗しました。\n", err)
	} else {
		fmt.Println("\nシリアライズしたデータ:\n", string(yamlData))
	}
}

// シリアライズ機能を表すインターフェイス
type Serialize interface {
	Serialize() ([]byte, error) // byte[]とerrorを返す
}

// アカウント構造体
type Account struct {
	No      int    `yaml:"no"`
	Name    string `yaml:"name"`
	Balance int    `yaml:"balance"`
}

// インターフェイスの実装（レシーバ）
// YAML形式に変換した結果を返す
func (ins *Account) Serialize() ([]byte, error) {
	if yamlData, err := yaml.Marshal(ins); err != nil {
		return nil, err
	} else {
		return yamlData, nil
	}
}

// go get gopkg.in/yaml.v3
// go getでyamlをインストールする
// go.modとgo.sumにインストールしたパッケージの情報が追加される
