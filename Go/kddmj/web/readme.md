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



http://127.0.0.1:11188

请求

```
{
	"handCards":[0, 0, 1, 1, 1],
	"laizi":[1]
}
```

回复

```
{"hu":true,"dianshu":10}
```


