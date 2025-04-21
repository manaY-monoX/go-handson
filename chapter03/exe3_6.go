package chapter03

// 練習問題3-6
// 社員を表す構造体を作成し、入力した値を保持した結果を表示する

import "fmt"

type Employee3_6 struct {
	EmpNo   int // 社員番号
	EmpName string
}

func Exe3_6() {
	var employee = new(Employee3_6)

	fmt.Print("社員番号を入力してください:")
	fmt.Scan(&employee.EmpNo)

	fmt.Print("社員名を入力してください:")
	fmt.Scan(&employee.EmpName)

	fmt.Println("\n社員番号:", employee.EmpNo, "\n社員名:", employee.EmpName)
}
