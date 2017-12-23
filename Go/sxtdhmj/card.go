package sxtdhmj

/*
常用中英文对照(http://www.xqbase.com/other/mahjongg_english.htm):
万牌 Character
条牌 Bamboo
筒牌 Dot
风牌 Wind
箭牌 Dragon
花牌 Flower

序数牌	Rank
字牌 Honor

东、南、西、北	East / South / West / North Wind
中、发、白	Red / Green / White Dragon

春、夏、秋、冬	Spring / Summer / Autumn / Winter
梅、兰、竹、菊	Plum / Orchid / Bamboo / Chrysanthemum
*/

// 麻将牌ID定义
const (
	MAN     = 0
	MAN1    = 0
	MAN2    = 1
	MAN3    = 2
	MAN4    = 3
	MAN5    = 4
	MAN6    = 5
	MAN7    = 6
	MAN8    = 7
	MAN9    = 8
	PIN     = 9 // 筒子
	PIN1    = 9
	PIN2    = 10
	PIN3    = 11
	PIN4    = 12
	PIN5    = 13
	PIN6    = 14
	PIN7    = 15
	PIN8    = 16
	PIN9    = 17
	SOU     = 18 // 索子
	SOU1    = 18
	SOU2    = 19
	SOU3    = 20
	SOU4    = 21
	SOU5    = 22
	SOU6    = 23
	SOU7    = 24
	SOU8    = 25
	SOU9    = 26
	SuitMax = SOU9
	TON     = 27 // 东
	NAN     = 28 // 南
	SHA     = 29 // 西
	PEI     = 30 // 北
	HAK     = 31 // Haku, 白
	HAT     = 32 // Hatsu, 发
	CHU     = 33 // Chun 中
	TILEMAX = 34
)

//CharacterOne 一万
const CharacterOne = 0

//CharacterNine 九万
const CharacterNine = 8

//BambooOne 一条，幺鸡
const BambooOne = 18

//BambooNine 九条
const BambooNine = 26

//DotOne 一筒
const DotOne = 9

//DotNine 九筒
const DotNine = 17

//EastWind 东风
const EastWind = 27

//SouthWind 南风
const SouthWind = 28

//WestWind 西风
const WestWind = 29

//NorthWind 北风
const NorthWind = 30

//RedDragon 红中
const RedDragon = 33

//GreenDragon 发财
const GreenDragon = 32

//WhiteDragon 白板
const WhiteDragon = 31

//SpringFlower 春
const SpringFlower = 38

//SummerFlower 夏
const SummerFlower = 39

//AutumnFlower 秋
const AutumnFlower = 40

//WinterFlower 冬
const WinterFlower = 41

//PlumFlower 梅
const PlumFlower = 34

//OrchidFlower 兰
const OrchidFlower = 35

//BambooFlower 竹
const BambooFlower = 36

//ChrysanthemumFlower 菊
const ChrysanthemumFlower = 37

//MaxGamePlayerUser 最大游戏玩家数
const MaxGamePlayerUser = 4

//AllCardsNum 麻将中牌数(36张万，36张条，36张筒，16张风，12张箭，一共136张)
const AllCardsNum = 136

//CardsTypeNum 所有类型牌的总数(9张万，9张条，9张筒，4张风，3张箭，一共34张)
const CardsTypeNum = 34

//MaxHandCardNum 最多的手牌数量为14张
const MaxHandCardNum = 14

//MaxFixCardNum 最多的固定牌数量
const MaxFixCardNum = 4

//MaxRankNum 序数牌的张数: 1~9
const MaxRankNum = 9

//MaxWindNum 风牌的张数: 东南西北
const MaxWindNum = 4

//MaxDragonNum 箭牌的张数: 中发白
const MaxDragonNum = 3

//MaxGangCount 最大杠操作,因为手牌最多14张，一个杠要用4张牌，所以手中最多同时存在3个杠。
const MaxGangCount = 3

//MJTypeNum 牌的种类数量(万，条，饼，风，箭）
const MJTypeNum = 5

//IsValidCard 判断是否是有效的牌
func IsValidCard(c int) bool {
	if IsSuit(c) || IsHonor(c) {
		return true
	}
	return false
}

// IsSuit 是否数牌：万、条、筒
func IsSuit(tileID int) bool {
	return tileID >= MAN && tileID <= SOU9
}

//IsCharacter 判断是否是万牌
func IsCharacter(c int) bool {
	return c >= MAN && c <= MAN9
}

//IsBamboo 判断是否是条牌
func IsBamboo(c int) bool {
	return c >= SOU && c <= SOU9
}

//IsDot 判断是否是筒牌
func IsDot(c int) bool {
	return c >= PIN && c <= PIN9
}

// IsHonor 是否字牌：东、南、西、北、中、发、白
func IsHonor(tileID int) bool {
	return tileID >= TON && tileID <= CHU
}

// IsWind 是否风牌：东、南、西、北
func IsWind(tileID int) bool {
	return tileID >= TON && tileID <= PEI
}

// IsDragon 是否箭牌：中、发、白
func IsDragon(tileID int) bool {
	return tileID >= HAK && tileID <= CHU
}
