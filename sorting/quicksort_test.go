package sorting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotReallyQuicksort(t *testing.T) {
	lst := []int{33, 15, 7, 59, 1, 6, 57, 44, 8, 36, 6, 47}
	expected := []int{1, 6, 6, 7, 8, 15, 33, 36, 44, 47, 57, 59}
	require.Equal(t, expected, NotReallyQuicksort(lst))
}

func TestQuicksort2(t *testing.T) {
	lst := []int{33, 15, 7, 59, 1, 6, 57, 44, 8, 36, 6, 47}
	expected := []int{1, 6, 6, 7, 8, 15, 33, 36, 44, 47, 57, 59}
	require.Equal(t, expected, Quicksort2(lst))
}

func TestQuicksort4(t *testing.T) {
	lst := []int{33, 15, 7, 59, 1, 6, 57, 44, 8, 36, 6, 47}
	expected := []int{1, 6, 6, 7, 8, 15, 33, 36, 44, 47, 57, 59}
	out := qsort(lst)
	require.Equal(t, expected, out)
}

func TestQuicksort5(t *testing.T) {
	lst := []int{-1, 2, -1}
	expected := []int{-1, -1, 2}
	out := qsort(lst)
	require.Equal(t, expected, out)
}
