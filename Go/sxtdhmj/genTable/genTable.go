package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

const LaiZiNum = 9

//只保存最终正确的结果。
var tableMgr *sxtdhmj.TableMgr

//除了保存最终正确的结果，还要把中间计算过的错误结果也保存起来，防止重复计算。
var tableXuShuTemp [LaiZiNum]*map[int]int
var tableXuShuWithEyeTemp = [LaiZiNum]*map[int]int{}

//风牌
var tableFengKeTemp = [LaiZiNum]*map[int]int{}
var tableFengKeWithEyeTemp = [LaiZiNum]*map[int]int{}
var tableFengTemp = [LaiZiNum]*map[int]int{}
var tableFengWithEyeTemp = [LaiZiNum]*map[int]int{}

//箭牌
var tableJianKeTemp = [LaiZiNum]*map[int]int{}
var tableJianKeWithEyeTemp = [LaiZiNum]*map[int]int{}
var tableJianTemp = [LaiZiNum]*map[int]int{}
var tableJianWithEyeTemp = [LaiZiNum]*map[int]int{}

var tableZiTemp = [LaiZiNum]*map[int]int{}
var tableZiWithEyeTemp = [LaiZiNum]*map[int]int{}

var curTable *[LaiZiNum]*map[int]int
var curTableTemp *[LaiZiNum]*map[int]int
var curCardsTypeNum int

type PuZi struct {
	PuZi [3]int
}

var heiSanFeng []*PuZi

func genHeiSanFengAllPossible() {
	hsfPuZi1 := &PuZi{}
	hsfPuZi1.PuZi[0] = sxtdhmj.TON
	hsfPuZi1.PuZi[1] = sxtdhmj.NAN
	hsfPuZi1.PuZi[2] = sxtdhmj.SHA
	heiSanFeng = append(heiSanFeng, hsfPuZi1)

	hsfPuZi2 := &PuZi{}
	hsfPuZi2.PuZi[0] = sxtdhmj.TON
	hsfPuZi2.PuZi[1] = sxtdhmj.NAN
	hsfPuZi2.PuZi[2] = sxtdhmj.PEI
	heiSanFeng = append(heiSanFeng, hsfPuZi2)

	hsfPuZi3 := &PuZi{}
	hsfPuZi3.PuZi[0] = sxtdhmj.TON
	hsfPuZi3.PuZi[1] = sxtdhmj.SHA
	hsfPuZi3.PuZi[2] = sxtdhmj.PEI
	heiSanFeng = append(heiSanFeng, hsfPuZi3)

	hsfPuZi4 := &PuZi{}
	hsfPuZi4.PuZi[0] = sxtdhmj.SHA
	hsfPuZi4.PuZi[1] = sxtdhmj.NAN
	hsfPuZi4.PuZi[2] = sxtdhmj.PEI
	heiSanFeng = append(heiSanFeng, hsfPuZi4)
}

func main() {

	fmt.Println("main start")

	genHeiSanFengAllPossible()

	tableMgr = sxtdhmj.NewTableMgr()

	for i := 0; i < LaiZiNum; i++ {
		tableXuShuTemp[i] = &map[int]int{}
		tableXuShuWithEyeTemp[i] = &map[int]int{}

		tableFengKeTemp[i] = &map[int]int{}
		tableFengKeWithEyeTemp[i] = &map[int]int{}
		tableFengTemp[i] = &map[int]int{}
		tableFengWithEyeTemp[i] = &map[int]int{}

		tableJianKeTemp[i] = &map[int]int{}
		tableJianKeWithEyeTemp[i] = &map[int]int{}
		tableJianTemp[i] = &map[int]int{}
		tableJianWithEyeTemp[i] = &map[int]int{}

		tableZiTemp[i] = &map[int]int{}
		tableZiWithEyeTemp[i] = &map[int]int{}
	}

	genTableXuShu()
	genTableXuShuWithEye()

	genTableFengKe()
	genTableFengKeWithEye()
	genTableFeng()
	genTableFengWithEye()

	genTableJianKe()
	genTableJianKeWithEye()
	genTableJian()
	genTableJianWithEye()

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
		//fmt.Println("genTableXuShuWithEye jiang = ", i)
		addToXuShu(cards)
		genXuShuPuZi(cards, 1)
		cards[i] = 0
	}

	fmt.Println("genTableXuShuWithEye success")
}

func genTableFengKe() {
	fmt.Println("genTableFengKe start")

	curTable = &tableMgr.TableFengKe.Map
	curTableTemp = &tableFengKeTemp
	curCardsTypeNum = 4
	cards := []int{0, 0, 0, 0}
	genZiPuZi(cards, 1)

	fmt.Println("genTableFengKe success")
}

func genTableFengKeWithEye() {
	fmt.Println("genTableFengKeWithEye start")

	curTable = &tableMgr.TableFengKeWithEye.Map
	curTableTemp = &tableFengKeWithEyeTemp
	curCardsTypeNum = 4

	cards := []int{0, 0, 0, 0}
	for i := 0; i < curCardsTypeNum; i++ {
		cards[i] = 2
		//fmt.Println("genTableZiWithEye jiang = ", i)
		addToXuShu(cards)
		genZiPuZi(cards, 1)
		cards[i] = 0
	}

	fmt.Println("genTableFengKeWithEye success")
}

func genTableFeng() {
	fmt.Println("genTableFeng start")

	curTable = &tableMgr.TableFeng.Map
	curTableTemp = &tableFengTemp
	curCardsTypeNum = 4
	cards := []int{0, 0, 0, 0}
	genHeiSanFengPuZi(cards, 1, 0)

	fmt.Println("genTableFeng success")
}

