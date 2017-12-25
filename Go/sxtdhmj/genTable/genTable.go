package main

import (
	"fmt"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

type HandCardsList struct {
	CommonList     []*sxtdhmj.PuZi //保存这副牌中，除了黑三风和中发白的其他牌
	HeiSanFengList []*sxtdhmj.PuZi //保存黑三风和中发白
}

func (cardList *HandCardsList) addNoFeng(c1 int, c2 int, c3 int) {
	puzi := &sxtdhmj.PuZi{}
	puzi.PuZi[0] = c1
	puzi.PuZi[1] = c2
	puzi.PuZi[2] = c3
	cardList.CommonList = append(cardList.CommonList, puzi)
}

func (cardList *HandCardsList) removeNoFeng(c1 int, c2 int, c3 int) {
	found := false
	index := 0
	for i, v := range cardList.CommonList {
		if v.PuZi[0] == c1 && v.PuZi[1] == c2 && v.PuZi[2] == c3 {
			found = true
			index = i
		}
	}

	if found {
		cardList.CommonList = append(cardList.CommonList[:index], cardList.CommonList[index+1:]...)
	}
}

func (cardList *HandCardsList) addFeng(c1 int, c2 int, c3 int) {
	puzi := &sxtdhmj.PuZi{}
	puzi.PuZi[0] = c1
	puzi.PuZi[1] = c2
	puzi.PuZi[2] = c3
	cardList.HeiSanFengList = append(cardList.HeiSanFengList, puzi)
}

func (cardList *HandCardsList) removeFeng(c1 int, c2 int, c3 int) {
	found := false
	index := 0
	for i, v := range cardList.HeiSanFengList {
		if v.PuZi[0] == c1 && v.PuZi[1] == c2 && v.PuZi[2] == c3 {
			found = true
			index = i
		}
	}

	if found {
		cardList.HeiSanFengList = append(cardList.HeiSanFengList[:index], cardList.HeiSanFengList[index+1:]...)
	}
}

const LaiZiNum = 9

//只保存最终正确的结果。
var tableMgr *sxtdhmj.TableMgr

//除了保存最终正确的结果，还要把中间计算过的错误结果也保存起来，防止重复计算。
var tableXuShuTemp [LaiZiNum]*map[int]*sxtdhmj.TdhValue
var tableXuShuWithEyeTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}

//风牌
var tableFengKeTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}
var tableFengKeWithEyeTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}
var tableFengTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}
var tableFengWithEyeTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}

//箭牌
var tableJianKeTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}
var tableJianKeWithEyeTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}
var tableJianTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}
var tableJianWithEyeTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}

var tableZiTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}
var tableZiWithEyeTemp = [LaiZiNum]*map[int]*sxtdhmj.TdhValue{}

var curTable *[LaiZiNum]*map[int]*sxtdhmj.TdhValue
var curTableTemp *[LaiZiNum]*map[int]*sxtdhmj.TdhValue
var curCardsTypeNum int

var heiSanFeng []*sxtdhmj.PuZi

func genHeiSanFengAllPossible() {
	hsfPuZi1 := &sxtdhmj.PuZi{}
	hsfPuZi1.PuZi[0] = sxtdhmj.TON
	hsfPuZi1.PuZi[1] = sxtdhmj.NAN
	hsfPuZi1.PuZi[2] = sxtdhmj.SHA
	heiSanFeng = append(heiSanFeng, hsfPuZi1)

	hsfPuZi2 := &sxtdhmj.PuZi{}
	hsfPuZi2.PuZi[0] = sxtdhmj.TON
	hsfPuZi2.PuZi[1] = sxtdhmj.NAN
	hsfPuZi2.PuZi[2] = sxtdhmj.PEI
	heiSanFeng = append(heiSanFeng, hsfPuZi2)

	hsfPuZi3 := &sxtdhmj.PuZi{}
	hsfPuZi3.PuZi[0] = sxtdhmj.TON
	hsfPuZi3.PuZi[1] = sxtdhmj.SHA
	hsfPuZi3.PuZi[2] = sxtdhmj.PEI
	heiSanFeng = append(heiSanFeng, hsfPuZi3)

	hsfPuZi4 := &sxtdhmj.PuZi{}
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
		tableXuShuTemp[i] = &map[int]*sxtdhmj.TdhValue{}
		tableXuShuWithEyeTemp[i] = &map[int]*sxtdhmj.TdhValue{}

		tableFengKeTemp[i] = &map[int]*sxtdhmj.TdhValue{}
		tableFengKeWithEyeTemp[i] = &map[int]*sxtdhmj.TdhValue{}
		tableFengTemp[i] = &map[int]*sxtdhmj.TdhValue{}
		tableFengWithEyeTemp[i] = &map[int]*sxtdhmj.TdhValue{}

		tableJianKeTemp[i] = &map[int]*sxtdhmj.TdhValue{}
		tableJianKeWithEyeTemp[i] = &map[int]*sxtdhmj.TdhValue{}
		tableJianTemp[i] = &map[int]*sxtdhmj.TdhValue{}
		tableJianWithEyeTemp[i] = &map[int]*sxtdhmj.TdhValue{}

		tableZiTemp[i] = &map[int]*sxtdhmj.TdhValue{}
		tableZiWithEyeTemp[i] = &map[int]*sxtdhmj.TdhValue{}
	}

	// genTableXuShu()
	// genTableXuShuWithEye()

	// genTableFengKe()
	// genTableFengKeWithEye()
	// genTableFeng()
	// genTableFengWithEye()

	// genTableJianKe()
	// genTableJianKeWithEye()
	// genTableJian()
	// genTableJianWithEye()

	// genTableZi()
	// genTableZiWithEye()

	genTableFeng()
	genTableFengWithEye()
	genTableJian()
	genTableJianWithEye()

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
	cardsList := &HandCardsList{}
	genHeiSanFengPuZi(cardsList, cards, 1, 0)

	fmt.Println("genTableFeng success")
}

