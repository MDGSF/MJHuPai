package sxtjmj

import (
	"fmt"
	"log"
)

var tableMgr *TableMgr

func init() {
	tableMgr = NewTableMgr()
	tableMgr.Load("E:\\Go\\GOPATH\\src\\github.com\\MDGSF\\MJHuPai\\Go\\sxtjmj\\genTable\\")
	//tableMgr.Load(".\\")
}

/*
CanHu 胡牌
handCards: 手牌数组，最多14张。
huType: 胡牌类型，1自摸，2点炮。
huCard: 胡的那张牌，自摸的那张 或者是 点炮的那张。
HeiSanFeng: 是否开启黑三风

return:
	true可以胡牌，最大黑三风的数量。
	false不可以胡牌，0。
*/
func CanHu(handCards []int, huType int, huCard int, HeiSanFeng bool, laizi []int) (bool, int) {

	if !IsValidHandCards(handCards) {
		return false, 0
	}

	if isAllLaiZi(handCards, laizi) {
		return true, 0
	}

	TableXuShuWithEye := tableMgr.TableXuShuWithEye
	TableXuShu := tableMgr.TableXuShu
	TableFeng, TableFengWithEye := getFengTable(HeiSanFeng)
	TableJian := tableMgr.TableJianKe
	TableJianWithEye := tableMgr.TableJianKeWithEye

	allPossibleSlotHandCards := genAllPossibleSlotHandCards(handCards, laizi)

	log.Println("handCards = ", handCards)
	log.Println("laizi = ", laizi)
	log.Println("allPossibleSlotHandCards = ", allPossibleSlotHandCards)

	bCanHu := false
	maxFengNum := 0 //这个是黑三风的数量

	for _, v := range allPossibleSlotHandCards {
		XuShu, Feng, Jian := getArray(v.slots)

		if len(XuShu) > 0 {
			ret1, fengNum1 := walkThroughTableXuShuJiang(v.laiziNum, XuShu, Feng, Jian,
				TableXuShuWithEye, TableXuShu, TableFeng, TableJian, huType, huCard)
			if ret1 && fengNum1 >= maxFengNum {
				bCanHu = true
				maxFengNum = fengNum1
				log.Println("walkThroughTableXuShuJiang bCanHu = true, fengNum1 = ", fengNum1)
				log.Println(XuShu, Feng, Jian, v.laiziNum)
			}
		}

		if len(Feng) > 0 {
			ret2, fengNum2 := walkThroughTableFengJiang(v.laiziNum, Feng, XuShu, Jian,
				TableFengWithEye, TableXuShu, TableJian, huType, huCard)
			if ret2 && fengNum2 >= maxFengNum {
				bCanHu = true
				maxFengNum = fengNum2
				log.Println("walkThroughTableFengJiang bCanHu = true, fengNum2 = ", fengNum2)
				log.Println(XuShu, Feng, Jian, v.laiziNum)
			}
		}

		if len(Jian) > 0 {
			ret3, fengNum3 := walkThroughTableJianJiang(v.laiziNum, Jian, XuShu, Feng,
				TableJianWithEye, TableXuShu, TableFeng, huType, huCard)
			if ret3 && fengNum3 >= maxFengNum {
				bCanHu = true
				maxFengNum = fengNum3
				log.Println("walkThroughTableJianJiang bCanHu = true, fengNum3 = ", fengNum3)
				log.Println(XuShu, Feng, Jian, v.laiziNum)
			}
		}

		// log.Println("ret1 = ", ret1, ", fengNum1 = ", fengNum1, ", heiSanFengNum1 = ", heiSanFengNum1, ", zhongFaBaiNum1 = ", zhongFaBaiNum1)
		// log.Println("ret2 = ", ret2, ", fengNum2 = ", fengNum2, ", heiSanFengNum2 = ", heiSanFengNum2, ", zhongFaBaiNum2 = ", zhongFaBaiNum2)
		// log.Println("ret3 = ", ret3, ", fengNum3 = ", fengNum3, ", heiSanFengNum3 = ", heiSanFengNum3, ", zhongFaBaiNum3 = ", zhongFaBaiNum3)
	}

	// log.Println("bCanHu = ", bCanHu)
	// log.Println("maxHeiSanFengNum = ", maxHeiSanFengNum)
	// log.Println("maxZhongFaBaiNum = ", maxZhongFaBaiNum)
	// log.Println("maxFengNum = ", maxFengNum)

	return bCanHu, maxFengNum
}

