package kddmahjong

import "log"

var (
	//OneDragon 一条龙
	OneDragon *Table

	//OneDragonNoLaiZi 一条龙没有赖子的那张表
	OneDragonNoLaiZi *map[int]int
)

func init() {
	OneDragon = NewTable()
	OneDragonNoLaiZi = OneDragon.Map[0]

	genOneDragonTable()
}

func genOneDragonTable() {

	log.Println("Start generate oneDragon table.")

	//一条龙有9张牌，14-9=5，那就遍历剩下的5张牌。

	//第一种情况，就是只有9张牌
	handCards := []int{}
	for i := 0; i <= 8; i++ {
		handCards = append(handCards, i)
	}
	key := oneDragonCalcKey(handCards)
	(*OneDragonNoLaiZi)[key] = 1
	addToOneDragonSub(handCards, 1)

	//第二种情况，9张牌 + 2张将牌
	for i := 0; i <= 8; i++ {
		handCards := []int{}
		for m := 0; m <= 8; m++ { //先把一条龙加入手牌中
			handCards = append(handCards, m)
		}

		handCards = append(handCards, i)
		handCards = append(handCards, i)

		key = oneDragonCalcKey(handCards)
		(*OneDragonNoLaiZi)[key] = 1
		addToOneDragonSub(handCards, 1)
	}

	//第三种情况，9张牌 + 3张牌
	onePuZiChan := make(chan oneDragonPuZi)
	go genPuZi(onePuZiChan)
	for one := range onePuZiChan {
		handCards := []int{}
		for m := 0; m <= 8; m++ { //先把一条龙加入手牌中
			handCards = append(handCards, m)
		}

		var handCards1 = handCards
		handCards1 = append(handCards1, one.PuZi[0])
		handCards1 = append(handCards1, one.PuZi[1])
		handCards1 = append(handCards1, one.PuZi[2])

		key = oneDragonCalcKey(handCards1)
		(*OneDragonNoLaiZi)[key] = 1
		addToOneDragonSub(handCards1, 1)
	}

	//第四种情况，9张牌 + 2张将牌 + 3张牌
	for i := 0; i <= 8; i++ {
		handCards := []int{}
		for m := 0; m <= 8; m++ { //先把一条龙加入手牌中
			handCards = append(handCards, m)
		}

		handCards = append(handCards, i)
		handCards = append(handCards, i)

		onePuZiChan := make(chan oneDragonPuZi)
		go genPuZi(onePuZiChan)
		for one := range onePuZiChan {

			var handCards1 = handCards
			handCards1 = append(handCards1, one.PuZi[0])
			handCards1 = append(handCards1, one.PuZi[1])
			handCards1 = append(handCards1, one.PuZi[2])
			if !IsValidHandCards(handCards1) {
				continue
			}

			key = oneDragonCalcKey(handCards1)
			(*OneDragonNoLaiZi)[key] = 1
			addToOneDragonSub(handCards1, 1)
		}
	}

	count := 0
	for k := range *OneDragonNoLaiZi {
		count++
		log.Println(k)
	}
	log.Println("count = ", count)

}

func addToOneDragonSub(cardsTemp []int, iLaiZiNum int) {
	if iLaiZiNum >= LaiZiNum {
		return
	}

	cards := []int{}
	cards = append(cards, cardsTemp...)

	for i := 0; i <= 8; i++ {
		if cards[i] == 0 {
			continue
		}

		cards[i]--
		key := oneDragonCalcKey(cards)
		if _, ok := (*OneDragon.Map[iLaiZiNum])[key]; ok {
			continue
		}
		(*OneDragon.Map[iLaiZiNum])[key] = 1

		addToOneDragonSub(cards, iLaiZiNum+1)
		cards[i]++
	}
}

type oneDragonPuZi struct {
	PuZi [3]int
}

func genPuZi(puziChan chan oneDragonPuZi) {
	for i := 0; i <= 8; i++ {
		if !IsValidCard(int(i)) {
			continue
		}

		var pu oneDragonPuZi
		pu.PuZi[0] = int(i)
		pu.PuZi[1] = int(i)
		pu.PuZi[2] = int(i)
		puziChan <- pu
	}

	for i := 0; i <= 8; i++ {
		if !(i+1 <= 8 && i+2 <= 8) {
			continue
		}

		var pu oneDragonPuZi
		pu.PuZi[0] = int(i)
		pu.PuZi[1] = int(i + 1)
		pu.PuZi[2] = int(i + 2)
		puziChan <- pu
	}

	close(puziChan)
}

func oneDragonCalcKey(handCards []int) int {
	var slot [9]int
	for _, v := range handCards {
		slot[v]++
	}

	num := 0
	for _, v := range slot {
		num = num*10 + v
	}
	return num
}
