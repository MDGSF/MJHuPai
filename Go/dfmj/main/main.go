package main

import (
	"fmt"
	"strconv"

	"github.com/MDGSF/MJHuPai/Go/dfmj"
)

func main() {
	slots := []int{
		2, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 3, 3, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}
	ret := dfmj.CalcKey(slots, nil)
	fmt.Println("ret = ", ret, strconv.FormatInt(int64(ret), 16), dfmj.IsWinable(slots))
}
