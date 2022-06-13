package trees

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertion(t *testing.T) {
	values := []int{33, 15, 7, 59, 1, 6, 57, 44, 8, 36, 6, 47}
	tree := &Tree{}
	for _, val := range values {
		tree.Insert(val)
	}
}

func TestFind(t *testing.T) {
	values := []int{33, 15, 7, 59, 1, 6, 57, 44, 8, 36, 6, 47}
	tree := &Tree{}
	for _, val := range values {
		tree.Insert(val)
	}
	require.True(t, tree.Find(8))
	require.False(t, tree.Find(32))
}
