package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/sxtjmj"
)

func main() {
	handCards := []int{1, 2, 3, 27, 27, 28, 28, 29, 29, 30, 30}
	laizi := []int{1, 2}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 27, true, laizi)
	if !ok {
		sxtjmj.ShowHandCards(handCards)
		fmt.Println("error")
	}
	fmt.Println("fengNum = ", fengNum)
}
