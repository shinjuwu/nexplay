/*
浮點數運算 (float64)
1.比較大小
2.浮點數只保留兩位小數
*/
package utils

import (
	"math"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

// 精準度設定
func Accuracy() float64 {
	return 0.001
}

// a == b
func Equal(a, b float64) bool {
	return math.Abs(a-b) < Accuracy()
}

// a > b
func Greater(a, b float64) bool {
	return math.Max(a, b) == a && math.Abs(a-b) > Accuracy()
}

// a < b
func Smaller(a, b float64) bool {
	return math.Max(a, b) == b && math.Abs(a-b) > Accuracy()
}

// a >= b
func GreaterOrEqual(a, b float64) bool {
	return math.Max(a, b) == a || math.Abs(a-b) < Accuracy()
}

// a <= b
func SmallerOrEqual(a, b float64) bool {
	return math.Max(a, b) == b || math.Abs(a-b) < Accuracy()
}

// 只保留浮點數後兩位小數(不進位)
func Decimal2(value float64) float64 {
	value, _ = strconv.ParseFloat(ChangeNumber(value, 2), 64)
	// value, _ = decimal.NewFromFloatWithExponent(value, -2).Float64()
	return value
}

// 只保留浮點數後三位小數(不進位)
func Decimal3(value float64) float64 {
	value, _ = strconv.ParseFloat(ChangeNumber(value, 3), 64)
	// value, _ = decimal.NewFromFloatWithExponent(value, -3).Float64()
	return value
}

// 只保留浮點數後四位小數(不進位)
func Decimal4(value float64) float64 {
	value, _ = strconv.ParseFloat(ChangeNumber(value, 4), 64)
	// value, _ = decimal.NewFromFloatWithExponent(value, -4).Float64()
	return value
}

// float64 轉 string
// f: 欲轉換的數字
// m: 保留小數後幾位
func ChangeNumber(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

// n1 + n2 (回傳值只到小數點第四位)
func DecimalAdd(n1 float64, n2 float64) (Value float64) {
	g1 := decimal.NewFromFloat(Decimal4(n1))
	g2 := decimal.NewFromFloat(Decimal4(n2))
	Value, _ = g1.Add(g2).Float64()
	return Value
}

// n1 - n2 (回傳值只到小數點第四位)
func DecimalSub(n1 float64, n2 float64) (Value float64) {
	g1 := decimal.NewFromFloat(Decimal4(n1))
	g2 := decimal.NewFromFloat(Decimal4(n2))
	Value, _ = g1.Sub(g2).Float64()
	return Value
}

// n1 * n2 (回傳值只到小數點第四位)
func DecimalMul(n1 float64, n2 float64) (Value float64) {
	g1 := decimal.NewFromFloat(Decimal4(n1))
	g2 := decimal.NewFromFloat(Decimal4(n2))
	Value, _ = g1.Mul(g2).Float64()
	return Value
}

// n1 / n2 (回傳值只到小數點第四位)
func DecimalDiv(n1 float64, n2 float64) (Value float64) {
	g1 := decimal.NewFromFloat(Decimal4(n1))
	g2 := decimal.NewFromFloat(Decimal4(n2))
	Value, _ = g1.Div(g2).Float64()
	return Value
}