func genTableFengWithEye() {
	fmt.Println("genTableFengWithEye start")

	curTable = &tableMgr.TableFengWithEye.Map
	curTableTemp = &tableFengWithEyeTemp
	curCardsTypeNum = 4

	cards := []int{0, 0, 0, 0}
	cardsList := &HandCardsList{}

	for i := 0; i < curCardsTypeNum; i++ {
		cards[i] = 2
		cardsList.addNoFeng(sxtdhmj.TON+i, sxtdhmj.TON+i, 0)
		//fmt.Println("genTableZiWithEye jiang = ", i)

		checkAndAddHeiSanFeng(cardsList, cards, 0, 0)
		genHeiSanFengPuZi(cardsList, cards, 1, 0)

		cardsList.removeNoFeng(sxtdhmj.TON+i, sxtdhmj.TON+i, 0)
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
	cardsList := &HandCardsList{}

	genZhongFaBaiPuZi(cardsList, cards, 1, 0)

	fmt.Println("genTableJian success")
}

func genTableJianWithEye() {
	fmt.Println("genTableJianWithEye start")

	curTable = &tableMgr.TableJianWithEye.Map
	curTableTemp = &tableJianWithEyeTemp
	curCardsTypeNum = 3

	cards := []int{0, 0, 0}
	cardsList := &HandCardsList{}

	for i := 0; i < curCardsTypeNum; i++ {
		cards[i] = 2
		cardsList.addNoFeng(sxtdhmj.HAK+i, sxtdhmj.HAK+i, 0)

		checkAndAddHeiSanFeng(cardsList, cards, 0, 0)
		genZhongFaBaiPuZi(cardsList, cards, 1, 0)

		cardsList.removeNoFeng(sxtdhmj.HAK+i, sxtdhmj.HAK+i, 0)
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

func genHeiSanFengPuZi(cardsList *HandCardsList, cards []int, level int, heiSanFengNum int) {
	if level > 4 {
		return
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 3 {
			continue
		}

		cards[i] += 3
		cardsList.addNoFeng(sxtdhmj.TON+i, sxtdhmj.TON+i, sxtdhmj.TON+i)

		checkAndAddHeiSanFeng(cardsList, cards, 0, heiSanFengNum)
		genHeiSanFengPuZi(cardsList, cards, level+1, heiSanFengNum)

		cardsList.removeNoFeng(sxtdhmj.TON+i, sxtdhmj.TON+i, sxtdhmj.TON+i)
		cards[i] -= 3
	}

	for _, v := range heiSanFeng {
		cards[v.PuZi[0]-sxtdhmj.TON]++
		cards[v.PuZi[1]-sxtdhmj.TON]++
		cards[v.PuZi[2]-sxtdhmj.TON]++
		cardsList.addFeng(v.PuZi[0], v.PuZi[1], v.PuZi[2])

		checkAndAddHeiSanFeng(cardsList, cards, 0, heiSanFengNum+1)
		genHeiSanFengPuZi(cardsList, cards, level+1, heiSanFengNum+1)

		cardsList.removeFeng(v.PuZi[0], v.PuZi[1], v.PuZi[2])
		cards[v.PuZi[0]-sxtdhmj.TON]--
		cards[v.PuZi[1]-sxtdhmj.TON]--
		cards[v.PuZi[2]-sxtdhmj.TON]--
	}
}

func genZhongFaBaiPuZi(cardsList *HandCardsList, cards []int, level int, zhongFaBaiNum int) {
	if level > 4 {
		return
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 3 {
			continue
		}

		cards[i] += 3
		cardsList.addNoFeng(sxtdhmj.HAK+i, sxtdhmj.HAK+i, sxtdhmj.HAK+i)

		checkAndAddHeiSanFeng(cardsList, cards, 0, zhongFaBaiNum)
		genZhongFaBaiPuZi(cardsList, cards, level+1, zhongFaBaiNum)

		cardsList.removeNoFeng(sxtdhmj.HAK+i, sxtdhmj.HAK+i, sxtdhmj.HAK+i)
		cards[i] -= 3
	}

	cards[0]++
	cards[1]++
	cards[2]++
	cardsList.addFeng(sxtdhmj.HAK+0, sxtdhmj.HAK+1, sxtdhmj.HAK+2)

	checkAndAddHeiSanFeng(cardsList, cards, 0, zhongFaBaiNum+1)
	genZhongFaBaiPuZi(cardsList, cards, level+1, zhongFaBaiNum+1)

	cardsList.removeFeng(sxtdhmj.HAK+0, sxtdhmj.HAK+1, sxtdhmj.HAK+2)
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

	v1 := &sxtdhmj.TdhValue{}
	(*HandCardsMapTemp)[key] = v1

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 4 {
			return true //这里用true是说这种情况不行，但是如果有赖子的话，还是可能可以的。
		}
	}

	HandCardsMap := curTable[iLaiZiNum]
	v2 := &sxtdhmj.TdhValue{}
	(*HandCardsMap)[key] = v2
	return true
}

func checkAndAddHeiSanFeng(cardsList *HandCardsList, cards []int, iLaiZiNum int, heiSanFengNum int) bool {

	key := 0
	for i := 0; i < curCardsTypeNum; i++ {
		key = key*10 + cards[i]
	}

	ziMoList := calcZiMoList(cardsList)

	HandCardsMapTemp := curTableTemp[iLaiZiNum]
	v, exists := (*HandCardsMapTemp)[key]
	if exists {
		if heiSanFengNum <= v.FengNum && len(cardsList.CommonList) == 0 {
			return false //这里说明这个情况处理过了，去重。
		}
		(*HandCardsMapTemp)[key].FengNum = heiSanFengNum
		//(*HandCardsMapTemp)[key].ZiMoList = ziMoList
	} else {
		v1 := &sxtdhmj.TdhValue{}
		v1.FengNum = heiSanFengNum
		//v1.ZiMoList = ziMoList
		(*HandCardsMapTemp)[key] = v1
	}

	for i := 0; i < curCardsTypeNum; i++ {
		if cards[i] > 4 {
			return true //这里用true是说这种情况不行，但是如果有赖子的话，还是可能可以的。
		}
	}

	if key == 4301 {
		fmt.Printf("cardsList.CommonList = ")
		for _, v := range cardsList.CommonList {
			fmt.Printf("%d ", v)
		}
		fmt.Println()

		fmt.Printf("cardsList.HeiSanFengList = ")
		for _, v := range cardsList.HeiSanFengList {
			fmt.Printf("%d ", v)
		}
		fmt.Println()

		fmt.Println("(*curTable[iLaiZiNum])[key] = ", (*curTable[iLaiZiNum])[key])
		fmt.Println("----------------------------------------------------")
	}

	HandCardsMap := curTable[iLaiZiNum]
	if exists {
		(*HandCardsMap)[key].FengNum = heiSanFengNum
		mofidyZiMoList(cardsList, (*HandCardsMap)[key])
		//(*HandCardsMap)[key].ZiMoList = ziMoList
	} else {
		v2 := &sxtdhmj.TdhValue{}
		v2.FengNum = heiSanFengNum
		v2.ZiMoList = ziMoList
		(*HandCardsMap)[key] = v2
	}
	return true
}

func mofidyZiMoList(cardsList *HandCardsList, v *sxtdhmj.TdhValue) {
	if len(cardsList.CommonList) == 0 {
		return
	}

	slots := [sxtdhmj.TILEMAX]int{}
	for _, v := range v.ZiMoList {
		slots[v] = 1
	}

	for _, v := range cardsList.CommonList {
		slots[v.PuZi[0]] = 0
		slots[v.PuZi[1]] = 0
		slots[v.PuZi[2]] = 0
	}

	ziMoList := []int{}
	for k, v := range slots {
		if v == 1 {
			ziMoList = append(ziMoList, k)
		}
	}

	v.ZiMoList = nil
	v.ZiMoList = append(v.ZiMoList, ziMoList...)
}

func calcZiMoList(cardsList *HandCardsList) []int {

	if len(cardsList.HeiSanFengList) == 0 {
		return nil
	}

	slots := [sxtdhmj.TILEMAX]int{}
	for _, v := range cardsList.HeiSanFengList {
		//log.Println("HeiSanFengList = ", v)
		slots[v.PuZi[0]] = 1
		slots[v.PuZi[1]] = 1
		slots[v.PuZi[2]] = 1
	}

	for _, v := range cardsList.CommonList {
		//log.Println("CommonList = ", v)
		slots[v.PuZi[0]] = 0
		slots[v.PuZi[1]] = 0
		slots[v.PuZi[2]] = 0
	}

	ziMoList := []int{}
	for k, v := range slots {
		if v == 1 {
			ziMoList = append(ziMoList, k)
		}
	}

	return ziMoList
}
