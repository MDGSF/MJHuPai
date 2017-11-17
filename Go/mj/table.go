package mj

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const LaiZiNum = 9

type Table struct {
	Map [LaiZiNum]*map[int]bool
}

func NewTable() *Table {
	table := &Table{}
	for i := 0; i < LaiZiNum; i++ {
		table.Map[i] = &map[int]bool{}
	}
	return table
}

func (table *Table) IsInTable(num int) bool {
	for i := 0; i < LaiZiNum; i++ {
		if table.IsInTableMap(num, i) {
			return true
		}
	}
	return false
}

func (table *Table) IsInTableMap(num int, iLaiZiNum int) bool {
	_, ok := (*table.Map[iLaiZiNum])[num]
	return ok
}

func (table *Table) Load(prefix string) {
	for i := 0; i < LaiZiNum; i++ {
		name := prefix + "_" + strconv.Itoa(i) + ".tbl"
		loadFromFile(name, table.Map[i])
	}
}

func (table *Table) Dump(prefix string) {
	for i := 0; i < LaiZiNum; i++ {
		name := prefix + "_" + strconv.Itoa(i) + ".tbl"
		dumpToFile(name, table.Map[i])
	}
}

func loadFromFile(name string, table *map[int]bool) {
	file, _ := os.Open(name)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(buf)
		key, _ := strconv.Atoi(str)
		(*table)[key] = true
	}
}

func dumpToFile(name string, table *map[int]bool) {
	file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	buf := bufio.NewWriter(file)
	for key := range *table {
		fmt.Fprintf(buf, "%d\n", key)
	}
	buf.Flush()
}