func isAllLaiZi(handCards []int, laizi []int) bool {
	var bHasCommonCard = false
	laiziNum := 0
	for _, c := range handCards {
		if isLaizi(c, laizi) {
			laiziNum++
		} else {
			bHasCommonCard = true
		}
	}

	if !bHasCommonCard && laiziNum > 0 {
		//手牌全部都是赖子。
		return true
	}

	return false
}

func walkThroughTableJianJiang(laiziNum int, Jian []int, XuShu []int, Feng []int,
	TableJianWithEye *Table, TableXuShu *Table, TableFeng *Table,
	huType int, huCard int) (CanHu bool, FengNum int) {

	hasLaiZiNum := laiziNum
	maxFengNum := 0
	needLaiZiNum := 0
	ok := false

	if len(Jian) > 0 {
		_, needLaiZiNum, ok = TableJianWithEye.IsInTable(Jian[0])
		if !ok || needLaiZiNum > hasLaiZiNum {
			return false, 0
		}
		hasLaiZiNum -= needLaiZiNum
	}

	for _, num := range XuShu {
		_, needLaiZiNum, ok = TableXuShu.IsInTable(num)
		if !ok || needLaiZiNum > hasLaiZiNum {
			continue
		}
		hasLaiZiNum -= needLaiZiNum
	}

	if len(Feng) > 0 {
		if IsWind(huCard) {
			maxFengNum, needLaiZiNum, ok = TableFeng.IsValid(Feng[0], huType, huCard)
			if !ok || needLaiZiNum > hasLaiZiNum {
				return false, 0
			}
			hasLaiZiNum -= needLaiZiNum
		} else {
			maxFengNum, needLaiZiNum, ok = TableFeng.IsInTable(Feng[0])
			if !ok || needLaiZiNum > hasLaiZiNum {
				return false, 0
			}
			hasLaiZiNum -= needLaiZiNum
		}
	}

	return true, maxFengNum
}

func walkThroughTableFengJiang(laiziNum int, Feng []int,
	XuShu []int, Jian []int,
	TableFengWithEye *Table,
	TableXuShu *Table, TableJian *Table,
	huType int, huCard int) (CanHu bool, FengNum int) {

	hasLaiZiNum := laiziNum
	maxFengNum := 0
	needLaiZiNum := 0
	ok := false

	if len(Feng) > 0 {
		if IsWind(huCard) {
			maxFengNum, needLaiZiNum, ok = TableFengWithEye.IsValid(Feng[0], huType, huCard)
			if !ok || needLaiZiNum > hasLaiZiNum {
				return false, 0
			}
			hasLaiZiNum -= needLaiZiNum
		} else {
			maxFengNum, needLaiZiNum, ok = TableFengWithEye.IsInTable(Feng[0])
			if !ok || needLaiZiNum > hasLaiZiNum {
				return false, 0
			}
			hasLaiZiNum -= needLaiZiNum
		}
	}

	for _, num := range XuShu {
		_, needLaiZiNum, ok = TableXuShu.IsInTable(num)
		if !ok || needLaiZiNum > hasLaiZiNum {
			continue
		}
		hasLaiZiNum -= needLaiZiNum
	}

	if len(Jian) > 0 {
		_, needLaiZiNum, ok := TableJian.IsInTable(Jian[0])
		if !ok || needLaiZiNum > hasLaiZiNum {
			return false, 0
		}
		hasLaiZiNum -= needLaiZiNum
	}

	return true, maxFengNum
}

