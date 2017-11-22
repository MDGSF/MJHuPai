package main

import (
	"jian/mj"
	"testing"
)

func Test1(t *testing.T) {
	handCards := []mj.Card{0x43, 0x43, 0x01, 0x02, 0x03}
	if !mj.CanHu(handCards) {
		mj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestLaiZi(t *testing.T) {
	handCards := []mj.Card{0x43, 0x43, 0x01, 0x02, 0x03}
	laizi := []mj.Card{0x11}
	if !mj.CanHuWithLaiZi(handCards, laizi) {
		mj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestOneJiang(t *testing.T) {
	for i := 1; i <= mj.MAX_CARD; i++ {
		if mj.IsValidCard(mj.Card(i)) {
			handCards := []mj.Card{}
			handCards = append(handCards, mj.Card(i))
			handCards = append(handCards, mj.Card(i))
			if !mj.CanHu(handCards) {
				mj.ShowHandCards(handCards)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestOneJiangWithOnePu(t *testing.T) {
	for i := 1; i <= mj.MAX_CARD; i++ {
		if !mj.IsValidCard(mj.Card(i)) {
			continue
		}

		handCards := []mj.Card{}
		handCards = append(handCards, mj.Card(i))
		handCards = append(handCards, mj.Card(i))

		for j := 1; j <= mj.MAX_CARD; j++ {
			if !mj.IsValidCard(mj.Card(j)) {
				continue
			}
			if j == i {
				continue
			}

			var handCardsTemp = handCards

			handCardsTemp = append(handCardsTemp, mj.Card(j))
			handCardsTemp = append(handCardsTemp, mj.Card(j))
			handCardsTemp = append(handCardsTemp, mj.Card(j))

			if !mj.CanHu(handCardsTemp) {
				mj.ShowHandCards(handCardsTemp)
				t.Error("CanHu failed.")
			}
		}

		for j := 1; j <= mj.MAX_CARD; j++ {
			if !mj.IsXuShu(mj.Card(j)) || !mj.IsXuShu(mj.Card(j+1)) || !mj.IsXuShu(mj.Card(j+2)) {
				continue
			}

			var handCardsTemp = handCards

			handCardsTemp = append(handCardsTemp, mj.Card(j))
			handCardsTemp = append(handCardsTemp, mj.Card(j+1))
			handCardsTemp = append(handCardsTemp, mj.Card(j+2))

			if !mj.CanHu(handCardsTemp) {
				mj.ShowHandCards(handCardsTemp)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestOneJiangWithTwoPu(t *testing.T) {
	for i := 1; i <= mj.MAX_CARD; i++ {
		if !mj.IsValidCard(mj.Card(i)) {
			continue
		}

		handCards := []mj.Card{}
		handCards = append(handCards, mj.Card(i))
		handCards = append(handCards, mj.Card(i))

		for j := 1; j <= mj.MAX_CARD; j++ {
			if !mj.IsValidCard(mj.Card(j)) {
				continue
			}

			var handCardsj = handCards
			handCardsj = append(handCardsj, mj.Card(j))
			handCardsj = append(handCardsj, mj.Card(j))
			handCardsj = append(handCardsj, mj.Card(j))
			if !mj.IsValidHandCards(handCardsj) {
				continue
			}

			for k := 1; k <= mj.MAX_CARD; k++ {
				if !mj.IsValidCard(mj.Card(j)) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, mj.Card(k))
				handCardsk = append(handCardsk, mj.Card(k))
				handCardsk = append(handCardsk, mj.Card(k))
				if !mj.IsValidHandCards(handCardsk) {
					continue
				}

				if !mj.CanHu(handCardsk) {
					mj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}

			for k := 1; k <= mj.MAX_CARD; k++ {
				if !mj.IsXuShu(mj.Card(k)) || !mj.IsXuShu(mj.Card(k+1)) || !mj.IsXuShu(mj.Card(k+2)) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, mj.Card(k))
				handCardsk = append(handCardsk, mj.Card(k+1))
				handCardsk = append(handCardsk, mj.Card(k+2))
				if !mj.IsValidHandCards(handCardsk) {
					continue
				}

				if !mj.CanHu(handCardsk) {
					mj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}
		}

		for j := 1; j <= mj.MAX_CARD; j++ {
			if !mj.IsXuShu(mj.Card(j)) || !mj.IsXuShu(mj.Card(j+1)) || !mj.IsXuShu(mj.Card(j+2)) {
				continue
			}

			var handCardsj = handCards
			handCardsj = append(handCardsj, mj.Card(j))
			handCardsj = append(handCardsj, mj.Card(j+1))
			handCardsj = append(handCardsj, mj.Card(j+2))
			if !mj.IsValidHandCards(handCardsj) {
				continue
			}

			for k := 1; k <= mj.MAX_CARD; k++ {
				if !mj.IsXuShu(mj.Card(k)) || !mj.IsXuShu(mj.Card(k+1)) || !mj.IsXuShu(mj.Card(k+2)) {
					continue
				}

				var handCardsk = handCardsj
				handCardsk = append(handCardsk, mj.Card(k))
				handCardsk = append(handCardsk, mj.Card(k+1))
				handCardsk = append(handCardsk, mj.Card(k+2))
				if !mj.IsValidHandCards(handCardsk) {
					continue
				}

				if !mj.CanHu(handCardsk) {
					mj.ShowHandCards(handCardsk)
					t.Error("CanHu failed.")
				}
			}
		}
	}
}

func TestOneJiangWithThreePu(t *testing.T) {

	count := 0
	t.Log("count = ", count)

	jiangChan := make(chan mj.Card)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []mj.Card{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		onePuZiChan := make(chan PuZi)
		go genPuZi(onePuZiChan)
		for one := range onePuZiChan {

			var handCards1 = handCards
			handCards1 = append(handCards1, one.PuZi[0])
			handCards1 = append(handCards1, one.PuZi[1])
			handCards1 = append(handCards1, one.PuZi[2])
			if !mj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !mj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !mj.IsValidHandCards(handCards3) {
						continue
					}

					count++

					if !mj.CanHu(handCards3) {
						mj.ShowHandCards(handCards3)
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

	jiangChan := make(chan mj.Card)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []mj.Card{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		onePuZiChan := make(chan PuZi)
		go genPuZi(onePuZiChan)
		for one := range onePuZiChan {

			var handCards1 = handCards
			handCards1 = append(handCards1, one.PuZi[0])
			handCards1 = append(handCards1, one.PuZi[1])
			handCards1 = append(handCards1, one.PuZi[2])
			if !mj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !mj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !mj.IsValidHandCards(handCards3) {
						continue
					}

					fourPuZiChan := make(chan PuZi)
					go genPuZi(fourPuZiChan)
					for four := range fourPuZiChan {

						var handCards4 = handCards3
						handCards4 = append(handCards4, four.PuZi[0])
						handCards4 = append(handCards4, four.PuZi[1])
						handCards4 = append(handCards4, four.PuZi[2])
						if !mj.IsValidHandCards(handCards4) {
							continue
						}

						count++

						if !mj.CanHu(handCards4) {
							mj.ShowHandCards(handCards4)
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

	jiangChan := make(chan mj.Card)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []mj.Card{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		onePuZiChan := make(chan PuZi)
		go genPuZi(onePuZiChan)
		for one := range onePuZiChan {

			var handCards1 = handCards
			handCards1 = append(handCards1, one.PuZi[0])
			handCards1 = append(handCards1, one.PuZi[1])
			handCards1 = append(handCards1, one.PuZi[2])
			if !mj.IsValidHandCards(handCards1) {
				continue
			}

			twoPuZiChan := make(chan PuZi)
			go genPuZi(twoPuZiChan)
			for two := range twoPuZiChan {

				var handCards2 = handCards1
				handCards2 = append(handCards2, two.PuZi[0])
				handCards2 = append(handCards2, two.PuZi[1])
				handCards2 = append(handCards2, two.PuZi[2])
				if !mj.IsValidHandCards(handCards2) {
					continue
				}

				threePuZiChan := make(chan PuZi)
				go genPuZi(threePuZiChan)
				for three := range threePuZiChan {

					var handCards3 = handCards2
					handCards3 = append(handCards3, three.PuZi[0])
					handCards3 = append(handCards3, three.PuZi[1])
					handCards3 = append(handCards3, three.PuZi[2])
					if !mj.IsValidHandCards(handCards3) {
						continue
					}

					fourPuZiChan := make(chan PuZi)
					go genPuZi(fourPuZiChan)
					for four := range fourPuZiChan {

						var handCards4 = handCards3
						handCards4 = append(handCards4, four.PuZi[0])
						handCards4 = append(handCards4, four.PuZi[1])
						handCards4 = append(handCards4, four.PuZi[2])
						if !mj.IsValidHandCards(handCards4) {
							continue
						}

						count++

						for i := 1; i <= mj.MAX_CARD; i++ {
							laizi := []mj.Card{}
							laizi = append(laizi, mj.Card(i))
							if !mj.CanHuWithLaiZi(handCards4, laizi) {
								mj.ShowHandCards(handCards4)
								t.Error("CanHu failed.")
							}
						}

					}

				}

			}

		}
	}

	t.Log("count = ", count)
}

func genJiang(jiangChan chan mj.Card) {
	for i := 1; i <= mj.MAX_CARD; i++ {
		if !mj.IsValidCard(mj.Card(i)) {
			continue
		}

		jiangChan <- mj.Card(i)
	}
	close(jiangChan)
}

type PuZi struct {
	PuZi [3]mj.Card
}

func genPuZi(puziChan chan PuZi) {
	for i := 1; i <= mj.MAX_CARD; i++ {
		if !mj.IsValidCard(mj.Card(i)) {
			continue
		}

		var pu PuZi
		pu.PuZi[0] = mj.Card(i)
		pu.PuZi[1] = mj.Card(i)
		pu.PuZi[2] = mj.Card(i)
		puziChan <- pu
	}

	for i := 1; i <= mj.MAX_CARD; i++ {
		if !mj.IsXuShu(mj.Card(i)) || !mj.IsXuShu(mj.Card(i+1)) || !mj.IsXuShu(mj.Card(i+2)) {
			continue
		}

		var pu PuZi
		pu.PuZi[0] = mj.Card(i)
		pu.PuZi[1] = mj.Card(i + 1)
		pu.PuZi[2] = mj.Card(i + 2)
		puziChan <- pu
	}

	close(puziChan)
}

func main() {

}
