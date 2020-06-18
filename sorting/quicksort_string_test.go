package sorting

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuicksortStr(t *testing.T) {
	lst := []rune(strings.Join([]string{"c", "a", "b", "o", "p", "k"}, ""))
	fmt.Println(lst)
	expected := []rune(strings.Join([]string{"a", "b", "c", "k", "o", "p"}, ""))
	out := QSortStr(lst)
	fmt.Println(lst, out, expected)
	require.Equal(t, expected, out)
}
