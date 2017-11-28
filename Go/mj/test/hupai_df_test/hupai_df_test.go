package main

import (
	"testing"

	"github.com/MDGSF/MJHuPai/Go/mj"
)

var (
	mjslots = [][]int{
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4, 0, 0},
	}

	testCount = 0
)

// TestWinable 测试胡牌判断
func TestWinable(t *testing.T) {
	var hai = []int{
		0, 1, 2,
		0, 0, 0,
		9, 9, 9,
		11, 11, 11,
		12, 12}

	if !mj.IsWinableTest(hai) {
		t.Errorf("winable test failed")
	}
}

// 测试4个面子牌组+一对将
// 注意有大量重复，如果希望过滤重复则需要用一个哈希表来存储已枚举的组合
func Test4Melds(t *testing.T) {
	testCount = 0
	_arr := make([]int, 14)
	for j := 0; j < 4; j++ {
		for i := 0; i < 9; i++ {
			if mjslots[j][i] < 2 {
				// System.Diagnostics.Debug.Assert(j == 3 && (i == 7 || i == 8));
				continue
			}

			// remove head elements
			mjslots[j][i] = mjslots[j][i] - 2
			_arr[0] = 9*j + i
			_arr[1] = 9*j + i

			if !mj.IsWinableTest(_arr[:2]) {
				t.Errorf("Test4Melds test failed")
			}

			Four3x(4, _arr, t)

			//_set.Clear()
			// restore head elements
			mjslots[j][i] = mjslots[j][i] + 2
		}
	}

	t.Logf("Test4Melds test count:%d\n", testCount)
}

func Four3x(n int, _arr []int, t *testing.T) {

	if !mj.IsWinableTest(_arr[:(14 - n*3)]) {
		t.Errorf("Test4Melds test failed")
	}

	// 顺子
	for j := 0; j < 3; j++ {
		for i := 0; i < 7; i++ {
			if mjslots[j][i] > 0 && mjslots[j][i+1] > 0 && mjslots[j][i+2] > 0 {
				_arr[2+(4-n)*3] = 9*j + i
				_arr[2+(4-n)*3+1] = 9*j + i + 1
				_arr[2+(4-n)*3+2] = 9*j + i + 2

				if n-1 > 0 {
					// remove 3x elements
					mjslots[j][i]--
					mjslots[j][i+1]--
					mjslots[j][i+2]--

					Four3x(n-1, _arr, t)
					// restore 3x elements
					mjslots[j][i]++
					mjslots[j][i+1]++
					mjslots[j][i+2]++
				} else {
					testCount++

					//Array.Sort(_arr);
					//System.Diagnostics.Debug.Assert(jap.CanHu(_arr));
					//SaveCalcKey(_arr);
					//_set.Add(new SetItem(_arr));
					//Keep6(_arr);
					if !mj.IsWinableTest(_arr) {
						t.Errorf("Test4Melds test failed")
					}
				}
			}
		}
	}

	// 刻子
	for j := 0; j < 4; j++ {
		for i := 0; i < 9; i++ {
			if mjslots[j][i] > 2 {
				_arr[2+(4-n)*3] = 9*j + i
				_arr[2+(4-n)*3+1] = 9*j + i
				_arr[2+(4-n)*3+2] = 9*j + i
				if n-1 > 0 {
					// remove 3x elements
					mjslots[j][i] -= 3

					Four3x(n-1, _arr, t)

					// restore 3x elements
					mjslots[j][i] += 3
				} else {
					testCount++
					//Array.Sort(_arr);
					//System.Diagnostics.Debug.Assert(jap.CanHu(_arr));
					//SaveCalcKey(_arr);
					//_set.Add(new SetItem(_arr));
					//Keep6(_arr);
					if !mj.IsWinableTest(_arr) {
						t.Errorf("Test4Melds test failed")
					}
				}

			}
		}
	}
}
