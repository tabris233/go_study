package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint // uint 在32位机器是uint32 在64位机器是uint64
}

const (
	bitNum = 32 << (^uint(0) >> 63) // bit number 用来检测当前机器是32位机器还是64位机器.
)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bitNum, uint(x%bitNum)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bitNum, uint(x%bitNum)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitNum*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len  return the number of elements
func (s *IntSet) Len() int {
	len := 0

	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}

	return len
}

// Remove  remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/bitNum, uint(x%bitNum)
	s.words[word] &^= (1 << bit)
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = *new([]uint)
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var ans IntSet

	for _, sword := range s.words {
		ans.words = append(ans.words, sword)
	}

	return &ans
}

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

// IntersectWith sets s to the Intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
}

// DifferenceWith sets s to the Difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

// SymmetricDifference sets s to the Difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	var ans IntSet
	for i, tword := range t.words {
		if i < len(s.words) {
			ans.words = append(ans.words, (s.words[i]|tword)&^(s.words[i]&tword))
		} else {
			ans.words = append(ans.words, s.words[i])
		}
	}
	return &ans
}

// Elems  for range
func (s *IntSet) Elems() []uint {
	var ans []uint
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				ans = append(ans, uint(i*bitNum+j))
			}
		}
	}
	return ans
}