func walkThroughTableXuShuJiang(laiziNum int, XuShu []int,
	Feng []int, Jian []int,
	TableXuShuWithEye *Table, TableXuShu *Table,
	TableFeng *Table, TableJian *Table,
	huType int, huCard int) (CanHu bool, FengNum int) {

	bFound := false
	maxFengNum := 0

	for i, iNum := range XuShu {

		curFengNum := 0
		needLaiZiNum := 0
		hasLaiZiNum := laiziNum
		success := true

		_, jiangNeedLaiZiNum, ok := TableXuShuWithEye.IsInTable(iNum)
		if !ok || jiangNeedLaiZiNum > hasLaiZiNum {
			continue
		}
		hasLaiZiNum -= jiangNeedLaiZiNum

		for j, jNum := range XuShu {
			if i == j {
				continue
			}

			_, needLaiZiNum, ok = TableXuShu.IsInTable(jNum)
			if !ok || needLaiZiNum > hasLaiZiNum {
				success = false
				break
			}
			hasLaiZiNum -= needLaiZiNum
		}
		if !success {
			continue
		}

		if len(Feng) > 0 {
			if IsWind(huCard) {
				curFengNum, needLaiZiNum, ok = TableFeng.IsValid(Feng[0], huType, huCard)
				if !ok || needLaiZiNum > hasLaiZiNum {
					continue
				}
				hasLaiZiNum -= needLaiZiNum
			} else {
				curFengNum, needLaiZiNum, ok = TableFeng.IsInTable(Feng[0])
				if !ok || needLaiZiNum > hasLaiZiNum {
					continue
				}
				hasLaiZiNum -= needLaiZiNum
			}
		}

		if len(Jian) > 0 {
			_, needLaiZiNum, ok := TableJian.IsInTable(Jian[0])
			if !ok || needLaiZiNum > hasLaiZiNum {
				continue
			}
			hasLaiZiNum -= needLaiZiNum
		}

		if curFengNum > maxFengNum {
			maxFengNum = curFengNum
		}

		bFound = true
	}

	if bFound {
		return true, maxFengNum
	}
	return false, 0
}

func getFengTable(HeiSanFeng bool) (TableFeng, TableFengWithEye *Table) {
	if HeiSanFeng {
		TableFeng = tableMgr.TableFeng
		TableFengWithEye = tableMgr.TableFengWithEye
	} else {
		TableFeng = tableMgr.TableFengKe
		TableFengWithEye = tableMgr.TableFengKeWithEye
	}
	return
}

func getArray(slots [TILEMAX]int) (XuShu []int, Feng []int, Jian []int) {
	if getWan(slots) > 0 {
		XuShu = append(XuShu, getWan(slots))
	}

	if getTiao(slots) > 0 {
		XuShu = append(XuShu, getTiao(slots))
	}

	if getTong(slots) > 0 {
		XuShu = append(XuShu, getTong(slots))
	}

	if getFeng(slots) > 0 {
		Feng = append(Feng, getFeng(slots))
	}

	if getJian(slots) > 0 {
		Jian = append(Jian, getJian(slots))
	}

	return
}

/*
IsValidHandCards 判断手牌是否合法
handCards: 手牌数组，最多14张。
return: true手牌合法，false手牌非法。
*/
func IsValidHandCards(handCards []int) bool {
	cardsNum := len(handCards)
	if cardsNum <= 0 || cardsNum%3 != 2 || cardsNum > MaxHandCardNum {
		return false
	}

	for _, c := range handCards {
		if !IsValidCard(c) {
			return false
		}
	}

	slots := GenSlots(handCards)
	for _, num := range slots {
		if num > 4 {
			return false
		}
	}

	return true
}

/*
GenSlots 将手牌转换为对应的表示每张牌数量的数组。
例子：
handCards []int = {0x01, 0x01, 0x11, 0x12, 0x13, 0x33, 0x43, 0x43}
[TILEMAX]int = {
	2, 0, 0, 0, 0, 0, 0, 0, 0, //万
	1, 1, 1, 0, 0, 0, 0, 0, 0, //条
	0, 0, 0, 0, 0, 0, 0, 0, 0, //筒
	0, 0, 1, 0, //风
	0, 0, 2, //箭
}
*/
func GenSlots(handCards []int) [TILEMAX]int {
	var slots [TILEMAX]int

	for _, c := range handCards {
		slots[c]++
	}

	return slots
}

type slotHandCards struct {
	slots    [TILEMAX]int
	laiziNum int
}

func genFengKeys(slots [TILEMAX]int) int {
	key := slots[TON]*1000 + slots[NAN]*100 + slots[SHA]*10 + slots[PEI]
	return key
}

