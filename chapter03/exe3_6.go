package chapter03

// 練習問題3-6
// 社員を表す構造体を作成し、入力した値を保持した結果を表示する

import "fmt"

type Employee struct {
	EmpNo   int // 社員番号
	EmpName string
}

func Exe3_6() {
	fmt.Print("社員番号を入力してください:")
	var empNo int
	fmt.Scan(&empNo)

	fmt.Print("社員名を入力してください:")
	var empName string
	fmt.Scan(&empName)

	employee := Employee{EmpNo: empNo, EmpName: empName}
	fmt.Println("\n社員番号:", employee.EmpNo, "\n社員名:", employee.EmpName)
}
