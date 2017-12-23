package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

func main() {
	//handCards := []int{0, 0, 31, 31, 31}
	handCards := []int{9, 9, 2, 2, 2, 3, 4, 5, 9, 10, 11, 12, 13, 14}
	laizi := []int{5}
	ok, dianshu := sxtdhmj.CanHuWithLaiZi(handCards, laizi)
	if !ok {
		sxtdhmj.ShowHandCards(handCards)
		fmt.Println("error")
	}
	fmt.Println("dianshu = ", dianshu)
}
