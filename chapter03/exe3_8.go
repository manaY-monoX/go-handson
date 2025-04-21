package chapter03

// 練習問題3-8
// 演習3-6で作成した構造体をポインタ変数で利用する。
// また、「補足」のコードを利用して、構造体のアドレスを出力する。

import (
	"fmt"
	"reflect"
)

type Employee3_8 struct {
	EmpNo   int
	EmpName string
}

func Exe3_8() {
	var employee *Employee3_8 = new(Employee3_8)
	fmt.Println("構造体のアドレス:", reflect.ValueOf(employee).Pointer())

	// fmt.Print("社員番号を入力してください:")
	// fmt.Scan(&employee.EmpNo)

	// fmt.Print("社員名を入力してください:")
	// fmt.Scan(&employee.EmpName)

	// fmt.Println("\n社員番号:", employee.EmpNo, "\n社員名:", employee.EmpName)
}

// reflect: 実行時の型情報を利用した動的な処理を可能にします。
