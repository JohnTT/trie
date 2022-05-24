// Package trie implements the trie data structure using runes.
// https://en.wikipedia.org/wiki/Trie
//
// This can be tested on LeetCode problem below:
// https://leetcode.com/problems/implement-trie-prefix-tree/
package trie

// Trie defines the underlying data structure.
type Trie struct {
	Children   map[rune]*Trie
	IsTerminal bool
}

// Constructor is an alias for Make() to satisfy LeetCode API.
func Constructor() Trie {
	return Make()
}

// Make returns a Trie.
func Make() Trie {
	return Trie{make(map[rune]*Trie), false}
}

// New returns a pointer to a Trie.
func New() *Trie {
	return &Trie{make(map[rune]*Trie), false}
}

// Insert puts a UTF-8 string into the Trie.
func (this *Trie) Insert(word string) {
	cur := this
	for _, r := range word {
		// Create the next node if it doesn't exist.
		if _, ok := cur.Children[r]; !ok {
			cur.Children[r] = New()
		}

		// Move to the next node in the Trie.
		cur = cur.Children[r]
	}

	// Mark the terminal node.
	cur.IsTerminal = true
}

// Search checks if a UTF-8 string exists in the Trie.
func (this *Trie) Search(word string) bool {
	cur := this
	for _, r := range word {
		// Return False if the next node if it doesn't exist.
		if _, ok := cur.Children[r]; !ok {
			return false
		}

		// Move to the next node in the Trie.
		cur = cur.Children[r]
	}

	return cur.IsTerminal
}

// StartsWith checks if a UTF-8 prefix exists in the Trie.
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, r := range prefix {
		// Return False if the next node if it doesn't exist.
		if _, ok := cur.Children[r]; !ok {
			return false
		}

		// Move to the next node in the Trie.
		cur = cur.Children[r]
	}

	return true
}
