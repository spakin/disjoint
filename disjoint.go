/*
Package disjoint implements a disjoint-set data structure.

A disjoint-set—also called union-find—data structure keeps track of
nonoverlapping partitions of a collection of data elements.  Initially, each
data element belongs to its own, singleton, set.  The following operations can
then be performed on these sets:

• Union merges two sets into a single set containing the union of their
elements.

• Find returns an arbitrary element from a set.

The critical feature is that Find returns the same element when given any
element in a set.  The implication is that two sets A and B belong to the same
union if and only if A.Find() == B.Find().

Disjoint sets are often used in graph algorithms, for example to find a minimal
spanning tree for a graph or to determine if adding a given edge to a graph
would create a cycle.

Both Union and Find run in amortized near-constant time.  See
http://en.wikipedia.org/wiki/Disjoint_set_forest for more information.
*/
package disjoint

// An Element represents a value that can be associated with a set of other
// values.
type Element struct {
	parent *Element // Parent element
	rank   int      // Rank (approximate depth) of the subtree with this element as root
}

// NewElement returns an element that is not associated with any other element.
func NewElement() *Element {
	s := &Element{rank: 0}
	s.parent = s
	return s
}

// Find returns an arbitrary Element belonging to a union of Elements.  The
// important feature is that it returns the same Element for any Element in the
// union so it can be used to test if two Elements are in the same union.
func (s *Element) Find() *Element {
	for s.parent != s {
		s.parent = s.parent.parent
		s = s.parent
	}
	return s
}

// Union associates two Elements—and by transitivity, all Elements with which
// they are associated.  Thenceforth, Find will return the same Element when
// given any Element in the union.
func (s *Element) Union(t *Element) {
	// Ensure the two Elements aren't already part of the same union.
	sRoot := s.Find()
	tRoot := t.Find()
	if sRoot == tRoot {
		return
	}

	// Create a union by making the shorter tree point to the root of the
	// larger tree.
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
