package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/kddmj"
)

func main() {
	//handCards := []int{0, 0, 31, 31, 31}
	handCards := []int{9, 9, 2, 2, 2, 3, 4, 5, 9, 10, 11, 12, 13, 14}
	laizi := []int{5}
	ok, dianshu := kddmj.CanHuWithLaiZi(handCards, laizi)
	if !ok {
		kddmj.ShowHandCards(handCards)
		fmt.Println("error")
	}
	fmt.Println("dianshu = ", dianshu)
}
