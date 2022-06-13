package anagram

import (
	"fmt"
	"sync"

	"github.com/daemonfire300/algodat-playground/sorting"
)

// This function should illustrate how we could span the initial function
// over the network or threads or green threads.
// It will not outperform the other solutions for small data-sets.
//
// It still has some global state, i.e., the wordIndex, therefore you cannot truly
// or simply distribute the workload with the conceptual implementation below.
// However, it gives many hints at how a partial distribution can happen and some tricks
// like for instance using buffered channels so that the reading of a data source which might happen
// faster than the hashing of the words and updating of the wordIndex can continue for some time
// without blocking the process by filling the buffer.
func anagramStreaming(in []string) [][]string {
	var wordIndex = make(map[string][]int) // maps hash to indices that match that hash
	// n in this context refers to len(in)
	out := make([][]string, 0)

	wg := &sync.WaitGroup{}
	quit := make(chan struct{}, 1)
	wordQueue := make(chan struct { // inline definition of struct because this is sample code
		word string
		idx  int
	}, 32) // non-blocking channel that can be filled with up to 32 items
	wordIndexQueue := make(chan struct {
		hash string
		idx  int
	}, 32)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Quit (word index processing)")
			case item := <-wordIndexQueue:
				if _, ok := wordIndex[item.hash]; !ok {
					wordIndex[item.hash] = make([]int, 0)
				}
				wordIndex[item.hash] = append(wordIndex[item.hash], item.idx)
				wg.Done()
			}
		}
	}()
	for i := 0; i < 4; i++ { // let's start some routines, 4 at most
		go func() {
			for {
				select {
				case <-quit:
					fmt.Println("Quit (processing)")
					return
				case item := <-wordQueue:
					h := hash(string(sorting.QSortStr(copyStringToRunSlice(item.word))))
					wg.Add(1)
					wordIndexQueue <- struct {
						hash string
						idx  int
					}{
						hash: h,
						idx:  item.idx,
					}
					wg.Done()
				}
			}
		}()
	}

	// Populate wordIndex, O(n)
	// Keep track of order (ensure stable algorithm)
	// Memory O(n+u) where u are unique hashes
	for i, word := range in {
		wg.Add(1)
		wordQueue <- struct {
			word string
			idx  int
		}{
			word: word,
			idx:  i,
		}
	}
	wg.Wait()

	// Turn index into word list
	for _, indices := range wordIndex {
		tmp := make([]string, len(indices))
		for j, index := range indices {
			tmp[j] = in[index]
		}
		out = append(out, tmp)
	}
	return out
}
