package main

import (
	"fmt"
	"math/big"
)

// SetPrec(200) - Here 200 is the number of precision in bits = 60 decimal points

func main() {
	// EXAMPLE 1
	usdtAmountStr := "4"
	usdtToBtcRateStr := "0.00034"

	usdtAmount := new(big.Float).SetPrec(200)
	usdtAmount.SetString(usdtAmountStr)

	usdtToBtcRate := new(big.Float).SetPrec(200)
	usdtToBtcRate.SetString(usdtToBtcRateStr)

	// BTC value = USDT amount * USDT to BTC rate
	btcValue := new(big.Float).SetPrec(200)
	btcValue.Mul(usdtAmount, usdtToBtcRate)

	fmt.Printf("%.18f USDT is equivalent to %.18f BTC\n", usdtAmount, btcValue)
	// ################################## //


	// EXAMPLE 2
	originalValueStr := "1000.123456789"
	percentageStr := "12.5"

	originalValue := new(big.Float).SetPrec(200)
	originalValue.SetString(originalValueStr)

	percentage := new(big.Float).SetPrec(200)
	percentage.SetString(percentageStr)

	// percentage value = (original value * percentage) / 100
	percentageValue := new(big.Float).SetPrec(200)
	percentageValue.Mul(originalValue, percentage)
	percentageValue.Quo(percentageValue, big.NewFloat(100))
	// ################################## //


	// EXAMPLE 3 - MATHEMATICAL OPERATIONS
	x := new(big.Float).SetPrec(200)
	x.SetString("123456789.123456789123456789")

	y := new(big.Float).SetPrec(200)
	y.SetString("987654321.987654321987654321")

	// Addition
	resultAdd := new(big.Float).SetPrec(200)
	resultAdd.Add(x, y)
	fmt.Println("Addition:", resultAdd)

	// Subtraction
	resultSub := new(big.Float).SetPrec(200)
	resultSub.Sub(x, y)
	fmt.Println("Subtraction:", resultSub)

	// Multiplication
	resultMul := new(big.Float).SetPrec(200)
	resultMul.Mul(x, y)
	fmt.Println("Multiplication:", resultMul)

	// Division
	resultDiv := new(big.Float).SetPrec(200)
	resultDiv.Quo(x, y)
	fmt.Println("Division:", resultDiv)
	// ################################## //
}
