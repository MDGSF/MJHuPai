package mj

import (
	"fmt"
)

var tableMgr *TableMgr

func init() {
	tableMgr = NewTableMgr()
	tableMgr.Load("E:\\Go\\GOPATH\\src\\github.com\\MDGSF\\MJHuPai\\Go\\mj\\genTable\\")
	initM1()
}

/*
CanHuWithLaiZi 赖子胡牌
handCards: 手牌数组，最多14张。
laizi: 赖子数组，保存可能的赖子。
return: true可以胡牌，false不可以胡牌。
*/
func CanHuWithLaiZi(handCards []Card, laizi []Card) bool {

	if !IsValidHandCards(handCards) {
		return false
	}

	laiziNum := 0
	var slots [MaxCardArraySize]Card
	for _, c := range handCards {
		if isLaizi(c, laizi) {
			laiziNum++
		} else {
			slots[c]++
		}
	}

	var XuShu []int
	if getWan(slots) > 0 {
		XuShu = append(XuShu, getWan(slots))
	}
	if getTiao(slots) > 0 {
		XuShu = append(XuShu, getTiao(slots))
	}
	if getTong(slots) > 0 {
		XuShu = append(XuShu, getTong(slots))
	}

	var Zi []int
	if getFeng(slots) > 0 {
		Zi = append(Zi, getFeng(slots))
	}
	if getJian(slots) > 0 {
		Zi = append(Zi, getJian(slots))
	}

	for i, iNum := range XuShu {
		success := true
		hasLaiZiNum := laiziNum
		needLaiZiNum, ok := tableMgr.TableXuShuWithEye.IsInTable(iNum)
		if !ok || needLaiZiNum > hasLaiZiNum {
			continue
		}
		hasLaiZiNum -= needLaiZiNum

		for j, jNum := range XuShu {
			if i == j {
				continue
			}
			needLaiZiNum, ok := tableMgr.TableXuShu.IsInTable(jNum)
			if !ok || needLaiZiNum > hasLaiZiNum {
				success = false
				break
			}
			hasLaiZiNum -= needLaiZiNum
		}
		if !success {
			continue
		}

		for _, num := range Zi {
			needLaiZiNum, ok := tableMgr.TableZi.IsInTable(num)
			if !ok || needLaiZiNum > hasLaiZiNum {
				success = false
				break
			}
		}
		if !success {
			continue
		}

		return true
	}

	for i, iNum := range Zi {
		success := true
		hasLaiZiNum := laiziNum
		needLaiZiNum, ok := tableMgr.TableZiWithEye.IsInTable(iNum)
		if !ok || needLaiZiNum > hasLaiZiNum {
			continue
		}
		hasLaiZiNum -= needLaiZiNum

		for j, jNum := range Zi {
			if i == j {
				continue
			}
			needLaiZiNum, ok := tableMgr.TableZi.IsInTable(jNum)
			if !ok || needLaiZiNum > hasLaiZiNum {
				success = false
				break
			}
			hasLaiZiNum -= needLaiZiNum
		}
		if !success {
			continue
		}

		for _, num := range XuShu {
			needLaiZiNum, ok := tableMgr.TableXuShu.IsInTable(num)
			if !ok || needLaiZiNum > hasLaiZiNum {
				success = false
				break
			}
		}
		if !success {
			continue
		}

		return true
	}

	return false
}

func isLaizi(c Card, LaiZiArr []Card) bool {
	for _, laizi := range LaiZiArr {
		if c == laizi {
			return true
		}
	}
	return false
}

/*
CanHu 胡牌
handCards: 手牌数组，最多14张。
return: true可以胡牌，false不可以胡牌。
*/
func CanHu(handCards []Card) bool {

	if !IsValidHandCards(handCards) {
		return false
	}

	slots := GenSlots(handCards)

	// fmt.Println("getWan = ", getWan(slots))
	// fmt.Println("getTiao = ", getTiao(slots))
	// fmt.Println("getTong = ", getTong(slots))
	// fmt.Println("getFeng = ", getFeng(slots))
	// fmt.Println("getJian = ", getJian(slots))

	wan := getWan(slots)
	tiao := getTiao(slots)
	tong := getTong(slots)
	feng := getFeng(slots)
	jian := getJian(slots)

	var XuShu []int
	if wan > 0 {
		XuShu = append(XuShu, wan)
	}
	if tiao > 0 {
		XuShu = append(XuShu, tiao)
	}
	if tong > 0 {
		XuShu = append(XuShu, tong)
	}

	var Zi []int
	if feng > 0 {
		Zi = append(Zi, feng)
	}
	if jian > 0 {
		Zi = append(Zi, jian)
	}

	for i, iNum := range XuShu {
		if !tableMgr.TableXuShuWithEye.IsInTableMap(iNum, 0) {
			continue
		}

		for j, jNum := range XuShu {
			if i == j {
				continue
			}

			if !tableMgr.TableXuShu.IsInTableMap(jNum, 0) {
				return false
			}
		}

		for _, num := range Zi {
			if !tableMgr.TableZi.IsInTableMap(num, 0) {
				return false
			}
		}

		return true
	}

	for i, iNum := range Zi {
		if !tableMgr.TableZiWithEye.IsInTableMap(iNum, 0) {
			continue
		}

		for j, jNum := range Zi {
			if i == j {
				continue
			}

			if !tableMgr.TableZi.IsInTableMap(jNum, 0) {
				return false
			}
		}

		for _, num := range XuShu {
			if !tableMgr.TableXuShu.IsInTableMap(num, 0) {
				return false
			}
		}

		return true
	}

	return false
}

