package sxtdhmj

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type PuZi struct {
	PuZi [3]int
}

type HuNode struct {
	Card    int //胡的那张牌
	FengNum int //胡这张牌的时候，带有的黑三风或中发白的数量
}

//TdhValue 生成的表的value定义
type TdhValue struct {
	FengNum  int   //保存黑三风、中发白的数量
	ZiMoList []int //保存只能自摸的牌，胡黑三风和中发白时，只能自摸。

	HuZiMo    map[int]int //保存所有可以自摸胡的牌，和对应的风的数量 <key: , vlaue: >
	HuDianPao map[int]int //保存所有可以点炮胡的牌，和对应的风的数量 <key: , vlaue: >
}

//NewTable 新建TdhValue
func NewTdhValue() *TdhValue {
	value := &TdhValue{}
	value.HuZiMo = make(map[int]int)
	value.HuDianPao = make(map[int]int)
	return value
}

//LaiZiNum 赖子的数量
const LaiZiNum = 9

/*
Table 表
每张表中都有LaiZiNum个map，
map[0],map[1],...,map[LaiZiNum-1]
分别对应着赖子的数量，即
map[0]中保存着当赖子的数量为0的时候，可以胡牌的所有可能性。
map[1]中保存着当赖子的数量为1的时候，可以胡牌的所有可能性。
...
*/
type Table struct {
	Map [LaiZiNum]*map[int]*TdhValue
}

//NewTable 新建一张表
func NewTable() *Table {
	table := &Table{}
	for i := 0; i < LaiZiNum; i++ {
		table.Map[i] = &map[int]*TdhValue{}
	}
	return table
}

//IsInTable 判断num是否在这个表中
func (table *Table) IsInTable(num int) (*TdhValue, int, bool) {
	for i := 0; i < LaiZiNum; i++ {
		value, ok := table.IsInTableMap(num, i)
		if ok {
			return value, i, true
		}
	}
	return nil, 0, false
}

//IsInTableMap 判断num是不是在有iLaiZiNum个赖子的那个map中
func (table *Table) IsInTableMap(num int, iLaiZiNum int) (*TdhValue, bool) {
	value, ok := (*table.Map[iLaiZiNum])[num]
	return value, ok
}

/*
IsValid 判断是不是可以胡牌
num:
huType: 胡牌类型，1自摸，2点炮。
huCard: 胡的那张牌，自摸的那张 或者是 点炮的那张。
*/
func (table *Table) IsValid(num int, huType int, huCard int) (bInTable bool, maxFengNum int) {

	//因为推倒胡没有赖子，所以这里直接就查第一张表就好了。
	value, ok := (*table.Map[0])[num]
	if !ok {
		return false, 0
	}

	//不是黑三风和中发白
	if len(value.HuZiMo) == 0 && len(value.HuDianPao) == 0 {
		return true, 0
	}

	if huType == 1 {
		if v, ok := value.HuZiMo[huCard]; ok {
			return true, v
		}
	} else if huType == 2 {
		if v, ok := value.HuDianPao[huCard]; ok {
			return true, v
		}
	}

	return false, 0
}

//Load 加载表到内存中
func (table *Table) Load(prefix string) {
	for i := 0; i < LaiZiNum; i++ {
		name := prefix + "_" + strconv.Itoa(i) + ".tbl"
		loadFromFile(name, table.Map[i])
	}
}

//Dump 固化内存中的表。
func (table *Table) Dump(prefix string) {
	for i := 0; i < LaiZiNum; i++ {
		name := prefix + "_" + strconv.Itoa(i) + ".tbl"
		dumpToFile(name, table.Map[i])
	}
}

func loadFromFile(name string, table *map[int]*TdhValue) {
	file, _ := os.Open(name)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(buf)
		result := strings.Split(str, "=")
		key, _ := strconv.Atoi(result[0])
		sValue := strings.Split(result[1], "|")

		tdh := NewTdhValue()
		tdh.FengNum, _ = strconv.Atoi(sValue[0])

		ziMoValue := sValue[1]

		dianPaoValue := sValue[2]

		if len(ziMoValue) > 0 {
			sub := strings.Split(ziMoValue, ",")
			if len(sub) > 0 {
				for i := 0; i < len(sub); i += 2 {
					node := &HuNode{}
					node.Card, _ = strconv.Atoi(sub[i])
					node.FengNum, _ = strconv.Atoi(sub[i+1])
					tdh.HuZiMo[node.Card] = node.FengNum
				}
			}
		}

		if len(dianPaoValue) > 0 {
			sub := strings.Split(dianPaoValue, ",")
			if len(sub) > 0 {
				for i := 0; i < len(sub); i += 2 {
					node := &HuNode{}
					node.Card, _ = strconv.Atoi(sub[i])
					node.FengNum, _ = strconv.Atoi(sub[i+1])
					tdh.HuDianPao[node.Card] = node.FengNum
				}
			}
		}

		(*table)[key] = tdh
	}
}

func dumpToFile(name string, table *map[int]*TdhValue) {
	file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	buf := bufio.NewWriter(file)
	for key, value := range *table {
		//fmt.Fprintf(buf, "%d=%d\n", key, value)
		fmt.Fprintf(buf, "%d=", key)
		fmt.Fprintf(buf, "%d", value.FengNum)
		fmt.Fprintf(buf, "|")

		bFirst := true
		for i, v := range value.HuZiMo {
			if bFirst {
				bFirst = false
				fmt.Fprintf(buf, "%d,%d", i, v)
			} else {
				fmt.Fprintf(buf, ",%d,%d", i, v)
			}
		}
		fmt.Fprintf(buf, "|")

		bFirst = true
		for i, v := range value.HuDianPao {
			if bFirst {
				bFirst = false
				fmt.Fprintf(buf, "%d,%d", i, v)
			} else {
				fmt.Fprintf(buf, ",%d,%d", i, v)
			}
		}

		fmt.Fprintf(buf, "\n")
	}
	buf.Flush()
}
