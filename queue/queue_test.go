package leetcode

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDequeue(t *testing.T) {
	ans := maxSlidingWindow([]int{7, 2, 4}, 2)
	expected := []int{7, 4}
	require.True(t, reflect.DeepEqual(ans, expected))
}
