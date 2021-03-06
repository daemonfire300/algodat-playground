package anagram

import (
	"crypto/sha1"
	"fmt"
	"sync"
	"unicode/utf8"

	"github.com/daemonfire300/algodat-playground/sorting"
)

func anagram(in []string) [][]string {
	var wordIndex = make(map[string][]int) // maps hash to indices that match that hash
	// n in this context refers to len(in)
	out := make([][]string, 0)

	// Populate wordIndex, O(n)
	// Keep track of order (ensure stable algorithm)
	// Memory O(n+u) where u are unique hashes
	order := make([]string, 0)
	for i, word := range in {
		h := hash(string(sorting.QSortStr(copyStringToRunSlice(word))))
		if _, ok := wordIndex[h]; !ok {
			wordIndex[h] = make([]int, 0)
			order = append(order, h)
		}
		wordIndex[h] = append(wordIndex[h], i)
	}

	// Turn index into word list
	for i, h := range order {
		indices := wordIndex[h]
		out = append(out, make([]string, len(indices)))
		for j, index := range indices {
			out[i][j] = in[index]
		}
	}
	return out
}

func copyStringToRunSlice(word string) []rune {
	runes := make([]rune, utf8.RuneCountInString(word))
	for i, r := range word {
		runes[i] = r
	}
	return runes
}

func hash(word string) string {
	return hashWithCacheConcurrencySafe(word)
}

var hashCache = make(map[string]string)

func hashWithCache(word string) string {
	if h, ok := hashCache[word]; ok {
		return h
	}
	hashCache[word] = fmt.Sprintf("%x", sha1.Sum([]byte(word)))
	return hashCache[word]
}

var hashMutex = &sync.Mutex{}

func hashWithCacheConcurrencySafe(word string) string {
	hashMutex.Lock()
	defer hashMutex.Unlock() // Could also use sync.Map instead, but then
	// we deal with interface values instead of a typed map.
	// And for range becames a map.Range(func()...)
	if h, ok := hashCache[word]; ok {
		return h
	}
	hashCache[word] = fmt.Sprintf("%x", sha1.Sum([]byte(word)))
	return hashCache[word]
}