/*
IsValidHandCards 判断手牌是否合法
handCards: 手牌数组，最多14张。
return: true手牌合法，false手牌非法。
*/
func IsValidHandCards(handCards []Card) bool {
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
handCards []Card = {0x01, 0x01, 0x11, 0x12, 0x13, 0x33, 0x43, 0x43}
[MaxCardArraySize]Card = {
	2, 0, 0, 0, 0, 0, 0, 0, 0, //万
	1, 1, 1, 0, 0, 0, 0, 0, 0, //条
	0, 0, 0, 0, 0, 0, 0, 0, 0, //筒
	0, 0, 1, 0, //风
	0, 0, 2, //箭
}
*/
func GenSlots(handCards []Card) [MaxCardArraySize]Card {
	var slots [MaxCardArraySize]Card

	for _, c := range handCards {
		slots[c]++
	}

	return slots
}

func getWan(slots [MaxCardArraySize]Card) int {
	return getNum(slots, CharacterOne, CharacterNine)
}

func getTiao(slots [MaxCardArraySize]Card) int {
	return getNum(slots, BambooOne, BambooNine)
}

func getTong(slots [MaxCardArraySize]Card) int {
	return getNum(slots, DotOne, DotNine)
}

func getFeng(slots [MaxCardArraySize]Card) int {
	return getNum(slots, EastWind, NorthWind)
}

func getJian(slots [MaxCardArraySize]Card) int {
	return getNum(slots, RedDragon, WhiteDragon)
}

func getNum(slots [MaxCardArraySize]Card, iStart int, iEnd int) int {
	num := 0
	for i := iStart; i <= iEnd; i++ {
		num = num*10 + int(slots[i])
	}
	return num
}

/*
ShowHandCards 打印手牌
*/
func ShowHandCards(handCards []Card) {
	for _, c := range handCards {
		fmt.Printf("0x%02X, ", c)
	}
	fmt.Println()
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

var M1 map[int]Card

func initM1() {

	M1 = make(map[int]Card)

	M1[0] = Card(0x01)
	M1[1] = Card(0x02)
	M1[2] = Card(0x03)
	M1[3] = Card(0x04)
	M1[4] = Card(0x05)
	M1[5] = Card(0x06)
	M1[6] = Card(0x07)
	M1[7] = Card(0x08)
	M1[8] = Card(0x09)

	M1[9] = Card(0x11)
	M1[10] = Card(0x12)
	M1[11] = Card(0x13)
	M1[12] = Card(0x14)
	M1[13] = Card(0x15)
	M1[14] = Card(0x16)
	M1[15] = Card(0x17)
	M1[16] = Card(0x18)
	M1[17] = Card(0x19)

	M1[18] = Card(0x21)
	M1[19] = Card(0x22)
	M1[20] = Card(0x23)
	M1[21] = Card(0x24)
	M1[22] = Card(0x25)
	M1[23] = Card(0x26)
	M1[24] = Card(0x27)
	M1[25] = Card(0x28)
	M1[26] = Card(0x29)

	M1[27] = Card(0x31)
	M1[28] = Card(0x32)
	M1[29] = Card(0x33)
	M1[30] = Card(0x34)

	M1[31] = Card(0x41)
	M1[32] = Card(0x42)
	M1[33] = Card(0x43)

	M1[34] = Card(0x51)
	M1[35] = Card(0x52)
	M1[36] = Card(0x53)
	M1[37] = Card(0x54)
	M1[38] = Card(0x55)
	M1[39] = Card(0x56)
	M1[40] = Card(0x57)
	M1[41] = Card(0x58)
}

//IsWinableTest test
func IsWinableTest(hai []int) bool {

	if len(hai) > 14 {
		return false
	}

	handCards := make([]Card, 0)

	for _, v := range hai {
		handCards = append(handCards, M1[v])
	}

	return CanHu(handCards)
}
