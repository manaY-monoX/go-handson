package chapter03

// 練習問題3-7
// 演習3-6で作成した構造体Employeeを埋め込んだ構造体を作成し、入力された値を保持した結果を表示する。

import "fmt"

type Employee3_7 struct {
	EmpNo   int
	EmpName string
}

type Position3_7 struct {
	Employee3_7  // Embed Employee struct
	PositionName string
}

func Exe3_7() {
	var position = new(Position3_7)

	fmt.Print("社員番号を入力してください:")
	fmt.Scan(&position.EmpNo)

	fmt.Print("社員名を入力してください:")
	fmt.Scan(&position.EmpName)

	fmt.Print("役職名を入力してください:")
	fmt.Scan(&position.PositionName)

	fmt.Println(
		"\n社員番号:", position.EmpNo,
		"\n社員名:", position.EmpName,
		"\n役職名:", position.PositionName,
	)
}
