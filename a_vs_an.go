/*
Package indart provides the Find function to identify the correct indefinite article for a word.
*/
package indart

import (
	"strings"
)

type dataNode struct {
	Article article
	Next    map[rune]*dataNode
}

type article bool

const (
	a  = article(true)
	an = article(false)
)

// Find returns the correct indefinite article for word. It returns either "a" or "an".
func Find(word string) string {
	n := dataTree()
	w := strings.TrimLeft(word, ignoredRuneInPrefixSet)
	for _, c := range w {
		nn, ok := n.Next[c]
		if !ok {
			return a2s(n.Article)
		}
		n = nn
	}
	return a2s(dataRoot.Article)
}
