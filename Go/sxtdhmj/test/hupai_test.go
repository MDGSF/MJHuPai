package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/MDGSF/MJHuPai/Go/sxtdhmj"
)

func Test1(t *testing.T) {
	handCards := []int{31, 31}
	if !sxtdhmj.CanHu(handCards) {
		sxtdhmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestLaiZi(t *testing.T) {
	handCards := []int{0, 0, 1, 2, 3}
	laizi := []int{9}
	ok, _ := sxtdhmj.CanHuWithLaiZi(handCards, laizi)
	if !ok {
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
			if !sxtdhmj.CanHu(handCards) {
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

			if !sxtdhmj.CanHu(handCardsTemp) {
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

			if !sxtdhmj.CanHu(handCardsTemp) {
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

				if !sxtdhmj.CanHu(handCardsk) {
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

				if !sxtdhmj.CanHu(handCardsk) {
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

				if !sxtdhmj.CanHu(handCardsk) {
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

					if !sxtdhmj.CanHu(handCards3) {
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

						if !sxtdhmj.CanHu(handCards4) {
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

func TestLaiZiOneJiang1(t *testing.T) {
	for i := sxtdhmj.MAN; i <= sxtdhmj.CHU; i++ {
		if sxtdhmj.IsValidCard(int(i)) {
			handCards := []int{}
			handCards = append(handCards, int(i))
			handCards = append(handCards, int(i))
			laizi := []int{}
			ok, dianshu := sxtdhmj.CanHuWithLaiZi(handCards, laizi)
			if !ok || dianshu != 0 {
				sxtdhmj.ShowHandCards(handCards)
				t.Error("CanHu failed.")
			}
		}
	}
}

func TestLaiZiOneJiang2(t *testing.T) {
	for i := sxtdhmj.MAN; i <= sxtdhmj.CHU; i++ {
		if sxtdhmj.IsValidCard(int(i)) {
			handCards := []int{}
			handCards = append(handCards, int(i))
			handCards = append(handCards, int(i))
			laizi := []int{}
			laizi = append(laizi, int(i))
			ok, dianshu := sxtdhmj.CanHuWithLaiZi(handCards, laizi)
			if !ok || dianshu != 10 {
				sxtdhmj.ShowHandCards(handCards)
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
			if !sxtdhmj.IsValidHandCards(handCards1) {
				continue
			}

			laizi := []int{}
			laizi = append(laizi, jiang)
			ok, dianshu := sxtdhmj.CanHuWithLaiZi(handCards1, laizi)
			if !ok || dianshu != 10 {
				sxtdhmj.ShowHandCards(handCards1)
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

	count = 0
	start := time.Now()

	jiangChan := make(chan int)
	go genJiang(jiangChan)
	for jiang := range jiangChan {

		handCards := []int{}
		handCards = append(handCards, jiang)
		handCards = append(handCards, jiang)

		AddPuZiToHandCards(t, handCards, jiang, 3)
	}

	elapsed := time.Since(start)

	fmt.Println("count = ", count, ", elapsed=", elapsed)
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
		count++
		laizi := []int{}
		laizi = append(laizi, jiang)
		ok, dianshu := sxtdhmj.CanHuWithLaiZi(handCards, laizi)
		if !ok || dianshu != 10 {
			sxtdhmj.ShowHandCards(handCards)
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
		if !sxtdhmj.IsValidHandCards(handCards1) {
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

						//for i := 1; i <= sxtdhmj.MaxCard; i++ {
						laizi := []int{0x01}
						//laizi = append(laizi, int(i))
						ok, _ := sxtdhmj.CanHuWithLaiZi(handCards4, laizi)
						if !ok {
							sxtdhmj.ShowHandCards(handCards4)
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

func TestGenAllPossible(t *testing.T) {
	count = 0
	var m1 map[int]*HuArray
	m1 = make(map[int]*HuArray)

	handCards := []int{}
	handCards = append(handCards, 0)
	handCards = append(handCards, 0)
	GenAllPossibleAddPuZiToHandCards(t, m1, handCards, 0, 4)

	file, _ := os.OpenFile("allPossible.log", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	buf := bufio.NewWriter(file)
	fmt.Fprintf(buf, "count = %d\n", count)
	for k, v := range m1 {
		for _, vj := range v.arr {
			fmt.Fprintf(buf, "%d=%v, %v, dianshu=%d, laizi=%d\n", k, vj.handCards, vj.hu, vj.dianshu, vj.laizi)
		}
	}
	buf.Flush()
}

func GenAllPossibleAddPuZiToHandCards(t *testing.T, m1 map[int]*HuArray, handCards []int, jiang int, level int) {
	if level <= 0 {

		count++

		key, distinctCards := calcHandCardsKey(handCards)
		if _, ok := m1[key]; !ok {

			arr := &HuArray{}

			for laiziCard := range distinctCards {
				laizi := []int{}
				laizi = append(laizi, laiziCard)
				ok, dianshu := sxtdhmj.CanHuWithLaiZi(handCards, laizi)

				ret := &HuRet{}
				ret.handCards = make([]int, len(handCards))
				copy(ret.handCards, handCards)
				ret.hu = ok
				ret.dianshu = dianshu
				ret.laizi = laiziCard
				arr.arr = append(arr.arr, ret)
			}
			m1[key] = arr
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
		if !sxtdhmj.IsValidHandCards(handCards1) {
			continue
		}

		GenAllPossibleAddPuZiToHandCards(t, m1, handCards1, jiang, level-1)
	}
}

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

// func getMaxOneCard(slots [sxtdhmj.TILEMAX]int) (bool, int) {
// 	for i := sxtdhmj.TON; i <= sxtdhmj.CHU; i++ {
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

// 		slots := sxtdhmj.GenSlots(handCards)
// 		ret, maxOneCard := getMaxOneCard(slots)
// 		if ret {
// 			laizi := []int{}
// 			laizi = append(laizi, maxOneCard)
// 			ok, dianshu := sxtdhmj.CanHuWithLaiZi(handCards, laizi)
// 			if !ok || dianshu != sxtdhmj.DianShuTable[maxOneCard] {
// 				fmt.Println("slots=", slots, ", maxOneCard=", maxOneCard,
// 					"sxtdhmj.DianShuTable[maxOneCard]=", sxtdhmj.DianShuTable[maxOneCard], "dianshu=", dianshu, ", ")
// 				sxtdhmj.ShowHandCards(handCards)
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
// 		if !sxtdhmj.IsValidHandCards(handCards1) {
// 			continue
// 		}

// 		ret := OneLaiZiAddPuZiToHandCards(t, handCards1, jiang, level-1)
// 		if !ret {
// 			return false
// 		}
// 	}

// 	return true
// }
