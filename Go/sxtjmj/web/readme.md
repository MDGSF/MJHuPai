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



http://127.0.0.1:11189

请求

```
{
	"handCards":[0, 0, 0, 3, 3, 3, 6, 6, 6, 27, 28, 29, 27, 27],
	"huType":1,
	"huCard":27,
	"heiSanFeng":true,
	"zhongFaBai":true,
	"zhongFaWu":true
}
```

```
handCards: 就是手牌
huType: 1自摸胡， 2点炮胡
huCard: 就是自摸的那张牌，或者是点炮的那张牌
heiSanFeng(黑三风): true开启，false关闭
zhongFaBai(中发白): true开启，false关闭
zhongFaWu(红发五): true开启，false关闭
```

回复

```
{"hu":true,"fengNum":1}
```

```
hu: true可以胡牌， false不可以胡牌
fengNum: 可以胡牌时，这个是手牌中风的数量(包括黑三风和中发白)。不能胡牌时，这个为0。
```

