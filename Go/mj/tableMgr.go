package mj

type TableMgr struct {
	TableXuShu        *Table
	TableXuShuWithEye *Table
	TableZi           *Table
	TableZiWithEye    *Table
}

func NewTableMgr() *TableMgr {
	tableMgr := &TableMgr{}
	tableMgr.TableXuShu = NewTable()
	tableMgr.TableXuShuWithEye = NewTable()
	tableMgr.TableZi = NewTable()
	tableMgr.TableZiWithEye = NewTable()
	return tableMgr
}

func (tableMgr *TableMgr) Load(directory string) {

	tableMgr.TableXuShu.Load(directory + "tbl/tableXuShu")
	tableMgr.TableXuShuWithEye.Load(directory + "tbl/tableEyeXuShu")
	tableMgr.TableZi.Load(directory + "tbl/tableZi")
	tableMgr.TableZiWithEye.Load(directory + "tbl/tableEyeZi")
}

func (tableMgr *TableMgr) Dump() {
	tableMgr.TableXuShu.Dump("tbl/tableXuShu")
	tableMgr.TableXuShuWithEye.Dump("tbl/tableEyeXuShu")
	tableMgr.TableZi.Dump("tbl/tableZi")
	tableMgr.TableZiWithEye.Dump("tbl/tableEyeZi")
}
