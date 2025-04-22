package chapter12

// 演習12-1
// 入力された年、月、日から曜日を求めて表示する。
// 日本語の曜日はMapから取得する。

import (
	"fmt"
	"time"
)

var dayOfWeek = map[time.Weekday]string{
	time.Monday:    "月",
	time.Tuesday:   "火",
	time.Wednesday: "水",
	time.Thursday:  "木",
	time.Friday:    "金",
	time.Saturday:  "土",
	time.Sunday:    "日",
}

// ⽇付や時間の⼀部を取得する
// author: Fullness, Inc.
// version: 1.0
// func DateTimeFormat() {
// 	t := time.Now()
// 	fmt.Printf("%d年%02d⽉%02d⽇%02d時:%02d分:%02d秒 曜日:%s\n",
// 		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(),
// 		t.Weekday().String())
// }

// func DateTimeNow() {
// 	t := time.Now()
// 	fmt.Println(t)
// }

func DisplayDayOfWeek(year int, month int, day int) {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	fmt.Println(t)
	fmt.Println("入力された日付は[", dayOfWeek[t.Weekday()], "]曜日です。")
}

func Exe12_1() {
	// DateTimeFormat()
	// DateTimeNow()
	fmt.Println("年月日を入力してください")
	var year, month, day int
	fmt.Scan(&year, &month, &day)
	DisplayDayOfWeek(year, month, day)
}
