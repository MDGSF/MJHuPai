package main

import (
	"sort"
	"testing"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

func Test1(t *testing.T) {
	handCards := []int{31, 31}
	if ok, _ := sxtdhmj.CanHu(handCards, false, false, false); !ok {
		sxtdhmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestOneJiang(t *testing.T) {
	for i := sxtdhmj.MAN; i <= sxtdhmj.CHU; i++ {
		if sxtdhmj.IsValidCard(int(i)) {
			handCards := []int{}
			handCards = append(handCards, int(i))
			handCards = append(handCards, int(i))
			if ok, _ := sxtdhmj.CanHu(handCards, false, false, false); !ok {
				sxtdhmj.ShowHandCards(handCards)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestOneJiangWithOnePu(t *testing.T) {
	for i := sxtdhmj.MAN; i <= sxtdhmj.CHU; i++ {
		if !sxtdhmj.IsValidCard(int(i)) {
			continue
		}

		handCards := []int{}
		handCards = append(handCards, int(i))
		handCards = append(handCards, int(i))

		for j := sxtdhmj.MAN; j <= sxtdhmj.CHU; j++ {
			if !sxtdhmj.IsValidCard(int(j)) {
				continue
			}
			if j == i {
				continue
			}

			var handCardsTemp = handCards

			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j))

			if ok, _ := sxtdhmj.CanHu(handCardsTemp, false, false, false); !ok {
				sxtdhmj.ShowHandCards(handCardsTemp)
				t.Error("CanHu failed.")
			}
		}

		for j := sxtdhmj.MAN; j <= sxtdhmj.CHU; j++ {
			if !((sxtdhmj.IsCharacter(j) && sxtdhmj.IsCharacter(j+1) && sxtdhmj.IsCharacter(j+2)) ||
				(sxtdhmj.IsBamboo(j) && sxtdhmj.IsBamboo(j+1) && sxtdhmj.IsBamboo(j+2)) ||
				(sxtdhmj.IsDot(j) && sxtdhmj.IsDot(j+1) && sxtdhmj.IsDot(j+2))) {
				continue
			}

			var handCardsTemp = handCards

			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j+1))
			handCardsTemp = append(handCardsTemp, int(j+2))

			if ok, _ := sxtdhmj.CanHu(handCardsTemp, false, false, false); !ok {
				sxtdhmj.ShowHandCards(handCardsTemp)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestOneJiangWithTwoPu(t *testing.T) {
	for i := sxtdhmj.MAN; i <= sxtdhmj.CHU; i++ {
		if !sxtdhmj.IsValidCard(int(i)) {
			continue
		}

		handCards := []int{}
		handCards = append(handCards, int(i))
		handCards = append(handCards, int(i))

		for j := sxtdhmj.MAN; j <= sxtdhmj.CHU; j++ {
			if !sxtdhmj.IsValidCard(int(j)) {
				continue
			}

			var handCardsj = handCards
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j))
			if !sxtdhmj.IsValidHandCards(handCardsj) {
				continue
			}

			for k := sxtdhmj.MAN; k <= sxtdhmj.CHU; k++ {
				if !sxtdhmj.IsValidCard(int(j)) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k))
				if !sxtdhmj.IsValidHandCards(handCardsk) {
					continue
				}

				if ok, _ := sxtdhmj.CanHu(handCardsk, false, false, false); !ok {
					sxtdhmj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}

			for k := sxtdhmj.MAN; k <= sxtdhmj.CHU; k++ {
				if !((sxtdhmj.IsCharacter(k) && sxtdhmj.IsCharacter(k+1) && sxtdhmj.IsCharacter(k+2)) ||
					(sxtdhmj.IsBamboo(k) && sxtdhmj.IsBamboo(k+1) && sxtdhmj.IsBamboo(k+2)) ||
					(sxtdhmj.IsDot(k) && sxtdhmj.IsDot(k+1) && sxtdhmj.IsDot(k+2))) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k+1))
				handCardsk = append(handCardsk, int(k+2))
				if !sxtdhmj.IsValidHandCards(handCardsk) {
					continue
				}

				if ok, _ := sxtdhmj.CanHu(handCardsk, false, false, false); !ok {
					sxtdhmj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}
		}

		for j := sxtdhmj.MAN; j <= sxtdhmj.CHU; j++ {
			if !((sxtdhmj.IsCharacter(j) && sxtdhmj.IsCharacter(j+1) && sxtdhmj.IsCharacter(j+2)) ||
				(sxtdhmj.IsBamboo(j) && sxtdhmj.IsBamboo(j+1) && sxtdhmj.IsBamboo(j+2)) ||
				(sxtdhmj.IsDot(j) && sxtdhmj.IsDot(j+1) && sxtdhmj.IsDot(j+2))) {
				continue
			}

			var handCardsj = handCards
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j+1))
			handCardsj = append(handCardsj, int(j+2))
			if !sxtdhmj.IsValidHandCards(handCardsj) {
				continue
			}

			for k := sxtdhmj.MAN; k <= sxtdhmj.CHU; k++ {
				if !((sxtdhmj.IsCharacter(k) && sxtdhmj.IsCharacter(k+1) && sxtdhmj.IsCharacter(k+2)) ||
					(sxtdhmj.IsBamboo(k) && sxtdhmj.IsBamboo(k+1) && sxtdhmj.IsBamboo(k+2)) ||
					(sxtdhmj.IsDot(k) && sxtdhmj.IsDot(k+1) && sxtdhmj.IsDot(k+2))) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k+1))
				handCardsk = append(handCardsk, int(k+2))
				if !sxtdhmj.IsValidHandCards(handCardsk) {
					continue
				}

				if ok, _ := sxtdhmj.CanHu(handCardsk, false, false, false); !ok {
					sxtdhmj.ShowHandCards(handCardsk)
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
			if !sxtdhmj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !sxtdhmj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !sxtdhmj.IsValidHandCards(handCards3) {
						continue
					}

					count++

					if ok, _ := sxtdhmj.CanHu(handCards3, false, false, false); !ok {
						sxtdhmj.ShowHandCards(handCards3)
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
			if !sxtdhmj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !sxtdhmj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !sxtdhmj.IsValidHandCards(handCards3) {
						continue
					}

					fourPuZiChan := make(chan PuZi)
					go genPuZi(fourPuZiChan)
					for four := range fourPuZiChan {

						var handCards4 = handCards3
						handCards4 = append(handCards4, four.PuZi[0])
						handCards4 = append(handCards4, four.PuZi[1])
						handCards4 = append(handCards4, four.PuZi[2])
						if !sxtdhmj.IsValidHandCards(handCards4) {
							continue
						}

						count++

						if ok, _ := sxtdhmj.CanHu(handCards4, false, false, false); !ok {
							sxtdhmj.ShowHandCards(handCards4)
							t.Error("CanHu failed.")
						}
					}

				}

			}

		}
	}

	t.Log("count = ", count)
}

