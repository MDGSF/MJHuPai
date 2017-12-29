package sxtjmj

//TableMgr 表管理工具
type TableMgr struct {

	// TableXuShu 序数牌不带将
	TableXuShu *Table

	// TableXuShuWithEye 序数牌带将
	TableXuShuWithEye *Table

	// TableZi 风牌只组成刻子不带将
	TableFengKe *Table

	// TableZi 风牌只组成刻子带将
	TableFengKeWithEye *Table

	// TableFeng 风牌组成刻子和黑三风不带将
	TableFeng *Table

	// TableFengWithEye 风牌组成刻子和黑三风带将
	TableFengWithEye *Table

	// TableJianKe 箭牌只组成刻子不带将
	TableJianKe *Table

	// TableJianKeWithEye 箭牌只组成刻子带将
	TableJianKeWithEye *Table
}

//NewTableMgr 新建一个表管理工具
func NewTableMgr() *TableMgr {
	tableMgr := &TableMgr{}

	//序数牌
	tableMgr.TableXuShu = NewTable()
	tableMgr.TableXuShuWithEye = NewTable()

	//风牌
	tableMgr.TableFengKe = NewTable()
	tableMgr.TableFengKeWithEye = NewTable()
	tableMgr.TableFeng = NewTable()
	tableMgr.TableFengWithEye = NewTable()

	//箭牌
	tableMgr.TableJianKe = NewTable()
	tableMgr.TableJianKeWithEye = NewTable()

	return tableMgr
}

//Load 加载所有的表到内存中。
func (tableMgr *TableMgr) Load(directory string) {

	//序数牌
	tableMgr.TableXuShu.Load(directory + "sxtj_tbl/TableXuShu")
	tableMgr.TableXuShuWithEye.Load(directory + "sxtj_tbl/TableXuShuWithEye")

	//风牌
	tableMgr.TableFengKe.Load(directory + "sxtj_tbl/TableFengKe")
	tableMgr.TableFengKeWithEye.Load(directory + "sxtj_tbl/TableFengKeWithEye")
	tableMgr.TableFeng.Load(directory + "sxtj_tbl/TableFeng")
	tableMgr.TableFengWithEye.Load(directory + "sxtj_tbl/TableFengWithEye")

	//箭牌
	tableMgr.TableJianKe.Load(directory + "sxtj_tbl/TableJianKe")
	tableMgr.TableJianKeWithEye.Load(directory + "sxtj_tbl/TableJianKeWithEye")
}

//Dump 固化所有内存中的表
func (tableMgr *TableMgr) Dump() {

	//序数牌
	tableMgr.TableXuShu.Dump("sxtj_tbl/TableXuShu")
	tableMgr.TableXuShuWithEye.Dump("sxtj_tbl/TableXuShuWithEye")

	//风牌
	tableMgr.TableFengKe.Dump("sxtj_tbl/TableFengKe")
	tableMgr.TableFengKeWithEye.Dump("sxtj_tbl/TableFengKeWithEye")
	tableMgr.TableFeng.Dump("sxtj_tbl/TableFeng")
	tableMgr.TableFengWithEye.Dump("sxtj_tbl/TableFengWithEye")

	//箭牌
	tableMgr.TableJianKe.Dump("sxtj_tbl/TableJianKe")
	tableMgr.TableJianKeWithEye.Dump("sxtj_tbl/TableJianKeWithEye")
}
