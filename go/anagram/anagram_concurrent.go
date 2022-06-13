package anagram

import (
	"sync"

	"github.com/daemonfire300/algodat-playground/sorting"
)

func anagramConcurrent(in []string) [][]string {
	var wordIndex = make(map[string][]int) // maps hash to indices that match that hash
	// n in this context refers to len(in)
	out := make([][]string, 0)

	wg := &sync.WaitGroup{}
	wordIndexMutex := &sync.Mutex{}
	// Populate wordIndex, O(n)
	// Keep track of order (ensure stable algorithm)
	// Memory O(n+u) where u are unique hashes
	order := make([]string, 0)
	for i, word := range in {
		wg.Add(1)
		go func(idx int, w string) { // TODO (JF) add configurable limit in order to prevent resource exhaustion by spawning to many go
			// routines
			defer wg.Done()
			h := hash(string(sorting.QSortStr(copyStringToRunSlice(w))))
			wordIndexMutex.Lock()
			defer wordIndexMutex.Unlock()
			if _, ok := wordIndex[h]; !ok {
				wordIndex[h] = make([]int, 0)
				order = append(order, h)
			}
			wordIndex[h] = append(wordIndex[h], idx)
		}(i, word)
	}
	wg.Wait()

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
