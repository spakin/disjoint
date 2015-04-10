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
element in a set.  The implication is that two elements A and B belong to the
same set if and only if A.Find() == B.Find().

Disjoint sets are more limited in functionality than conventional sets.  They
support only set union, not set intersection, set difference, or any other set
operation.  They don't allow an element to reside in more than one set.  They
don't even provide a way to enumerate the elements in a given set.  What makes
them useful, though, is that they're extremely fast, especially for large sets;
both Union and Find run in amortized near-constant time.  See
http://en.wikipedia.org/wiki/Disjoint_set_forest for more information.

Disjoint sets are often used in graph algorithms, for example to find a minimal
spanning tree for a graph or to determine if adding a given edge to a graph
would create a cycle.
*/
package disjoint

// An Element represents a single element of a set.
//
// The Data field lets a program store arbitrary data within an Element.  This
// can be used, for example, to keep track of the total number of elements in
// the associated set, the set's maximum-valued element, a map of attributes
// associated with the set's elements, or even a map or slice of all elements
// in the set.  (That last possibility would associate a linear-time cost with
// each Union operation but would not affect Find's near-constant running
// time.)
type Element struct {
	parent *Element    // Parent element
	rank   int         // Rank (approximate depth) of the subtree with this element as root
	Data   interface{} // Arbitrary user-provided payload
}

// NewElement returns an element belonging to a singleton set.
func NewElement() *Element {
	s := &Element{rank: 0}
	s.parent = s
	return s
}

// Find returns an arbitrary Element from a set containing a given Element.
// The important feature is that it returns the same value for every element in
// the set.  Consequently, it can be used to test if two Elements belong to the
// same set.
func (s *Element) Find() *Element {
	for s.parent != s {
		s.parent = s.parent.parent
		s = s.parent
	}
	return s
}

// Given an element from each of two sets, Union performs a destructive union
// of those sets.  Afterwards, the original sets no longer exist as separate
// entities.
//
// Note that, perhaps surprisingly, the package does not expose an explicit Set
// data type.  Sets exist only implicitly based on the sequence of Union
// operations performed on their elements.
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
