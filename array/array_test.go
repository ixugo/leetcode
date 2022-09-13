package array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchInsert(t *testing.T) {
	v := searchInsert([]int{1, 3, 5, 6}, 2)
	require.EqualValues(t, v, 1)
}

func TestLongestCommonPrefix(t *testing.T) {
	v := longestCommonPrefix([]string{"flower", "flower", "flower", "flower"})
	require.EqualValues(t, v, "fl")
}