func genTableFengWithEye() {
	fmt.Println("genTableFengWithEye start")

	curTable = &tableMgr.TableFengWithEye.Map
	curTableTemp = &tableFengWithEyeTemp
	curCardsTypeNum = 4

	cards := []int{0, 0, 0, 0}
	for i := 0; i < curCardsTypeNum; i++ {
		cards[i] = 2
		//fmt.Println("genTableZiWithEye jiang = ", i)
		checkAndAddHeiSanFeng(cards, 0, 0)
		genHeiSanFengPuZi(cards, 1, 0)
		cards[i] = 0
	}

	fmt.Println("genTableFengWithEye success")
}

func genTableJianKe() {
	fmt.Println("genTableJianKe start")

	curTable = &tableMgr.TableJianKe.Map
	curTableTemp = &tableJianKeTemp
	curCardsTypeNum = 3
	cards := []int{0, 0, 0}
	genZiPuZi(cards, 1)

	fmt.Println("genTableJianKe success")
}

func genTableJianKeWithEye() {
	fmt.Println("genTableJianKeWithEye start")

	curTable = &tableMgr.TableJianKeWithEye.Map
	curTableTemp = &tableJianKeWithEyeTemp
	curCardsTypeNum = 3

	cards := []int{0, 0, 0}
	for i := 0; i < curCardsTypeNum; i++ {
		cards[i] = 2
		//fmt.Println("genTableZiWithEye jiang = ", i)
		addToXuShu(cards)
		genZiPuZi(cards, 1)
		cards[i] = 0
	}

	fmt.Println("genTableJianKeWithEye success")
}

func genTableJian() {
	fmt.Println("genTableJian start")

	curTable = &tableMgr.TableJian.Map
	curTableTemp = &tableJianTemp
	curCardsTypeNum = 3
	cards := []int{0, 0, 0}
	genZhongFaBaiPuZi(cards, 1, 0)

	fmt.Println("genTableJian success")
}

func genTableJianWithEye() {
	fmt.Println("genTableJianWithEye start")

	curTable = &tableMgr.TableJianWithEye.Map
	curTableTemp = &tableJianWithEyeTemp
	curCardsTypeNum = 4

	cards := []int{0, 0, 0, 0}
	for i := 0; i < curCardsTypeNum; i++ {
		cards[i] = 2
		//fmt.Println("genTableZiWithEye jiang = ", i)
		checkAndAddHeiSanFeng(cards, 0, 0)
		genZhongFaBaiPuZi(cards, 1, 0)
		cards[i] = 0
	}

	fmt.Println("genTableJianWithEye success")
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
		//fmt.Println("genTableZiWithEye jiang = ", i)
		addToXuShu(cards)
		genZiPuZi(cards, 1)
		cards[i] = 0
	}

	fmt.Println("genTableZiWithEye success")
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

func genHeiSanFengPuZi(cards []int, level int, heiSanFengNum int) {
	if level > 4 {
		return
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 3 {
			continue
		}

		cards[i] += 3
		checkAndAddHeiSanFeng(cards, 0, heiSanFengNum)
		genHeiSanFengPuZi(cards, level+1, heiSanFengNum)
		cards[i] -= 3
	}

	for _, v := range heiSanFeng {
		cards[v.PuZi[0]-sxtdhmj.TON]++
		cards[v.PuZi[1]-sxtdhmj.TON]++
		cards[v.PuZi[2]-sxtdhmj.TON]++
		checkAndAddHeiSanFeng(cards, 0, heiSanFengNum+1)
		genHeiSanFengPuZi(cards, level+1, heiSanFengNum+1)
		cards[v.PuZi[0]-sxtdhmj.TON]--
		cards[v.PuZi[1]-sxtdhmj.TON]--
		cards[v.PuZi[2]-sxtdhmj.TON]--
	}
}

func genZhongFaBaiPuZi(cards []int, level int, zhongFaBaiNum int) {
	if level > 4 {
		return
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 3 {
			continue
		}

		cards[i] += 3
		checkAndAddHeiSanFeng(cards, 0, zhongFaBaiNum)
		genZhongFaBaiPuZi(cards, level+1, zhongFaBaiNum)
		cards[i] -= 3
	}

	cards[0]++
	cards[1]++
	cards[2]++
	checkAndAddHeiSanFeng(cards, 0, zhongFaBaiNum+1)
	genZhongFaBaiPuZi(cards, level+1, zhongFaBaiNum+1)
	cards[0]--
	cards[1]--
	cards[2]--
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

func addToXuShu(cards []int) {
	if !checkAndAdd(cards, 0) {
		return
	}

	//没有赖子，暂时不需要这些，先注释掉
	//addToXuShuSub(cards, 1)
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

	(*HandCardsMapTemp)[key] = 0

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 4 {
			return true //这里用true是说这种情况不行，但是如果有赖子的话，还是可能可以的。
		}
	}

	HandCardsMap := curTable[iLaiZiNum]
	(*HandCardsMap)[key] = 0
	return true
}

func checkAndAddHeiSanFeng(cards []int, iLaiZiNum int, heiSanFengNum int) bool {

	key := 0
	for i := 0; i < curCardsTypeNum; i++ {
		key = key*10 + cards[i]
	}

	HandCardsMapTemp := curTableTemp[iLaiZiNum]
	v, exists := (*HandCardsMapTemp)[key]
	if exists && heiSanFengNum <= v {
		return false //这里说明这个情况处理过了，去重。
	}

	(*HandCardsMapTemp)[key] = heiSanFengNum

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 4 {
			return true //这里用true是说这种情况不行，但是如果有赖子的话，还是可能可以的。
		}
	}

	HandCardsMap := curTable[iLaiZiNum]
	(*HandCardsMap)[key] = heiSanFengNum
	return true
}
