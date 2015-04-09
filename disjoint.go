/*
Package disjoint implements a disjoint-set data structure.

A disjoint-set data structure keeps track of nonoverlapping partitions of a
collection of data elements.  Initially, each data element belongs to its own,
singleton, set.  The following operations can then be performed on these sets:

• Union (called Merge in this package) merges two sets into a single set
containing the union of their elements.

• Find (called Representative in this package) returns an arbitrary set
belonging to a union of sets.

The critical feature is that the same representative is returned for all sets
in the same union.  The implication is that two sets A and B belong to the same
union if and only if A.Representative() == B.Representative().

Disjoint sets are often used in graph algorithms, for example to find a minimal
spanning tree for a graph or to determine if adding a given edge to a graph
would create a cycle.

Both Merge and Representative run in amortized near-constant time.  See
http://en.wikipedia.org/wiki/Disjoint_set_forest for more information.
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

// Representative returns an arbitrary Set belonging to a union of Sets.  The
// important feature is that it returns the same Set for any Set in the union
// so it can be used to test if two Sets are in the same union.
func (s *Set) Representative() *Set {
	for s.parent != s {
		s.parent = s.parent.parent
		s = s.parent
	}
	return s
}

// Merge puts two Sets into the same union.  Thenceforth, Representative will
// return the same Set when given either of the two Sets.
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
