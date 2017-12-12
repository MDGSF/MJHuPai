package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/kddmj"
)

func main() {
	//handCards := []int{0, 0, 31, 31, 31}
	handCards := []int{0, 0, 9, 9, 9, 0, 1, 2, 3, 4, 5, 10, 11, 12}
	laizi := []int{12}
	ok, dianshu := kddmj.CanHuWithLaiZi(handCards, laizi)
	if !ok {
		kddmj.ShowHandCards(handCards)
		fmt.Println("error")
	}
	fmt.Println("dianshu = ", dianshu)
}
