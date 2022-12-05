package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProduct(t *testing.T) {
	var data = []struct {
		array  []int
		answer int
	}{
		{
			array:  []int{1, 8, 5, 4, 10, 2, 6, 1, 1, 1, 9},
			answer: 72,
		},
		{
			array:  []int{3, 7},
			answer: 12,
		},
	}

	for _, v := range data {
		answer := maxProduct(v.array)
		require.EqualValues(t, answer, v.answer)
	}

}

func TestTotalMoney(t *testing.T) {
	v := totalMoney(10)
	require.EqualValues(t, v, 37)
}

func TestMoveZeroes(t *testing.T) {
	var data = []struct {
		array  []int
		answer []int
	}{
		{
			array:  []int{0, 1, 0},
			answer: []int{1, 0, 0},
		},
		{
			array:  []int{0, 1, 0, 3, 12},
			answer: []int{1, 3, 12, 0, 0},
		},

		{
			array:  []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0},
			answer: []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0},
		},
	}

	for _, v := range data {
		moveZeroes(v.array)
		require.EqualValues(t, reflect.DeepEqual(v.answer, v.array), true, v.array)
	}
}

func TestMaxProfit(t *testing.T) {
	s := maxProfit([]int{7, 1, 5, 3, 6, 4})
	require.EqualValues(t, s, 5)
}

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("()"))
}

func TestFirstUniqChar(t *testing.T) {
	i := firstUniqChar("leetocode")
	require.EqualValues(t, i, "0")

}
