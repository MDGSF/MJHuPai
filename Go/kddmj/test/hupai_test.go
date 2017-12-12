package main

import (
	"testing"

	"github.com/MDGSF/MJHuPai/Go/kddmj"
)

func Test1(t *testing.T) {
	handCards := []int{31, 31}
	if !kddmj.CanHu(handCards) {
		kddmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestLaiZi(t *testing.T) {
	handCards := []int{0, 0, 1, 2, 3}
	laizi := []int{9}
	ok, _ := kddmj.CanHuWithLaiZi(handCards, laizi)
	if !ok {
		kddmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestOneJiang(t *testing.T) {
	for i := kddmj.MAN; i <= kddmj.CHU; i++ {
		if kddmj.IsValidCard(int(i)) {
			handCards := []int{}
			handCards = append(handCards, int(i))
			handCards = append(handCards, int(i))
			if !kddmj.CanHu(handCards) {
				kddmj.ShowHandCards(handCards)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestOneJiangWithOnePu(t *testing.T) {
	for i := kddmj.MAN; i <= kddmj.CHU; i++ {
		if !kddmj.IsValidCard(int(i)) {
			continue
		}

		handCards := []int{}
		handCards = append(handCards, int(i))
		handCards = append(handCards, int(i))

		for j := kddmj.MAN; j <= kddmj.CHU; j++ {
			if !kddmj.IsValidCard(int(j)) {
				continue
			}
			if j == i {
				continue
			}

			var handCardsTemp = handCards

			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j))

			if !kddmj.CanHu(handCardsTemp) {
				kddmj.ShowHandCards(handCardsTemp)
				t.Error("CanHu failed.")
			}
		}

		for j := kddmj.MAN; j <= kddmj.CHU; j++ {
			if !((kddmj.IsCharacter(j) && kddmj.IsCharacter(j+1) && kddmj.IsCharacter(j+2)) ||
				(kddmj.IsBamboo(j) && kddmj.IsBamboo(j+1) && kddmj.IsBamboo(j+2)) ||
				(kddmj.IsDot(j) && kddmj.IsDot(j+1) && kddmj.IsDot(j+2))) {
				continue
			}

			var handCardsTemp = handCards

			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j+1))
			handCardsTemp = append(handCardsTemp, int(j+2))

			if !kddmj.CanHu(handCardsTemp) {
				kddmj.ShowHandCards(handCardsTemp)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestOneJiangWithTwoPu(t *testing.T) {
	for i := kddmj.MAN; i <= kddmj.CHU; i++ {
		if !kddmj.IsValidCard(int(i)) {
			continue
		}

		handCards := []int{}
		handCards = append(handCards, int(i))
		handCards = append(handCards, int(i))

		for j := kddmj.MAN; j <= kddmj.CHU; j++ {
			if !kddmj.IsValidCard(int(j)) {
				continue
			}

			var handCardsj = handCards
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j))
			if !kddmj.IsValidHandCards(handCardsj) {
				continue
			}

			for k := kddmj.MAN; k <= kddmj.CHU; k++ {
				if !kddmj.IsValidCard(int(j)) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k))
				if !kddmj.IsValidHandCards(handCardsk) {
					continue
				}

				if !kddmj.CanHu(handCardsk) {
					kddmj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}

			for k := kddmj.MAN; k <= kddmj.CHU; k++ {
				if !((kddmj.IsCharacter(k) && kddmj.IsCharacter(k+1) && kddmj.IsCharacter(k+2)) ||
					(kddmj.IsBamboo(k) && kddmj.IsBamboo(k+1) && kddmj.IsBamboo(k+2)) ||
					(kddmj.IsDot(k) && kddmj.IsDot(k+1) && kddmj.IsDot(k+2))) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k+1))
				handCardsk = append(handCardsk, int(k+2))
				if !kddmj.IsValidHandCards(handCardsk) {
					continue
				}

				if !kddmj.CanHu(handCardsk) {
					kddmj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}
		}

		for j := kddmj.MAN; j <= kddmj.CHU; j++ {
			if !((kddmj.IsCharacter(j) && kddmj.IsCharacter(j+1) && kddmj.IsCharacter(j+2)) ||
				(kddmj.IsBamboo(j) && kddmj.IsBamboo(j+1) && kddmj.IsBamboo(j+2)) ||
				(kddmj.IsDot(j) && kddmj.IsDot(j+1) && kddmj.IsDot(j+2))) {
				continue
			}

			var handCardsj = handCards
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j+1))
			handCardsj = append(handCardsj, int(j+2))
			if !kddmj.IsValidHandCards(handCardsj) {
				continue
			}

			for k := kddmj.MAN; k <= kddmj.CHU; k++ {
				if !((kddmj.IsCharacter(k) && kddmj.IsCharacter(k+1) && kddmj.IsCharacter(k+2)) ||
					(kddmj.IsBamboo(k) && kddmj.IsBamboo(k+1) && kddmj.IsBamboo(k+2)) ||
					(kddmj.IsDot(k) && kddmj.IsDot(k+1) && kddmj.IsDot(k+2))) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k+1))
				handCardsk = append(handCardsk, int(k+2))
				if !kddmj.IsValidHandCards(handCardsk) {
					continue
				}

				if !kddmj.CanHu(handCardsk) {
					kddmj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}
		}
	}
}

func TestOneJiangWithThreePu(t *testing.T) {

	count := 0
	t.Log("count = ", count)

	jiangChan := make(chan int)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []int{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		onePuZiChan := make(chan PuZi)
		go genPuZi(onePuZiChan)
		for one := range onePuZiChan {

			var handCards1 = handCards
			handCards1 = append(handCards1, one.PuZi[0])
			handCards1 = append(handCards1, one.PuZi[1])
			handCards1 = append(handCards1, one.PuZi[2])
			if !kddmj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !kddmj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !kddmj.IsValidHandCards(handCards3) {
						continue
					}

					count++

					if !kddmj.CanHu(handCards3) {
						kddmj.ShowHandCards(handCards3)
						t.Error("CanHu failed.")
					}
				}

			}

		}
	}

	t.Log("count = ", count)
}

