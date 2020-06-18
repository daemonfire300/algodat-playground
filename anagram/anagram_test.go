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

func BenchmarkAnagram(b *testing.B) {
	/*
		BenchmarkAnagram
		BenchmarkAnagram-8   	  180724	      5961 ns/op	    1212 B/op	      28 allocs/op
		BenchmarkAnagram
		BenchmarkAnagram-8   	  622585	      1758 ns/op	     489 B/op	      14 allocs/op
	*/
	anagrams := []string{
		"act",
		"cat",
		"tree",
		"race",
		"care",
		"acre",
		"bee",
	}
	for i := 0; i < b.N; i++ {
		anagram(anagrams)
	}
}
