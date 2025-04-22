package chapter08

// 練習問題8-2
// 練習問題8-1で作成したプログラムをゴルーチンを使って並列処理を実装する。
// ゴルーチンを使って並列処理を実装する。

import (
	"fmt"
	"time"
)

func selectCalc(operation int, channel chan<- string, endChannel chan<- int) {
	// for i := range 5 の代わりに、通常のforループを使用
	for i := 0; i < 5; i++ {
		if operation == 1 {
			channel <- fmt.Sprintf("%d回目の加算値: %d", i+1, i+i)
		} else if operation == 2 {
			channel <- fmt.Sprintf("%d回目の乗算値: %d", i+1, i*i)
		} else {
			panic("不正なoperation")
		}
		time.Sleep(1 * time.Second)
	}
	endChannel <- 1
}

func Exe8_2() {
	channel := make(chan string) // 計算結果の受信
	endChannel := make(chan int) // 終了の受信

	close := func() {
		close(channel)
		close(endChannel)
	}
	defer close()

	go selectCalc(1, channel, endChannel)
	go selectCalc(2, channel, endChannel)

	endCount := 0
	for {
		select {
		case getChannel := <-channel:
			fmt.Println(getChannel)
		case getEndChannel := <-endChannel:
			endCount += getEndChannel
			// チャネルから値を受信するごとにデバッグ情報を出力
			fmt.Printf("endCount: %d\n", endCount)
		default:
			// default句があると、select文はブロックせずに即座に実行されてしまいます
			// チャネルから値を受信できない場合にpanicが発生してしまうため、default句を削除
		}
		if endCount == 2 {
			break
		}
	}
	fmt.Println("プログラムを正常終了します。")
}