func TestOneJiangWithFourPu(t *testing.T) {

	count := 0
	t.Log("count = ", count)

	jiangChan := make(chan int)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []int{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		onePuZiChan := make(chan PuZi)
		go genPuZi(onePuZiChan)
		for one := range onePuZiChan {

			var handCards1 = handCards
			handCards1 = append(handCards1, one.PuZi[0])
			handCards1 = append(handCards1, one.PuZi[1])
			handCards1 = append(handCards1, one.PuZi[2])
			if !kddmj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !kddmj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !kddmj.IsValidHandCards(handCards3) {
						continue
					}

					fourPuZiChan := make(chan PuZi)
					go genPuZi(fourPuZiChan)
					for four := range fourPuZiChan {

						var handCards4 = handCards3
						handCards4 = append(handCards4, four.PuZi[0])
						handCards4 = append(handCards4, four.PuZi[1])
						handCards4 = append(handCards4, four.PuZi[2])
						if !kddmj.IsValidHandCards(handCards4) {
							continue
						}

						count++

						if !kddmj.CanHu(handCards4) {
							kddmj.ShowHandCards(handCards4)
							t.Error("CanHu failed.")
						}
					}

				}

			}

		}
	}

	t.Log("count = ", count)
}

func TestLaiZiOneJiang1(t *testing.T) {
	for i := kddmj.MAN; i <= kddmj.CHU; i++ {
		if kddmj.IsValidCard(int(i)) {
			handCards := []int{}
			handCards = append(handCards, int(i))
			handCards = append(handCards, int(i))
			laizi := []int{}
			ok, dianshu := kddmj.CanHuWithLaiZi(handCards, laizi)
			if !ok || dianshu != 0 {
				kddmj.ShowHandCards(handCards)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestLaiZiOneJiang2(t *testing.T) {
	for i := kddmj.MAN; i <= kddmj.CHU; i++ {
		if kddmj.IsValidCard(int(i)) {
			handCards := []int{}
			handCards = append(handCards, int(i))
			handCards = append(handCards, int(i))
			laizi := []int{}
			laizi = append(laizi, int(i))
			ok, dianshu := kddmj.CanHuWithLaiZi(handCards, laizi)
			if !ok || dianshu != 10 {
				kddmj.ShowHandCards(handCards)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestLaiZiOneJiangWithOnePu1(t *testing.T) {
	jiangChan := make(chan int)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []int{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		onePuZiChan := make(chan PuZi)
		go genPuZi(onePuZiChan)
		for one := range onePuZiChan {

			var handCards1 = handCards
			handCards1 = append(handCards1, one.PuZi[0])
			handCards1 = append(handCards1, one.PuZi[1])
			handCards1 = append(handCards1, one.PuZi[2])
			if !kddmj.IsValidHandCards(handCards1) {
				continue
			}

			laizi := []int{}
			laizi = append(laizi, jiang)
			ok, dianshu := kddmj.CanHuWithLaiZi(handCards1, laizi)
			if !ok || dianshu != 10 {
				kddmj.ShowHandCards(handCards1)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestLaiZiOneJiangWithTwoPu(t *testing.T) {
	jiangChan := make(chan int)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []int{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		AddPuZiToHandCards(t, handCards, jiang, 2)
	}
}

func TestLaiZiOneJiangWithThreePu(t *testing.T) {
	jiangChan := make(chan int)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []int{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		AddPuZiToHandCards(t, handCards, jiang, 3)
	}
}

func TestLaiZiOneJiangWithFourPu(t *testing.T) {
	jiangChan := make(chan int)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []int{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		AddPuZiToHandCards(t, handCards, jiang, 4)
	}
}

func AddPuZiToHandCards(t *testing.T, handCards []int, jiang int, level int) {
	if level <= 0 {
		laizi := []int{}
		laizi = append(laizi, jiang)
		ok, dianshu := kddmj.CanHuWithLaiZi(handCards, laizi)
		if !ok || dianshu != 10 {
			kddmj.ShowHandCards(handCards)
			t.Error("CanHu failed.")
		}
		return
	}

	onePuZiChan := make(chan PuZi)
	go genPuZi(onePuZiChan)
	for one := range onePuZiChan {

		var handCards1 = handCards
		handCards1 = append(handCards1, one.PuZi[0])
		handCards1 = append(handCards1, one.PuZi[1])
		handCards1 = append(handCards1, one.PuZi[2])
		if !kddmj.IsValidHandCards(handCards1) {
			continue
		}

		AddPuZiToHandCards(t, handCards1, jiang, level-1)
	}
}

func TestLaiZiOneJiangWithFourPu2(t *testing.T) {

	count := 0
	t.Log("count = ", count)

	jiangChan := make(chan int)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []int{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		onePuZiChan := make(chan PuZi)
		go genPuZi(onePuZiChan)
		for one := range onePuZiChan {

			var handCards1 = handCards
			handCards1 = append(handCards1, one.PuZi[0])
			handCards1 = append(handCards1, one.PuZi[1])
			handCards1 = append(handCards1, one.PuZi[2])
			if !kddmj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !kddmj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !kddmj.IsValidHandCards(handCards3) {
						continue
					}

					fourPuZiChan := make(chan PuZi)
					go genPuZi(fourPuZiChan)
					for four := range fourPuZiChan {

						var handCards4 = handCards3
						handCards4 = append(handCards4, four.PuZi[0])
						handCards4 = append(handCards4, four.PuZi[1])
						handCards4 = append(handCards4, four.PuZi[2])
						if !kddmj.IsValidHandCards(handCards4) {
							continue
						}

						count++

						//for i := 1; i <= kddmj.MaxCard; i++ {
						laizi := []int{0x01}
						//laizi = append(laizi, int(i))
						ok, _ := kddmj.CanHuWithLaiZi(handCards4, laizi)
						if !ok {
							kddmj.ShowHandCards(handCards4)
							t.Error("CanHu failed.")
						}
						//}

					}

				}

			}

		}
	}

	t.Log("count = ", count)
}

func genJiang(jiangChan chan int) {
	for i := kddmj.MAN; i <= kddmj.CHU; i++ {
		if !kddmj.IsValidCard(int(i)) {
			continue
		}

		jiangChan <- int(i)
	}
	close(jiangChan)
}

type PuZi struct {
	PuZi [3]int
}

func genPuZi(puziChan chan PuZi) {
	for i := kddmj.MAN; i <= kddmj.CHU; i++ {
		if !kddmj.IsValidCard(int(i)) {
			continue
		}

		var pu PuZi
		pu.PuZi[0] = int(i)
		pu.PuZi[1] = int(i)
		pu.PuZi[2] = int(i)
		puziChan <- pu
	}

	for i := kddmj.MAN; i <= kddmj.CHU; i++ {
		if !((kddmj.IsCharacter(i) && kddmj.IsCharacter(i+1) && kddmj.IsCharacter(i+2)) ||
			(kddmj.IsBamboo(i) && kddmj.IsBamboo(i+1) && kddmj.IsBamboo(i+2)) ||
			(kddmj.IsDot(i) && kddmj.IsDot(i+1) && kddmj.IsDot(i+2))) {
			continue
		}

		var pu PuZi
		pu.PuZi[0] = int(i)
		pu.PuZi[1] = int(i + 1)
		pu.PuZi[2] = int(i + 2)
		puziChan <- pu
	}

	close(puziChan)
}

// func getMaxOneCard(slots [kddmj.TILEMAX]int) (bool, int) {
// 	for i := kddmj.TON; i <= kddmj.CHU; i++ {
// 		if slots[i] == 1 {
// 			return true, i
// 		}
// 	}

// 	for i := 1; i <= 9; i++ {
// 		var k1 = 1*9 - i
// 		var k2 = 2*9 - i
// 		var k3 = 3*9 - i

// 		if slots[k1] > 1 || slots[k2] > 1 || slots[k3] > 1 {
// 			return false, 0
// 		}

// 		count := 0
// 		k := 0

// 		if slots[k1] == 1 {
// 			count++
// 			k = k1
// 		}

// 		if slots[k2] == 1 {
// 			count++
// 			k = k2
// 		}

// 		if slots[k3] == 1 {
// 			count++
// 			k = k3
// 		}

// 		if count == 1 {
// 			return true, k
// 		} else if count > 1 {
// 			return false, 0
// 		}
// 	}

// 	return false, 0
// }

// func TestOneLaiZiForAll(t *testing.T) {
// 	jiangChan := make(chan int)
// 	go genJiang(jiangChan)
// 	for jiang := range jiangChan {

// 		handCards := []int{}
// 		handCards = append(handCards, jiang)
// 		handCards = append(handCards, jiang)

// 		ret := OneLaiZiAddPuZiToHandCards(t, handCards, jiang, 4)
// 		if !ret {
// 			return
// 		}
// 	}
// }

// func OneLaiZiAddPuZiToHandCards(t *testing.T, handCards []int, jiang int, level int) bool {
// 	if level <= 0 {

// 		slots := kddmj.GenSlots(handCards)
// 		ret, maxOneCard := getMaxOneCard(slots)
// 		if ret {
// 			laizi := []int{}
// 			laizi = append(laizi, maxOneCard)
// 			ok, dianshu := kddmj.CanHuWithLaiZi(handCards, laizi)
// 			if !ok || dianshu != kddmj.DianShuTable[maxOneCard] {
// 				fmt.Println("slots=", slots, ", maxOneCard=", maxOneCard,
// 					"kddmj.DianShuTable[maxOneCard]=", kddmj.DianShuTable[maxOneCard], "dianshu=", dianshu, ", ")
// 				kddmj.ShowHandCards(handCards)
// 				t.Error("CanHu failed.")
// 				return false
// 			}
// 		}

// 		return true
// 	}

// 	onePuZiChan := make(chan PuZi)
// 	go genPuZi(onePuZiChan)
// 	for one := range onePuZiChan {

// 		var handCards1 = handCards
// 		handCards1 = append(handCards1, one.PuZi[0])
// 		handCards1 = append(handCards1, one.PuZi[1])
// 		handCards1 = append(handCards1, one.PuZi[2])
// 		if !kddmj.IsValidHandCards(handCards1) {
// 			continue
// 		}

// 		ret := OneLaiZiAddPuZiToHandCards(t, handCards1, jiang, level-1)
// 		if !ret {
// 			return false
// 		}
// 	}

// 	return true
// }
