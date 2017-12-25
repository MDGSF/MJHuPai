package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

func main() {
	//handCards := []int{0, 0, 31, 31, 31}
	handCards := []int{9, 9, 2, 2, 2, 33, 33, 33, 32, 32, 32, 31, 31, 31}
	ok, fengNum := sxtdhmj.CanHu(handCards, 2, 33, true, true, false)
	if !ok {
		sxtdhmj.ShowHandCards(handCards)
		fmt.Println("error")
	}
	fmt.Println("fengNum = ", fengNum)
}
