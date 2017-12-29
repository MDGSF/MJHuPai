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

	// TableJian 箭牌组成刻子和中发白不带将
	TableJian *Table

	// TableJianWithEye 箭牌组成刻子和中发白带将
	TableJianWithEye *Table

	// TableZi 字牌不带将
	TableZi *Table

	// TableZiWithEye 字牌带将
	TableZiWithEye *Table
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
	tableMgr.TableJian = NewTable()
	tableMgr.TableJianWithEye = NewTable()

	tableMgr.TableZi = NewTable()
	tableMgr.TableZiWithEye = NewTable()

	return tableMgr
}

//Load 加载所有的表到内存中。
func (tableMgr *TableMgr) Load(directory string) {

	//序数牌
	tableMgr.TableXuShu.Load(directory + "sxtdh_tbl/TableXuShu")
	tableMgr.TableXuShuWithEye.Load(directory + "sxtdh_tbl/TableXuShuWithEye")

	//风牌
	tableMgr.TableFengKe.Load(directory + "sxtdh_tbl/TableFengKe")
	tableMgr.TableFengKeWithEye.Load(directory + "sxtdh_tbl/TableFengKeWithEye")
	tableMgr.TableFeng.Load(directory + "sxtdh_tbl/TableFeng")
	tableMgr.TableFengWithEye.Load(directory + "sxtdh_tbl/TableFengWithEye")

	//箭牌
	tableMgr.TableJianKe.Load(directory + "sxtdh_tbl/TableJianKe")
	tableMgr.TableJianKeWithEye.Load(directory + "sxtdh_tbl/TableJianKeWithEye")
	tableMgr.TableJian.Load(directory + "sxtdh_tbl/TableJian")
	tableMgr.TableJianWithEye.Load(directory + "sxtdh_tbl/TableJianWithEye")

	tableMgr.TableZi.Load(directory + "sxtdh_tbl/TableZi")
	tableMgr.TableZiWithEye.Load(directory + "sxtdh_tbl/TableZiWithEye")
}

//Dump 固化所有内存中的表
func (tableMgr *TableMgr) Dump() {

	//序数牌
	tableMgr.TableXuShu.Dump("sxtdh_tbl/TableXuShu")
	tableMgr.TableXuShuWithEye.Dump("sxtdh_tbl/TableXuShuWithEye")

	//风牌
	tableMgr.TableFengKe.Dump("sxtdh_tbl/TableFengKe")
	tableMgr.TableFengKeWithEye.Dump("sxtdh_tbl/TableFengKeWithEye")
	tableMgr.TableFeng.Dump("sxtdh_tbl/TableFeng")
	tableMgr.TableFengWithEye.Dump("sxtdh_tbl/TableFengWithEye")

	//箭牌
	tableMgr.TableJianKe.Dump("sxtdh_tbl/TableJianKe")
	tableMgr.TableJianKeWithEye.Dump("sxtdh_tbl/TableJianKeWithEye")
	tableMgr.TableJian.Dump("sxtdh_tbl/TableJian")
	tableMgr.TableJianWithEye.Dump("sxtdh_tbl/TableJianWithEye")

	tableMgr.TableZi.Dump("sxtdh_tbl/TableZi")
	tableMgr.TableZiWithEye.Dump("sxtdh_tbl/TableZiWithEye")
}
