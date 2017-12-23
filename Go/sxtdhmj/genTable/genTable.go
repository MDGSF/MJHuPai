package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

const laiZiNum = 9

//只保存最终正确的结果。
var tableMgr *sxtdhmj.TableMgr

//除了保存最终正确的结果，还要把中间计算过的错误结果也保存起来，防止重复计算。
var tableXuShuTemp [laiZiNum]*map[int]int
var tableXuShuWithEyeTemp = [laiZiNum]*map[int]int{}
var tableZiTemp = [laiZiNum]*map[int]int{}
var tableZiWithEyeTemp = [laiZiNum]*map[int]int{}

var curTable *[laiZiNum]*map[int]int
var curTableTemp *[laiZiNum]*map[int]int
var curCardsTypeNum int

func main() {

	fmt.Println("main start")

	tableMgr = sxtdhmj.NewTableMgr()

	for i := 0; i < laiZiNum; i++ {
		tableXuShuTemp[i] = &map[int]int{}
		tableXuShuWithEyeTemp[i] = &map[int]int{}
		tableZiTemp[i] = &map[int]int{}
		tableZiWithEyeTemp[i] = &map[int]int{}
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
	ret, _ := checkAndAdd(cards, 0, 0, 0)
	if !ret {
		return
	}

	addToXuShuSub(cards, 1, 0)
}

func addToXuShuSub(cards []int, iLaiZiNum int, parentDianShu int) {
	if iLaiZiNum >= laiZiNum {
		return
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] == 0 {
			continue
		}

		cards[i]--
		ret, dianshu := checkAndAdd(cards, iLaiZiNum, i, parentDianShu)
		if !ret {
			cards[i]++
			continue
		}

		addToXuShuSub(cards, iLaiZiNum+1, dianshu)
		cards[i]++
	}
}

/*
@brief checkAndAdd:
@param cards:
@param iLaiZiNum: 赖子的数量。
@param laiZiStandFor: 当前这张赖子代替什么牌。当iLaiZiNum==0时，也就是没有赖子的时候，这个值没有用，填零就好了。
@param parentDianShu: 上一级的点数。
@return bool: true添加成功, false已经添加过了。
@return int: 成功时返回点数，失败时返回0。
*/
func checkAndAdd(cards []int, iLaiZiNum int, laiZiStandFor int, parentDianShu int) (bool, int) {
	if curCardsTypeNum == 9 {
		return checkAndAddXuShu(cards, iLaiZiNum, laiZiStandFor, parentDianShu)
	} else if curCardsTypeNum == 7 {
		return checkAndAddZi(cards, iLaiZiNum, laiZiStandFor, parentDianShu)
	}
	return false, 0
}

func checkAndAddXuShu(cards []int, iLaiZiNum int, laiZiStandFor int, parentDianShu int) (bool, int) {
	key := 0
	for i := 0; i < curCardsTypeNum; i++ {
		key = key*10 + cards[i]
	}

	newDianShu := max(laiZiStandFor+1, parentDianShu)

	HandCardsMapTemp := curTableTemp[iLaiZiNum]
	oldDianShu, exists := (*HandCardsMapTemp)[key]
	if exists && newDianShu <= oldDianShu {
		return false, 0 //这里说明这个情况处理过了，并且新的点数没有比旧的更大，去重。
	}

	if iLaiZiNum == 0 {
		(*HandCardsMapTemp)[key] = 0
	} else if iLaiZiNum == 1 {
		(*HandCardsMapTemp)[key] = laiZiStandFor + 1
	} else {
		(*HandCardsMapTemp)[key] = newDianShu
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 4 {
			return true, 0 //这里用true是说这种情况不行，但是如果有赖子的话，还是可能可以的。
		}
	}

	HandCardsMap := curTable[iLaiZiNum]
	if iLaiZiNum == 0 {
		(*HandCardsMap)[key] = 0
	} else if iLaiZiNum == 1 {
		(*HandCardsMap)[key] = laiZiStandFor + 1
	} else {
		(*HandCardsMap)[key] = newDianShu
	}
	return true, (*HandCardsMap)[key]
}

func checkAndAddZi(cards []int, iLaiZiNum int, laiZiStandFor int, parentDianShu int) (bool, int) {
	key := 0
	for i := 0; i < curCardsTypeNum; i++ {
		key = key*10 + cards[i]
	}

	HandCardsMapTemp := curTableTemp[iLaiZiNum]
	_, exists := (*HandCardsMapTemp)[key]
	if exists {
		return false, 0 //这里说明这个情况处理过了，并且新的点数没有比旧的更大，去重。
	}

	if iLaiZiNum == 0 {
		(*HandCardsMapTemp)[key] = 0
	} else {
		(*HandCardsMapTemp)[key] = 10
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 4 {
			return true, 0 //这里用true是说这种情况不行，但是如果有赖子的话，还是可能可以的。
		}
	}

	HandCardsMap := curTable[iLaiZiNum]
	if iLaiZiNum == 0 {
		(*HandCardsMap)[key] = 0
	} else {
		(*HandCardsMap)[key] = 10
	}
	return true, (*HandCardsMap)[key]
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
