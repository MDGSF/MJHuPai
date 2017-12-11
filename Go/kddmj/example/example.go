package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/kddmj"
)

func main() {
	//handCards := []int{0, 0, 31, 31, 31}
	handCards := []int{0, 0, 1, 1, 2, 2, 3, 3, 3, 4, 4, 31, 31, 31}
	laizi := []int{31}
	ok, dianshu := kddmj.CanHuWithLaiZi(handCards, laizi)
	if !ok {
		kddmj.ShowHandCards(handCards)
		fmt.Println("error")
	}
	fmt.Println("dianshu = ", dianshu)
}
