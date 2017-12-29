package main

import (
	"testing"

	"github.com/MDGSF/MJHuPai/Go/sxtjmj"
)

func Test1(t *testing.T) {
	handCards := []int{31, 31}
	if ok, _ := sxtjmj.CanHu(handCards, 1, handCards[0], false, false, false); !ok {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestOneJiang(t *testing.T) {
	for i := sxtjmj.MAN; i <= sxtjmj.CHU; i++ {
		if sxtjmj.IsValidCard(int(i)) {
			handCards := []int{}
			handCards = append(handCards, int(i))
			handCards = append(handCards, int(i))
			if ok, _ := sxtjmj.CanHu(handCards, 1, handCards[0], false, false, false); !ok {
				sxtjmj.ShowHandCards(handCards)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestOneJiangWithOnePu(t *testing.T) {
	for i := sxtjmj.MAN; i <= sxtjmj.CHU; i++ {
		if !sxtjmj.IsValidCard(int(i)) {
			continue
		}

		handCards := []int{}
		handCards = append(handCards, int(i))
		handCards = append(handCards, int(i))

		for j := sxtjmj.MAN; j <= sxtjmj.CHU; j++ {
			if !sxtjmj.IsValidCard(int(j)) {
				continue
			}
			if j == i {
				continue
			}

			var handCardsTemp = handCards

			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j))

			if ok, _ := sxtjmj.CanHu(handCardsTemp, 1, handCards[0], false, false, false); !ok {
				sxtjmj.ShowHandCards(handCardsTemp)
				t.Error("CanHu failed.")
			}
		}

		for j := sxtjmj.MAN; j <= sxtjmj.CHU; j++ {
			if !((sxtjmj.IsCharacter(j) && sxtjmj.IsCharacter(j+1) && sxtjmj.IsCharacter(j+2)) ||
				(sxtjmj.IsBamboo(j) && sxtjmj.IsBamboo(j+1) && sxtjmj.IsBamboo(j+2)) ||
				(sxtjmj.IsDot(j) && sxtjmj.IsDot(j+1) && sxtjmj.IsDot(j+2))) {
				continue
			}

			var handCardsTemp = handCards

			handCardsTemp = append(handCardsTemp, int(j))
			handCardsTemp = append(handCardsTemp, int(j+1))
			handCardsTemp = append(handCardsTemp, int(j+2))

			if ok, _ := sxtjmj.CanHu(handCardsTemp, 1, handCards[0], false, false, false); !ok {
				sxtjmj.ShowHandCards(handCardsTemp)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestOneJiangWithTwoPu(t *testing.T) {
	for i := sxtjmj.MAN; i <= sxtjmj.CHU; i++ {
		if !sxtjmj.IsValidCard(int(i)) {
			continue
		}

		handCards := []int{}
		handCards = append(handCards, int(i))
		handCards = append(handCards, int(i))

		for j := sxtjmj.MAN; j <= sxtjmj.CHU; j++ {
			if !sxtjmj.IsValidCard(int(j)) {
				continue
			}

			var handCardsj = handCards
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j))
			if !sxtjmj.IsValidHandCards(handCardsj) {
				continue
			}

			for k := sxtjmj.MAN; k <= sxtjmj.CHU; k++ {
				if !sxtjmj.IsValidCard(int(j)) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k))
				if !sxtjmj.IsValidHandCards(handCardsk) {
					continue
				}

				if ok, _ := sxtjmj.CanHu(handCardsk, 1, handCards[0], false, false, false); !ok {
					sxtjmj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}

			for k := sxtjmj.MAN; k <= sxtjmj.CHU; k++ {
				if !((sxtjmj.IsCharacter(k) && sxtjmj.IsCharacter(k+1) && sxtjmj.IsCharacter(k+2)) ||
					(sxtjmj.IsBamboo(k) && sxtjmj.IsBamboo(k+1) && sxtjmj.IsBamboo(k+2)) ||
					(sxtjmj.IsDot(k) && sxtjmj.IsDot(k+1) && sxtjmj.IsDot(k+2))) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k+1))
				handCardsk = append(handCardsk, int(k+2))
				if !sxtjmj.IsValidHandCards(handCardsk) {
					continue
				}

				if ok, _ := sxtjmj.CanHu(handCardsk, 1, handCards[0], false, false, false); !ok {
					sxtjmj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}
		}

		for j := sxtjmj.MAN; j <= sxtjmj.CHU; j++ {
			if !((sxtjmj.IsCharacter(j) && sxtjmj.IsCharacter(j+1) && sxtjmj.IsCharacter(j+2)) ||
				(sxtjmj.IsBamboo(j) && sxtjmj.IsBamboo(j+1) && sxtjmj.IsBamboo(j+2)) ||
				(sxtjmj.IsDot(j) && sxtjmj.IsDot(j+1) && sxtjmj.IsDot(j+2))) {
				continue
			}

			var handCardsj = handCards
			handCardsj = append(handCardsj, int(j))
			handCardsj = append(handCardsj, int(j+1))
			handCardsj = append(handCardsj, int(j+2))
			if !sxtjmj.IsValidHandCards(handCardsj) {
				continue
			}

			for k := sxtjmj.MAN; k <= sxtjmj.CHU; k++ {
				if !((sxtjmj.IsCharacter(k) && sxtjmj.IsCharacter(k+1) && sxtjmj.IsCharacter(k+2)) ||
					(sxtjmj.IsBamboo(k) && sxtjmj.IsBamboo(k+1) && sxtjmj.IsBamboo(k+2)) ||
					(sxtjmj.IsDot(k) && sxtjmj.IsDot(k+1) && sxtjmj.IsDot(k+2))) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, int(k))
				handCardsk = append(handCardsk, int(k+1))
				handCardsk = append(handCardsk, int(k+2))
				if !sxtjmj.IsValidHandCards(handCardsk) {
					continue
				}

				if ok, _ := sxtjmj.CanHu(handCardsk, 1, handCards[0], false, false, false); !ok {
					sxtjmj.ShowHandCards(handCardsk)
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
			if !sxtjmj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !sxtjmj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !sxtjmj.IsValidHandCards(handCards3) {
						continue
					}

					count++

					if ok, _ := sxtjmj.CanHu(handCards3, 1, handCards[0], false, false, false); !ok {
						sxtjmj.ShowHandCards(handCards3)
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
			if !sxtjmj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !sxtjmj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !sxtjmj.IsValidHandCards(handCards3) {
						continue
					}

					fourPuZiChan := make(chan PuZi)
					go genPuZi(fourPuZiChan)
					for four := range fourPuZiChan {

						var handCards4 = handCards3
						handCards4 = append(handCards4, four.PuZi[0])
						handCards4 = append(handCards4, four.PuZi[1])
						handCards4 = append(handCards4, four.PuZi[2])
						if !sxtjmj.IsValidHandCards(handCards4) {
							continue
						}

						count++

						if ok, _ := sxtjmj.CanHu(handCards4, 1, handCards[0], false, false, false); !ok {
							sxtjmj.ShowHandCards(handCards4)
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
	for i := sxtjmj.MAN; i <= sxtjmj.CHU; i++ {
		if !sxtjmj.IsValidCard(int(i)) {
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
	for i := sxtjmj.MAN; i <= sxtjmj.CHU; i++ {
		if !sxtjmj.IsValidCard(int(i)) {
			continue
		}

		var pu PuZi
		pu.PuZi[0] = int(i)
		pu.PuZi[1] = int(i)
		pu.PuZi[2] = int(i)
		puziChan <- pu
	}

	for i := sxtjmj.MAN; i <= sxtjmj.CHU; i++ {
		if !((sxtjmj.IsCharacter(i) && sxtjmj.IsCharacter(i+1) && sxtjmj.IsCharacter(i+2)) ||
			(sxtjmj.IsBamboo(i) && sxtjmj.IsBamboo(i+1) && sxtjmj.IsBamboo(i+2)) ||
			(sxtjmj.IsDot(i) && sxtjmj.IsDot(i+1) && sxtjmj.IsDot(i+2))) {
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
