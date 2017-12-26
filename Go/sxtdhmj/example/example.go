package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

func main() {
	//handCards := []int{0, 0, 31, 31, 31}
	handCards := []int{27, 27, 28, 28, 29, 29, 30, 30, 31, 31, 31, 31, 32, 33}
	ok, fengNum := sxtdhmj.CanHu(handCards, 1, 27, true, true, false)
	if !ok {
		sxtdhmj.ShowHandCards(handCards)
		fmt.Println("error")
	}
	fmt.Println("fengNum = ", fengNum)
}
