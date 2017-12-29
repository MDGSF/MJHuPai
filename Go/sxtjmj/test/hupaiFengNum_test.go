package main

import (
	"testing"

	"github.com/MDGSF/MJHuPai/Go/sxtjmj"
)

func TestFengNum1(t *testing.T) {
	handCards := []int{27, 27, 28, 28, 29, 29, 30, 30, 31, 31, 31, 31, 32, 33}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, handCards[0], true, true, false)
	if !ok || fengNum != 3 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum2(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		27, 27, 28, 29, 30}
	ok, fengNum := sxtjmj.CanHu(handCards, 2, 16, true, true, false)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum3(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16,
		27, 27, 27, 28, 29, 30}
	ok, fengNum := sxtjmj.CanHu(handCards, 2, 27, true, true, false)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum4(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16,
		27, 27, 28, 28, 29, 30}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 28, true, true, false)
	if !ok || fengNum != 2 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum5(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16,
		27, 27, 28, 29, 29, 30}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 29, true, true, false)
	if !ok || fengNum != 2 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum6(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16,
		27, 27, 28, 29, 30, 30}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 30, true, true, false)
	if !ok || fengNum != 2 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum7(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		17, 17,
		27, 28, 29}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 29, true, true, false)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum8(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		17, 17,
		27, 28, 30}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 30, true, true, false)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum9(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		27, 27, 27, 28, 28}
	ok, fengNum := sxtjmj.CanHu(handCards, 2, 28, true, true, false)
	if !ok || fengNum != 0 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum10(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		27, 27, 27, 28, 28}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 28, true, true, false)
	if !ok || fengNum != 0 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum11(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		27, 27, 27, 28, 29}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 29, true, true, false)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum12(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		27, 27, 27, 28, 30}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 30, true, true, false)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum13(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16,
		33, 33, 32, 32, 31, 31}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 31, true, true, false)
	if !ok || fengNum != 2 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum14(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		17, 17,
		33, 32, 31}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 31, true, true, false)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum15(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		33, 33, 33, 32, 32}
	ok, fengNum := sxtjmj.CanHu(handCards, 2, 32, true, true, false)
	if !ok || fengNum != 0 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum16(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		33, 33, 33, 32, 32}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 32, true, true, false)
	if !ok || fengNum != 0 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum17(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		33, 33, 33, 32, 31}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 31, true, true, false)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum18(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		13, 13,
		16, 16,
		33, 33, 32, 32}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 13, false, false, true)
	if !ok || fengNum != 2 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum19(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		13, 4,
		16, 16,
		33, 33, 32, 32}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 4, false, false, true)
	if !ok || fengNum != 2 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum20(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		13, 22,
		16, 16,
		33, 33, 32, 32}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 22, false, false, true)
	if !ok || fengNum != 2 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum21(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		13,
		16, 16, 16,
		17, 17,
		33, 32}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 32, false, false, true)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum22(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		33, 33, 33, 32, 32}
	ok, fengNum := sxtjmj.CanHu(handCards, 2, 32, false, false, true)
	if !ok || fengNum != 0 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum23(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		33, 33, 33, 32, 32}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 32, false, false, true)
	if !ok || fengNum != 0 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum24(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		33, 33, 33, 32, 4}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 4, false, false, true)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum25(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		33, 33, 33, 32, 13}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 13, false, false, true)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}

func TestFengNum26(t *testing.T) {
	handCards := []int{
		2, 3, 4,
		6, 6, 6,
		16, 16, 16,
		33, 33, 33, 32, 22}
	ok, fengNum := sxtjmj.CanHu(handCards, 1, 22, false, false, true)
	if !ok || fengNum != 1 {
		sxtjmj.ShowHandCards(handCards)
		t.Error("CanHu failed.")
	}
}
