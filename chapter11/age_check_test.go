package chapter11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 演習11-1
// go test

func Exe11_1(t *testing.T) {
	TestCase01(t)
	TestCase02(t)
	TestCase03(t)
}

func TestCase01(t *testing.T) {
	result, err := AgeCheck(0) // 0歳の場合
	assert.Equal(t, 1, result) // 結果が1であることを確認
	assert.Nil(t, err)         // エラーがnilであることを確認
}

func TestCase02(t *testing.T) {
	result, err := AgeCheck(20) // 20歳の場合
	assert.Equal(t, 2, result)  // 結果が2であることを確認
	assert.Nil(t, err)          // エラーがnilであることを確認
}

func TestCase03(t *testing.T) {
	result, err := AgeCheck(66) // 65歳の場合
	assert.Equal(t, 3, result)  // 結果が3であることを確認
	assert.Nil(t, err)          // エラーがnilであることを確認
}

func TestCase04(t *testing.T) {
	result, err := AgeCheck(1000)           // 1000歳の場合
	assert.Equal(t, -1, result)             // 結果が-1であることを確認
	assert.Equal(t, "不正な年齢です", err.Error()) // エラーが"不正な年齢です"であることを確認
}
