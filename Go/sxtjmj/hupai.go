package sxtjmj

import (
	"fmt"
)

var tableMgr *TableMgr

func init() {
	tableMgr = NewTableMgr()
	//tableMgr.Load("E:\\Go\\GOPATH\\src\\github.com\\MDGSF\\MJHuPai\\Go\\sxtjmj\\genTable\\")
	//tableMgr.Load(".\\")
}

/*
CanHu 胡牌
handCards: 手牌数组，最多14张。
huType: 胡牌类型，1自摸，2点炮。
huCard: 胡的那张牌，自摸的那张 或者是 点炮的那张。
HeiSanFeng: 是否开启黑三风
ZhongFaBai: 是否开启中发白
ZhongFaWu: 是否开启中发五: 五就是五万，五条，五筒可以代替白板。
return:
true可以胡牌，最大黑三风和中发白的数量。
false不可以胡牌，0。
*/
func CanHu(handCards []int, huType int, huCard int,
	HeiSanFeng bool, ZhongFaBai bool, ZhongFaWu bool) (bool, int) {

	if !IsValidHandCards(handCards) {
		return false, 0
	}

	TableXuShuWithEye := tableMgr.TableXuShuWithEye
	TableXuShu := tableMgr.TableXuShu
	TableFeng, TableFengWithEye := getFengTable(HeiSanFeng)
	TableJian, TableJianWithEye := getJianTable(ZhongFaBai, ZhongFaWu)

	slots := GenSlots(handCards)
	allPossibleSlotHandCards := genAllPossibleSlotHandCards(slots, ZhongFaWu)

	bCanHu := false
	maxHeiSanFengNum := 0 //黑三风数量
	maxZhongFaBaiNum := 0 //中发白数量
	maxFengNum := 0       //这个是黑三风和中发白的数量总和

	for _, v := range allPossibleSlotHandCards {
		XuShu, Feng, Jian := getArray(v.slots)

		if len(XuShu) > 0 {
			ret1, fengNum1, heiSanFengNum1, zhongFaBaiNum1 := walkThroughTableXuShuJiang(XuShu, Feng, Jian,
				TableXuShuWithEye, TableXuShu, TableFeng, TableJian, huType, huCard)
			if ret1 && fengNum1 >= maxFengNum {
				bCanHu = true
				maxHeiSanFengNum = heiSanFengNum1
				maxZhongFaBaiNum = zhongFaBaiNum1
				maxFengNum = maxHeiSanFengNum + maxZhongFaBaiNum
			}
		}

		if len(Feng) > 0 {
			ret2, fengNum2, heiSanFengNum2, zhongFaBaiNum2 := walkThroughTableFengJiang(Feng, XuShu, Jian,
				TableFengWithEye, TableXuShu, TableJian, huType, huCard)
			if ret2 && fengNum2 >= maxFengNum {
				bCanHu = true
				maxHeiSanFengNum = heiSanFengNum2
				maxZhongFaBaiNum = zhongFaBaiNum2
				maxFengNum = fengNum2
			}
		}

		if len(Jian) > 0 {
			ret3, fengNum3, heiSanFengNum3, zhongFaBaiNum3 := walkThroughTableJianJiang(Jian, XuShu, Feng,
				TableJianWithEye, TableXuShu, TableFeng, huType, huCard)
			if ret3 && fengNum3 >= maxFengNum {
				bCanHu = true
				maxHeiSanFengNum = heiSanFengNum3
				maxZhongFaBaiNum = zhongFaBaiNum3
				maxFengNum = fengNum3
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

func walkThroughTableJianJiang(Jian []int, XuShu []int, Feng []int,
	TableJianWithEye *Table, TableXuShu *Table, TableFeng *Table,
	huType int, huCard int) (CanHu bool,
	FengNum int, HeiSanFengNum int, ZhongFaBaiNum int) {

	bInTable := false
	maxHeiSanFengNum := 0
	maxZhongFaBaiNum := 0

	if len(Jian) > 0 {
		if IsDragon(huCard) {
			if maxZhongFaBaiNum, bInTable = TableJianWithEye.IsValid(Jian[0], huType, huCard); !bInTable {
				return false, 0, 0, 0
			}
		} else {
			if maxZhongFaBaiNum, bInTable = TableJianWithEye.IsInTableMap(Jian[0], 0); !bInTable {
				return false, 0, 0, 0
			}
		}
	}

	for _, num := range XuShu {
		if _, ok := TableXuShu.IsInTableMap(num, 0); !ok {
			return false, 0, 0, 0
		}
	}

	if len(Feng) > 0 {
		if IsWind(huCard) {
			if maxHeiSanFengNum, bInTable = TableFeng.IsValid(Feng[0], huType, huCard); !bInTable {
				return false, 0, 0, 0
			}
		} else {
			if maxHeiSanFengNum, bInTable = TableFeng.IsInTableMap(Feng[0], 0); !bInTable {
				return false, 0, 0, 0
			}
		}
	}

	return true, maxHeiSanFengNum + maxZhongFaBaiNum, maxHeiSanFengNum, maxZhongFaBaiNum
}

func walkThroughTableFengJiang(Feng []int,
	XuShu []int, Jian []int,
	TableFengWithEye *Table,
	TableXuShu *Table, TableJian *Table,
	huType int, huCard int) (CanHu bool,
	FengNum int, HeiSanFengNum int, ZhongFaBaiNum int) {

	bInTable := false
	maxHeiSanFengNum := 0
	maxZhongFaBaiNum := 0

	if len(Feng) > 0 {
		if IsWind(huCard) {
			if maxHeiSanFengNum, bInTable = TableFengWithEye.IsValid(Feng[0], huType, huCard); !bInTable {
				return false, 0, 0, 0
			}
		} else {
			if maxHeiSanFengNum, bInTable = TableFengWithEye.IsInTableMap(Feng[0], 0); !bInTable {
				return false, 0, 0, 0
			}
		}
	}

	for _, num := range XuShu {
		if _, ok := TableXuShu.IsInTableMap(num, 0); !ok {
			return false, 0, 0, 0
		}
	}

	if len(Jian) > 0 {
		if IsDragon(huCard) {
			if maxZhongFaBaiNum, bInTable = TableJian.IsValid(Jian[0], huType, huCard); !bInTable {
				return false, 0, 0, 0
			}
		} else {
			if maxZhongFaBaiNum, bInTable = TableJian.IsInTableMap(Jian[0], 0); !bInTable {
				return false, 0, 0, 0
			}
		}
	}

	return true, maxHeiSanFengNum + maxZhongFaBaiNum, maxHeiSanFengNum, maxZhongFaBaiNum
}

func walkThroughTableXuShuJiang(XuShu []int,
	Feng []int, Jian []int,
	TableXuShuWithEye *Table, TableXuShu *Table,
	TableFeng *Table, TableJian *Table,
	huType int, huCard int) (CanHu bool,
	FengNum int, HeiSanFengNum int, ZhongFaBaiNum int) {

	bFound := false

	maxFengNum := 0
	maxHeiSanFengNum := 0
	maxZhongFaBaiNum := 0

	for i, iNum := range XuShu {

		bInTable := false
		curFengNum := 0
		curHeiSanFengNum := 0
		curZhongFaBaiNum := 0

		_, ok := TableXuShuWithEye.IsInTableMap(iNum, 0)
		if !ok {
			continue
		}

		for j, jNum := range XuShu {
			if i == j {
				continue
			}

			_, ok := TableXuShu.IsInTableMap(jNum, 0)
			if !ok {
				return false, 0, 0, 0
			}
		}

		if len(Feng) > 0 {
			if IsWind(huCard) {
				if curHeiSanFengNum, bInTable = TableFeng.IsValid(Feng[0], huType, huCard); !bInTable {
					return false, 0, 0, 0
				}
			} else {
				if curHeiSanFengNum, bInTable = TableFeng.IsInTableMap(Feng[0], 0); !bInTable {
					return false, 0, 0, 0
				}
			}
		}

		if len(Jian) > 0 {
			if IsDragon(huCard) {
				if curZhongFaBaiNum, bInTable = TableJian.IsValid(Jian[0], huType, huCard); !bInTable {
					return false, 0, 0, 0
				}
			} else {
				if curZhongFaBaiNum, bInTable = TableJian.IsInTableMap(Jian[0], 0); !bInTable {
					return false, 0, 0, 0
				}
			}
		}

		curFengNum = curHeiSanFengNum + curZhongFaBaiNum
		if curFengNum > maxFengNum {
			maxFengNum = curFengNum
			maxHeiSanFengNum = curHeiSanFengNum
			maxZhongFaBaiNum = curZhongFaBaiNum
		}

		bFound = true
	}

	if bFound {
		return true, maxFengNum, maxHeiSanFengNum, maxZhongFaBaiNum
	}
	return false, 0, 0, 0
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

func getJianTable(ZhongFaBai bool, ZhongFaWu bool) (TableJian, TableJianWithEye *Table) {
	if ZhongFaBai || ZhongFaWu {
		TableJian = tableMgr.TableJian
		TableJianWithEye = tableMgr.TableJianWithEye
	} else {
		TableJian = tableMgr.TableJianKe
		TableJianWithEye = tableMgr.TableJianKeWithEye
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
	slots [TILEMAX]int
}

func genAllPossibleSlotHandCards(slots [TILEMAX]int, ZhongFaWu bool) (allPossibleSlotHandCards []slotHandCards) {
	var a slotHandCards
	for i := 0; i < TILEMAX; i++ {
		a.slots[i] = slots[i]
	}
	allPossibleSlotHandCards = append(allPossibleSlotHandCards, a)

	if !ZhongFaWu {
		return
	}

	man5Num := slots[MAN5]
	pin5Num := slots[PIN5]
	sou5Num := slots[SOU5]
	hakNum := slots[HAK]

	for i := 0; i <= man5Num; i++ {
		for j := 0; j <= pin5Num; j++ {
			for k := 0; k <= sou5Num; k++ {
				curMan5Num := i
				curPin5Num := j
				curSou5Num := k
				curHakNum := hakNum + (man5Num - i) + (pin5Num - j) + (sou5Num - k)
				if curHakNum > 4 {
					continue
				}

				var a slotHandCards
				for i := 0; i < TILEMAX; i++ {
					a.slots[i] = slots[i]
				}
				a.slots[MAN5] = curMan5Num
				a.slots[PIN5] = curPin5Num
				a.slots[SOU5] = curSou5Num
				a.slots[HAK] = curHakNum
				allPossibleSlotHandCards = append(allPossibleSlotHandCards, a)
			}
		}
	}

	return
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
