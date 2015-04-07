/*
Package disjoint implements a disjoint-set data structure.

A disjoint-set data structure keeps track of nonoverlapping partitions
of a collection of data elements.  Initially, each data element
belongs to its own, singleton, set.  The following operations can then
be performed on these sets:

* Union (called Merge in this package) merges two sets into a single
set containing the union of their elements.

* Find (called Representative in this package) returns an arbitrary
set from a union of sets.

Importantly, the same representative is returned for all sets in the
same union.  The implication is that two sets A and B belong to the
same union if and only if A.Representative() == B.Representative().

Disjoint sets are useful for determining if adding a given edge to a
graph would create a cycle.  Personally, I use them for maze
generation: repeatedly knock down walls in a maze between two sets of
rooms that are not part of the same union, then merge those rooms into
a single union.

Both Merge and Representative run in amortized near-constant time.
See http://en.wikipedia.org/wiki/Disjoint_set_forest for more information.
*/
package disjoint

// A Set represents a disjoint set in a disjoint-set forest.
type Set struct {
	parent *Set // Parent set
	rank   int  // Rank (approximate depth) of the set's subtree
}

// Singleton returns a one-element Set.
func Singleton() *Set {
	s := &Set{rank: 0}
	s.parent = s
	return s
}

// Representative returns an arbitrary Set from a union of Sets.  The important
// feature is that it returns the same Set for any Set in the union so it can
// be used to test if two Sets are in the same union.
func (s *Set) Representative() *Set {
	for s.parent != s {
		s.parent = s.parent.parent
		s = s.parent
	}
	return s
}

// Merge creates a union of a Set with another Set.  When this method returns,
// Representative will return the same Set for each of the two Sets.
func (s *Set) Merge(t *Set) {
	// Ensure the two Sets aren't already part of the same union.
	sRoot := s.Representative()
	tRoot := t.Representative()
	if sRoot == tRoot {
		return
	}

	// Merge the two sets by making the shorter tree point to the root of
	// the larger tree.
	switch {
	case sRoot.rank < tRoot.rank:
		sRoot.parent = tRoot
	case sRoot.rank > tRoot.rank:
		tRoot.parent = sRoot
	default:
		tRoot.parent = sRoot
		sRoot.rank++
	}
}
