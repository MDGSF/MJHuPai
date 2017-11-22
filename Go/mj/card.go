package mj

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

//Card 麻将子定义
type Card uint8

//CharacterOne 一万
const CharacterOne = 0x01

//CharacterNine 九万
const CharacterNine = 0x09

//BambooOne 一条，幺鸡
const BambooOne = 0x11

//BambooNine 九条
const BambooNine = 0x19

//DotOne 一筒
const DotOne = 0x21

//DotNine 九筒
const DotNine = 0x29

//EastWind 东风
const EastWind = 0x31

//SouthWind 南风
const SouthWind = 0x32

//WestWind 西风
const WestWind = 0x33

//NorthWind 北风
const NorthWind = 0x34

//RedDragon 红中
const RedDragon = 0x41

//GreenDragon 发财
const GreenDragon = 0x42

//WhiteDragon 白板
const WhiteDragon = 0x43

//SpringFlower 春
const SpringFlower = 0x51

//SummerFlower 夏
const SummerFlower = 0x52

//AutumnFlower 秋
const AutumnFlower = 0x53

//WinterFlower 冬
const WinterFlower = 0x54

//PlumFlower 梅
const PlumFlower = 0x55

//OrchidFlower 兰
const OrchidFlower = 0x56

//BambooFlower 竹
const BambooFlower = 0x57

//ChrysanthemumFlower 菊
const ChrysanthemumFlower = 0x58

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

//InvalidCard 一张非真实性的牌，填充假数据之用
const InvalidCard = 0xFF

//InvalidChair 一个非真实性的玩家椅子号
const InvalidChair = 0x7F

//MJTypeWan 万, 1-9,各4张，共36张
const MJTypeWan = 0

//MJTypeTiao 条, 1-9,各4张，共36张
const MJTypeTiao = 1

//MJTypeBing 饼, 1-9,各4张，共36张
const MJTypeBing = 2

//MJTypeFeng 东南西北各4张，共16张
const MJTypeFeng = 3

//MJTypeJian 中发白  各4张，共12张
const MJTypeJian = 4

//MJTypeFlower 花
const MJTypeFlower = 5

//MJTypeNum 牌的种类数量(万，条，饼，风，箭）
const MJTypeNum = 5

//MaxCard 最大牌面值
const MaxCard = 0x43

//MaxCardArraySize 数组大小，要比最大的牌面值大一。
const MaxCardArraySize = (MaxCard + 1)

//CharacterCards 万
var CharacterCards = []Card{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x00,
}

//BambooCards 条
var BambooCards = []Card{
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x00,
}

//DotCards 筒
var DotCards = []Card{
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x00,
}

//WindCards 风
var WindCards = []Card{
	0x31, 0x32, 0x33, 0x34,
	0x31, 0x32, 0x33, 0x34,
	0x31, 0x32, 0x33, 0x34,
	0x31, 0x32, 0x33, 0x34,
	0x00,
}

//DragonCards 箭
var DragonCards = []Card{
	0x41, 0x42, 0x43,
	0x41, 0x42, 0x43,
	0x41, 0x42, 0x43,
	0x41, 0x42, 0x43,
	0x00,
}

//FlowerCards 花
var FlowerCards = []Card{
	0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x00,
}

//MahjongCards 所有牌类型
var MahjongCards = []Card{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, //万
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, //条
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, //筒
	0x31, 0x32, 0x33, 0x34, //风
	0x41, 0x42, 0x43, //箭
}

//CardNum 获取牌对应的值
func CardNum(c Card) uint8 {
	return uint8((c) & 0x0F)
}

//CardType 获取牌对应的类型
func CardType(c Card) uint8 {
	return uint8(((c) & 0xF0) >> 4)
}

//IsValidCard 判断是否是有效的牌
func IsValidCard(c Card) bool {
	if IsRank(c) || IsHonor(c) {
		return true
	}
	return false
}

//IsRank 判断是否是序数牌
func IsRank(c Card) bool {
	if IsCharacter(c) || IsBamboo(c) || IsDot(c) {
		return true
	}
	return false
}

//IsHonor 判断是否是字牌
func IsHonor(c Card) bool {
	if IsWind(c) || IsDragon(c) {
		return true
	}
	return false
}

//IsCharacter 判断是否是万牌
func IsCharacter(c Card) bool {
	return c >= CharacterOne && c <= CharacterNine
}

//IsBamboo 判断是否是条牌
func IsBamboo(c Card) bool {
	return c >= BambooOne && c <= BambooNine
}

//IsDot 判断是否是筒牌
func IsDot(c Card) bool {
	return c >= DotOne && c <= DotNine
}

//IsWind 判断是否是风牌
func IsWind(c Card) bool {
	return c >= EastWind && c <= NorthWind
}

//IsDragon 判断是否是箭牌
func IsDragon(c Card) bool {
	return c >= RedDragon && c <= WhiteDragon
}
