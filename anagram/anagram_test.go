package anagram

import (
	"fmt"
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
		BenchmarkAnagram (baseline)
		BenchmarkAnagram-8   	  180724	      5961 ns/op	    1212 B/op	      28 allocs/op
		BenchmarkAnagram (with cache)
		BenchmarkAnagram-8   	  622585	      1758 ns/op	     489 B/op	      14 allocs/op
		BenchmarkAnagram (with concurrency-safe cache)
		BenchmarkAnagram-8   	  352878	      2970 ns/op	     664 B/op	      31 allocs/op
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

func TestAnagramConcurrent(t *testing.T) {
	anagrams := []string{
		"act",
		"cat",
		"tree",
		"race",
		"care",
		"acre",
		"bee",
	}
	res := anagramConcurrent(anagrams)
	// Algorithm is not stable
	require.Len(t, res, 4)
}

func BenchmarkAnagramConcurrent(b *testing.B) {
	/*
		BenchmarkAnagram (baseline)
		BenchmarkAnagram-8   	  180724	      5961 ns/op	    1212 B/op	      28 allocs/op
		BenchmarkAnagram (with cache)
		BenchmarkAnagram-8   	  622585	      1758 ns/op	     489 B/op	      14 allocs/op
		BenchmarkAnagram (with concurrency-safe cache)
		BenchmarkAnagram-8   	  352878	      2970 ns/op	     664 B/op	      31 allocs/op
		BenchmarkAnagramConcurrent (concurrent, with locks)
		BenchmarkAnagramConcurrent-8   	  140996	      7714 ns/op	    1120 B/op	      36 allocs/op
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
		anagramConcurrent(anagrams)
	}
}

func TestAnagramStreaming(t *testing.T) {
	anagrams := []string{
		"act",
		"cat",
		"tree",
		"race",
		"care",
		"acre",
		"bee",
	}
	res := anagramStreaming(anagrams)
	fmt.Println(res)
	// Algorithm is not stable
	require.Len(t, res, 4)
}

func BenchmarkAnagramStreaming(b *testing.B) {
	/*
		BenchmarkAnagram (baseline)
		BenchmarkAnagram-8   	  180724	      5961 ns/op	    1212 B/op	      28 allocs/op
		BenchmarkAnagram (with cache)
		BenchmarkAnagram-8   	  622585	      1758 ns/op	     489 B/op	      14 allocs/op
		BenchmarkAnagram (with concurrency-safe cache)
		BenchmarkAnagram-8   	  352878	      2970 ns/op	     664 B/op	      31 allocs/op
		BenchmarkAnagramConcurrent (concurrent, but not stable)
		BenchmarkAnagramConcurrent-8   	  410637	      2928 ns/op	     664 B/op	      31 allocs/op
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
		anagramStreaming(anagrams)
	}
}
