// An IntSet is a set of small non-negative integers.
// Its zero value represents an empty set

package main

// IntSet Structure
type IntSet struct {
	words []uint64
}

// Contains reports whether the set has a non-negative value x.
func (s *IntSet) has(x int) bool {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) 
	}
	s.words[word] |= 1 << bit
	

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}