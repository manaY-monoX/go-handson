package chapter11

// エラー構造体
type ValueError struct {
	message string
}

// エラー出力メソッド
func (value ValueError) Error() string {
	return value.message
}

// 演習11-1 testifyフレームワークを利用する
// テストターゲット関数
func AgeCheck(age int) (result int, err error) {
	switch {
	case age < 0 || age > 120: // 0歳未満、121歳以上
		return -1, ValueError{"不正な年齢です"}
	case age < 20: // 20歳未満
		result = 1
	case age >= 20 && age <= 65: // 20歳以上65歳以下
		result = 2
	default: // 66歳以上120歳以下
		result = 3
	}
	return
}