func genAllPossibleSlotHandCards(handCards []int, laizi []int) (allPossibleSlotHandCards []slotHandCards) {

	fengMap := make(map[int]int)
	slots := GenSlots(handCards)

	var a slotHandCards
	for i := 0; i < TILEMAX; i++ {
		if slots[i] > 0 {
			if isLaizi(i, laizi) {
				a.laiziNum += slots[i]
			} else {
				a.slots[i] = slots[i]
			}
		}
	}
	allPossibleSlotHandCards = append(allPossibleSlotHandCards, a)
	fengMap[genFengKeys(a.slots)] = 1

	ret := genFengAsLaiZi(TON, laizi, slots, 0, fengMap)
	allPossibleSlotHandCards = append(allPossibleSlotHandCards, ret...)

	return
}

func genFengAsLaiZi(curFeng int, laizi []int, slots [TILEMAX]int, laiziNum int, fengMap map[int]int) (allPossibleSlotHandCards []slotHandCards) {
	if curFeng > PEI {
		return
	}

	if isLaizi(curFeng, laizi) {

		for i := 0; i <= slots[curFeng]; i++ {

			fengNum := i
			fengAsLaiZiNum := slots[curFeng] - i

			backUpFengNum := slots[curFeng]
			backUpLaiZiNum := laiziNum

			slots[curFeng] = fengNum
			laiziNum = laiziNum + fengAsLaiZiNum

			if _, ok := fengMap[genFengKeys(slots)]; !ok {
				fengMap[genFengKeys(slots)] = 1
				var a slotHandCards
				a.laiziNum = laiziNum
				for j := 0; j < len(slots); j++ {
					a.slots[j] = slots[j]
				}
				allPossibleSlotHandCards = append(allPossibleSlotHandCards, a)

				for k := TON; k <= PEI; k++ {
					if k != curFeng {
						ret := genFengAsLaiZi(k, laizi, slots, laiziNum, fengMap)
						allPossibleSlotHandCards = append(allPossibleSlotHandCards, ret...)
					}
				}
			}

			slots[curFeng] = backUpFengNum
			laiziNum = backUpLaiZiNum
		}

	} else {
		ret := genFengAsLaiZi(curFeng+1, laizi, slots, laiziNum, fengMap)
		allPossibleSlotHandCards = append(allPossibleSlotHandCards, ret...)
	}

	return
}

func isLaizi(c int, LaiZiArr []int) bool {
	for _, laizi := range LaiZiArr {
		if c == laizi {
			return true
		}
	}
	return false
}

func getWan(slots [TILEMAX]int) int {
	return getNum(slots, CharacterOne, CharacterNine)
}

func getTiao(slots [TILEMAX]int) int {
	return getNum(slots, BambooOne, BambooNine)
}

func getTong(slots [TILEMAX]int) int {
	return getNum(slots, DotOne, DotNine)
}

func getFeng(slots [TILEMAX]int) int {
	return getNum(slots, EastWind, NorthWind)
}

func getJian(slots [TILEMAX]int) int {
	return getNum(slots, WhiteDragon, RedDragon)
}

func getNum(slots [TILEMAX]int, iStart int, iEnd int) int {
	num := 0
	for i := iStart; i <= iEnd; i++ {
		num = num*10 + int(slots[i])
	}
	return num
}

