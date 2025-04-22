package chapter10

// 練習問題10-1
// プログラムが配置されたパッケージにYAMLファイルを作成する。
// GOのファイル埋め込み機能を利用してYAMLファイルを埋め込んでビルド結果を実装する。

import (
	"embed"
	"fmt"
	"log"
	"os"
)

type Gretings struct {
	Morning string `yaml:"morning"`
	Daytime string `yaml:"daytime"`
	Evening string `yaml:"evening"`
}

//go:embed greeting.yml
var greetingsYamlFile embed.FS

func UseEmbed() {
	// ファイルの内容を読み込む
	data, err := greetingsYamlFile.ReadFile("greeting.yml")
	if err != nil {
		log.Fatal(err)
	}
	// ファイルの内容を表示する
	os.Stdout.Write(data)
}

func Exe10_1() {
	UseEmbed()
	fmt.Println("\nファイルの内容を表示しました")
}
