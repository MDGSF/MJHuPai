package kddmj

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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
	Map [LaiZiNum]*map[int]int
}

//NewTable 新建一张表
func NewTable() *Table {
	table := &Table{}
	for i := 0; i < LaiZiNum; i++ {
		table.Map[i] = &map[int]int{}
	}
	return table
}

//IsInTable 判断num是否在这个表中
func (table *Table) IsInTable(num int) (int, int, bool) {
	for i := 0; i < LaiZiNum; i++ {
		value, ok := table.IsInTableMap(num, i)
		if ok {
			return value, i, true
		}
	}
	return 0, 0, false
}

//IsInTableMap 判断num是不是在有iLaiZiNum个赖子的那个map中
func (table *Table) IsInTableMap(num int, iLaiZiNum int) (int, bool) {
	value, ok := (*table.Map[iLaiZiNum])[num]
	return value, ok
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

func loadFromFile(name string, table *map[int]int) {
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
		value, _ := strconv.Atoi(result[1])
		(*table)[key] = value
	}
}

func dumpToFile(name string, table *map[int]int) {
	file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	buf := bufio.NewWriter(file)
	for key, value := range *table {
		fmt.Fprintf(buf, "%d=%d\n", key, value)
	}
	buf.Flush()
}
