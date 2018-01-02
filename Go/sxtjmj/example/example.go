package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/sxtjmj"
)

func main() {
	handCards := []int{27, 27, 27, 28, 28, 28, 29, 29, 29, 31, 31}
	laizi := []int{27}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 27, true, laizi)
	if !ok {
		sxtjmj.ShowHandCards(handCards)
		fmt.Println("error")
	}
	fmt.Println("fengNum = ", fengNum)
}