/*
ShowHandCards 打印手牌
*/
func ShowHandCards(handCards []int) {
	for _, c := range handCards {
		//fmt.Printf("0x%02X, ", c)
		fmt.Printf("%d, ", c)
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
核心思想：
1、名词解释：eye(将），字牌（feng、东南西北中发白），花色（万、筒、条、字牌）
2、分而治之：检查手牌是否能胡是依次检查万、筒、条、字牌四种花色是否能组成胡牌的一部分。
3、单一花色要能满足胡牌的部分，则要么是3*n(不带将)，要么是3*n+2（带将）。3*n中的3带表三张牌一样的刻子，或三张连续的牌如1筒2筒3筒。
4、判断是否满足胡牌的单一花色部分，需要根据是否有将，有几个赖子，查询不同的表。表内容表示表里的元素加上对应的赖子数量能组成3*n 或3*n+2。
赖子数是表名最后的数字，带有eye的表表示满足3*n+2，没有的表示满足3*n。
5、查表的key值，是直接根据1-9有几张牌就填几（赖子不算），如1-9万各一张，则key为111111111。如1万3张，9万2张，则key为300000002。
6、组合多种花色，判断是否能胡牌。将赖子分配给不同的花色，有若干种分配方式，只要有一种分配能让所有花色满足单一花色胡牌部分，则手牌能胡。
如：手上有3个赖子，可以分配万、筒、条各一张，也可以万、同、字牌各一张
7、根据是否有将、是否字牌分为4种表，每种表又根据赖子个数0-8分别建表，共36张表，具体如下：

赖子个数     带将表        不带将的表         字牌带将表             字牌不带将表
0        eye_table_0     table_0       feng_eye_table_0        feng_table_0
1        eye_table_1     table_1       feng_eye_table_0        feng_table_1
2        eye_table_2     table_2       feng_eye_table_0        feng_table_2
3        eye_table_3     table_3       feng_eye_table_0        feng_table_3
4        eye_table_4     table_4       feng_eye_table_0        feng_table_4
5        eye_table_5     table_5       feng_eye_table_0        feng_table_5
6        eye_table_6     table_6       feng_eye_table_0        feng_table_6
7        eye_table_7     table_7       feng_eye_table_0        feng_table_7
8        eye_table_8     table_8       feng_eye_table_0        feng_table_8

步骤：
1、统计手牌中鬼牌个数 nGui，将鬼牌从牌数据中去除.
2、不同花色分开处理，分别校验是否能满足 将、顺子、刻子
3、分析东南西北风中发白时，不需要分析顺子的情况，简单很多
4、分析单一花色时，直接根据1-9点对应数字得出一个9位的整数，每位上为0-4代表该点数上有几张牌
比如：1筒2筒3筒3筒3筒3筒6筒7筒8筒2万3万3万3万4万
数字：筒: 1,1,4,0,0,1,1,1,0 得出的数字为114001110
      万 0,1,3,1,0,0,0,0,0 得出的数字为13100000
5、组合多种花色，判断是否能胡牌。将赖子分配给不同的花色，有若干种分配方式，只要有一种分配能让所有花色满足单一花色胡牌部分，则手牌能胡。
如：手上有3个赖子，可以分配万、筒、条各一张，也可以万、同、字牌各一张
每种花色与赖子组合，如果所有花色都能配型成功则可胡牌
检查配型时，每种花色的牌数量必需是3*n 或者 3*n + 2
根据赖子个数、带不带将，查找对应表，看能否满足3*n 或 3*n+2的牌型

非字牌表的产生:
1、穷举万字牌所有满足胡牌胡可能，将对应的牌型记录为数字，根据是否有将、放入eye_table_0或table_0中。
  具体是每次加入一个刻子，顺子或是将（将只能加入一对）,最多加入四组外带将牌。
2、将table_0中牌去掉一张，放入table_1中，表示去掉的牌用1张赖子代替。eye_table_0中牌去掉一张，放入到eye_table_1中。
3、将table_1中牌去掉一张，放入table_2中，表示去掉的牌用1张赖子代替。eye_table_1中牌去掉一张，放入到eye_table_2中。
4、将table_2中牌去掉一张，放入table_3中，表示去掉的牌用1张赖子代替。eye_table_2中牌去掉一张，放入到eye_table_3中。
5、将table_3中牌去掉一张，放入table_4中，表示去掉的牌用1张赖子代替。eye_table_3中牌去掉一张，放入到eye_table_4中。
6、将table_4中牌去掉一张，放入table_5中，表示去掉的牌用1张赖子代替。eye_table_4中牌去掉一张，放入到eye_table_5中。
7、将table_5中牌去掉一张，放入table_6中，表示去掉的牌用1张赖子代替。eye_table_5中牌去掉一张，放入到eye_table_6中。
8、将table_6中牌去掉一张，放入table_7中，表示去掉的牌用1张赖子代替。eye_table_6中牌去掉一张，放入到eye_table_7中。
9、将table_7中牌去掉一张，放入table_8中，表示去掉的牌用1张赖子代替。eye_table_7中牌去掉一张，放入到eye_table_8中。

字牌表的产生：
与非字牌表的产生方法相同，只是第一步中，不能加入顺子（除非麻将玩法字牌是能组成顺子的）

表的大小：总量在2M左右

表生成耗时：2-3S
*/
