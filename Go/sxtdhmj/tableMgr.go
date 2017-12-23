package sxtdhmj

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

	tableMgr.TableXuShu.Load(directory + "sxtdh_tbl/tableXuShu")
	tableMgr.TableXuShuWithEye.Load(directory + "sxtdh_tbl/tableEyeXuShu")
	tableMgr.TableZi.Load(directory + "sxtdh_tbl/tableZi")
	tableMgr.TableZiWithEye.Load(directory + "sxtdh_tbl/tableEyeZi")
}

//Dump 固化所有内存中的表
func (tableMgr *TableMgr) Dump() {
	tableMgr.TableXuShu.Dump("sxtdh_tbl/tableXuShu")
	tableMgr.TableXuShuWithEye.Dump("sxtdh_tbl/tableEyeXuShu")
	tableMgr.TableZi.Dump("sxtdh_tbl/tableZi")
	tableMgr.TableZiWithEye.Dump("sxtdh_tbl/tableEyeZi")
}