func genJiang(jiangChan chan int) {
	for i := sxtdhmj.MAN; i <= sxtdhmj.CHU; i++ {
		if !sxtdhmj.IsValidCard(int(i)) {
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
	for i := sxtdhmj.MAN; i <= sxtdhmj.CHU; i++ {
		if !sxtdhmj.IsValidCard(int(i)) {
			continue
		}

		var pu PuZi
		pu.PuZi[0] = int(i)
		pu.PuZi[1] = int(i)
		pu.PuZi[2] = int(i)
		puziChan <- pu
	}

	for i := sxtdhmj.MAN; i <= sxtdhmj.CHU; i++ {
		if !((sxtdhmj.IsCharacter(i) && sxtdhmj.IsCharacter(i+1) && sxtdhmj.IsCharacter(i+2)) ||
			(sxtdhmj.IsBamboo(i) && sxtdhmj.IsBamboo(i+1) && sxtdhmj.IsBamboo(i+2)) ||
			(sxtdhmj.IsDot(i) && sxtdhmj.IsDot(i+1) && sxtdhmj.IsDot(i+2))) {
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

type HuRet struct {
	handCards []int
	hu        bool
	dianshu   int
	laizi     int
}

type HuArray struct {
	arr []*HuRet
}

var count int

func calcHandCardsKey(handCards []int) (int, []int) {
	var slots [sxtdhmj.TILEMAX]int

	var distinctCards []int
	for _, c := range handCards {
		slots[c]++
		if slots[c] == 1 {
			distinctCards = append(distinctCards, c)
		}
	}

	sort.Ints(slots[:])

	num := 0
	for _, v := range slots {
		num = num*10 + v
	}

	return num, distinctCards
}
