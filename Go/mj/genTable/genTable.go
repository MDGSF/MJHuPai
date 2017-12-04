package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/mj"
)

const LaiZiNum = 9

//只保存最终正确的结果。
var tableMgr *mj.TableMgr

//除了保存最终正确的结果，还要把中间计算过的错误结果也保存起来，防止重复计算。
var tableXuShuTemp [LaiZiNum]*map[int]bool
var tableXuShuWithEyeTemp = [LaiZiNum]*map[int]bool{}
var tableZiTemp = [LaiZiNum]*map[int]bool{}
var tableZiWithEyeTemp = [LaiZiNum]*map[int]bool{}

var curTable *[LaiZiNum]*map[int]bool
var curTableTemp *[LaiZiNum]*map[int]bool
var curCardsTypeNum int

func main() {

	fmt.Println("main start")

	tableMgr = mj.NewTableMgr()

	for i := 0; i < LaiZiNum; i++ {
		tableXuShuTemp[i] = &map[int]bool{}
		tableXuShuWithEyeTemp[i] = &map[int]bool{}
		tableZiTemp[i] = &map[int]bool{}
		tableZiWithEyeTemp[i] = &map[int]bool{}
	}

	genTableXuShu()
	genTableXuShuWithEye()
	genTableZi()
	genTableZiWithEye()

	tableMgr.Dump()
}

func genTableXuShu() {
	fmt.Println("genTableXuShu start")

	curTable = &tableMgr.TableXuShu.Map
	curTableTemp = &tableXuShuTemp
	curCardsTypeNum = 9
	cards := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	genXuShuPuZi(cards, 1)

	fmt.Println("genTableXuShu success")
}

func genTableXuShuWithEye() {
	fmt.Println("genTableXuShuWithEye start")

	curTable = &tableMgr.TableXuShuWithEye.Map
	curTableTemp = &tableXuShuWithEyeTemp
	curCardsTypeNum = 9

	cards := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i <= 8; i++ {
		cards[i] = 2
		fmt.Println("genTableXuShuWithEye jiang = ", i)
		addToXuShu(cards)
		genXuShuPuZi(cards, 1)
		cards[i] = 0
	}

	fmt.Println("genTableXuShuWithEye success")
}

func genXuShuPuZi(cards []int, level int) {

	if level > 4 {
		return
	}

	for i := 0; i <= 8; i++ {

		if cards[i] <= 3 {
			cards[i] += 3
			addToXuShu(cards)
			genXuShuPuZi(cards, level+1)
			cards[i] -= 3
		}

		if i+1 <= 8 && i+2 <= 8 && cards[i] <= 5 && cards[i+1] <= 5 && cards[i+2] <= 5 {
			cards[i]++
			cards[i+1]++
			cards[i+2]++
			addToXuShu(cards)
			genXuShuPuZi(cards, level+1)
			cards[i]--
			cards[i+1]--
			cards[i+2]--
		}
	}
}

func addToXuShu(cards []int) {
	if !checkAndAdd(cards, 0) {
		return
	}

	addToXuShuSub(cards, 1)
}

func addToXuShuSub(cards []int, iLaiZiNum int) {
	if iLaiZiNum >= LaiZiNum {
		return
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] == 0 {
			continue
		}

		cards[i]--
		if !checkAndAdd(cards, iLaiZiNum) {
			cards[i]++
			continue
		}

		addToXuShuSub(cards, iLaiZiNum+1)
		cards[i]++
	}
}

func checkAndAdd(cards []int, iLaiZiNum int) bool {

	key := 0
	for i := 0; i < curCardsTypeNum; i++ {
		key = key*10 + cards[i]
	}

	HandCardsMapTemp := curTableTemp[iLaiZiNum]
	_, exists := (*HandCardsMapTemp)[key]
	if exists {
		return false //这里说明这个情况处理过了，去重。
	}

	(*HandCardsMapTemp)[key] = true

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 4 {
			return true //这里用true是说这种情况不行，但是如果有赖子的话，还是可能可以的。
		}
	}

	HandCardsMap := curTable[iLaiZiNum]
	(*HandCardsMap)[key] = true
	return true
}

func genTableZi() {
	fmt.Println("genTableZi start")

	curTable = &tableMgr.TableZi.Map
	curTableTemp = &tableZiTemp
	curCardsTypeNum = 7
	cards := []int{0, 0, 0, 0, 0, 0, 0}
	genZiPuZi(cards, 1)

	fmt.Println("genTableZi success")
}

func genTableZiWithEye() {
	fmt.Println("genTableZiWithEye start")

	curTable = &tableMgr.TableZiWithEye.Map
	curTableTemp = &tableZiWithEyeTemp
	curCardsTypeNum = 7

	cards := []int{0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < curCardsTypeNum; i++ {
		cards[i] = 2
		fmt.Println("genTableZiWithEye jiang = ", i)
		addToXuShu(cards)
		genZiPuZi(cards, 1)
		cards[i] = 0
	}

	fmt.Println("genTableZiWithEye success")
}

func genZiPuZi(cards []int, level int) {
	if level > 4 {
		return
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 3 {
			continue
		}

		cards[i] += 3
		addToXuShu(cards)
		genZiPuZi(cards, level+1)
		cards[i] -= 3
	}
}
