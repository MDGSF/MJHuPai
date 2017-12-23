package kddmj

//TableMgr 表管理工具
type TableMgr struct {
	TableXuShu        *Table
	TableXuShuWithEye *Table
	TableZi           *Table
	TableZiWithEye    *Table
}

//NewTableMgr 新建一个表管理工具
func NewTableMgr() *TableMgr {
	tableMgr := &TableMgr{}
	tableMgr.TableXuShu = NewTable()
	tableMgr.TableXuShuWithEye = NewTable()
	tableMgr.TableZi = NewTable()
	tableMgr.TableZiWithEye = NewTable()
	return tableMgr
}

//Load 加载所有的表到内存中。
func (tableMgr *TableMgr) Load(directory string) {

	tableMgr.TableXuShu.Load(directory + "tbl/tableXuShu")
	tableMgr.TableXuShuWithEye.Load(directory + "tbl/tableEyeXuShu")
	tableMgr.TableZi.Load(directory + "tbl/tableZi")
	tableMgr.TableZiWithEye.Load(directory + "tbl/tableEyeZi")
}

//Dump 固化所有内存中的表
func (tableMgr *TableMgr) Dump() {
	tableMgr.TableXuShu.Dump("tbl/tableXuShu")
	tableMgr.TableXuShuWithEye.Dump("tbl/tableEyeXuShu")
	tableMgr.TableZi.Dump("tbl/tableZi")
	tableMgr.TableZiWithEye.Dump("tbl/tableEyeZi")
}
