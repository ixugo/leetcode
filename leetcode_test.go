package leetcode

import (
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
