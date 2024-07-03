package main

import (
	"fmt"
	"math/big"

	"github.com/cockroachdb/apd"
	"github.com/shopspring/decimal"
)


func MathBigBasicAddition(a, b string) {
	x := new(big.Float).SetPrec(200)
	x.SetString(a)

	y := new(big.Float).SetPrec(200)
	y.SetString(b)

	resultAdd := new(big.Float).SetPrec(200)
	resultAdd.Add(x, y)
	fmt.Println("Addition:", resultAdd)
}

func MathBigBasicSubtraction(a, b string) {
	x := new(big.Float).SetPrec(200)
	x.SetString(a)

	y := new(big.Float).SetPrec(200)
	y.SetString(b)

	// Subtraction
	resultSub := new(big.Float).SetPrec(200)
	resultSub.Sub(x, y)
	fmt.Println("Subtraction:", resultSub)
}

func MathBigBasicMultiplication(a, b string) {
	// EXAMPLE 3 - MATHEMATICAL OPERATIONS
	x := new(big.Float).SetPrec(200)
	x.SetString(a)

	y := new(big.Float).SetPrec(200)
	y.SetString(b)

	// Multiplication
	resultMul := new(big.Float).SetPrec(200)
	resultMul.Mul(x, y)
	fmt.Println("Multiplication:", resultMul)
}

func MathBigDivision(a, b string) {
	x := new(big.Float).SetPrec(200)
	x.SetString(a)

	y := new(big.Float).SetPrec(200)
	y.SetString(b)

	// Division x/y
	resultDiv := new(big.Float).SetPrec(200)
	resultDiv.Quo(x, y)
	fmt.Println("Division:", resultDiv)
}

func DecimalAdd(a, b string) {
	btcAmount1, _ := decimal.NewFromString(a)
  btcAmount2, _ := decimal.NewFromString(b)

	resultAddDecimal := btcAmount1.Add(btcAmount2)

	fmt.Println("result ", resultAddDecimal)
}

func DecimalSub(a, b string) {
	btcAmount1, _ := decimal.NewFromString(a)
  btcAmount2, _ := decimal.NewFromString(b)

	resultAddDecimal := btcAmount1.Sub(btcAmount2)

	fmt.Println("result ", resultAddDecimal)
}

func DecimalMul(a, b string) {
	btcAmount1, _ := decimal.NewFromString(a)
  btcAmount2, _ := decimal.NewFromString(b)

	resultAddDecimal := btcAmount1.Mul(btcAmount2)

	fmt.Println("result ", resultAddDecimal)
}

func DecimalDiv(a, b string) {
	btcAmount1, _ := decimal.NewFromString(a)
  btcAmount2, _ := decimal.NewFromString(b)

	resultAddDecimal := btcAmount1.Div(btcAmount2)

	fmt.Println("result ", resultAddDecimal)
}

func CockroachAdd(a, b string) {
	btcAmount1 := new(apd.Decimal)
	_, _, _ = btcAmount1.SetString(a)

	btcAmount2 := new(apd.Decimal)
    _, _, _ = btcAmount2.SetString(b)

	resultAdd := new(apd.Decimal)
  _, _ = apd.BaseContext.Add(resultAdd, btcAmount1, btcAmount2)

	fmt.Println("final result", resultAdd)
}

func CockroachMul(a, b string) {
	btcAmount1 := new(apd.Decimal)
	_, _, _ = btcAmount1.SetString(a)

	btcAmount2 := new(apd.Decimal)
    _, _, _ = btcAmount2.SetString(b)

	resultMul := new(apd.Decimal)
   _, _ = apd.BaseContext.Mul(resultMul, btcAmount1, btcAmount2)

	fmt.Println("final value ", resultMul)
}

func main() {
	fmt.Println("hello")
}
