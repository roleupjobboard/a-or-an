package indefart

import (
	"strings"
)

// It is the implementation in Go of:
// https://eamonnerbonne.github.io/a-vs-an/AvsAnDemo/
// https://github.com/EamonNerbonne/a-vs-an/

type dataNode struct {
	Article article
	Next    map[rune]*dataNode
}

type article bool

const (
	a  = article(true)
	an = article(false)
)

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
