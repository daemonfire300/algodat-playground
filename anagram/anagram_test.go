package anagram

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnagram(t *testing.T) {
	anagrams := []string{
		"act",
		"cat",
		"tree",
		"race",
		"care",
		"acre",
		"bee",
	}
	expected := [][]string{
		{"act", "cat"},
		{"tree"},
		{"race", "care", "acre"},
		{"bee"},
	}
	res := anagram(anagrams)
	require.Equal(t, expected, res)
}
