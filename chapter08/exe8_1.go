package chapter08

// 練習問題8-1
// 指定された値とカウンタ変数を使って、加算または乗算を５回実⾏する関数を作成する。
// 関数をゴルーチンとして実行し、sync.WaitGroup を利⽤して終了を同期化する。

import (
	"fmt"
	"sync"
	"time"
)

func syncCalc(wait *sync.WaitGroup, operation int) {
	for i := range 5 {
		if operation == 1 {
			fmt.Println("加算:", i+1, "回目の結果:", i+i)
		} else if operation == 2 {
			fmt.Println("乗算:", i+1, "回目の結果:", i*i)
		} else {
			panic("不明な操作です")
		}
		time.Sleep(1 * time.Second)
	}
	wait.Done()
}

func Exe8_1() {
	wait := new(sync.WaitGroup)
	wait.Add(2)
	go syncCalc(wait, 1)
	go syncCalc(wait, 2)
	wait.Wait()
	fmt.Println("すべての計算が完了しました")
}
