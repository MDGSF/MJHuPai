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
	laizi := []int{0x11}
	if !kddmj.CanHuWithLaiZi(handCards, laizi) {
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

func TestLaiZiOneJiangWithFourPu(t *testing.T) {

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
						if !kddmj.CanHuWithLaiZi(handCards4, laizi) {
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
