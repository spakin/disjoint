// Test disjoint-set forests.

package disjoint

import (
	"math/rand"
	"testing"
)

// TestEvenOdd puts even numbers in one union and odd numbers in other, which
// is easy to test.
func TestEvenOdd(t *testing.T) {
	// Create a bunch of singleton sets.
	const N = 1000
	sets := make([]*Element, N)
	for i := 0; i < N; i++ {
		sets[i] = NewElement()
	}

	// Merge each even number with its predecessor and each odd number with
	// its predecessor.
	for i := 2; i < N; i += 2 {
		Union(sets[i], sets[i-2])
	}
	for i := 3; i < N; i += 2 {
		Union(sets[i], sets[i-2])
	}

	// Ensure that even numbers are in the same union as other even numbers
	// and odd numbers are in the same union as other oddn numbers.
	for i := 0; i < N*3; i++ {
		s1 := rand.Intn(N)
		s2 := rand.Intn(N)
		sameMod2 := s1%2 == s2%2
		sameRep := sets[s1].Find() == sets[s2].Find()
		if sameMod2 != sameRep {
			t.Fatalf("Should %d and %d lie in the same set?  The package incorrectly says %v.",
				s1, s2, sameRep)
		}
	}
}

// TestExtraData tests disjoint sets that carry auxiliary data.
func TestExtraData(t *testing.T) {
	// Create six singleton sets, each maintaining integer data.
	sets := make([]*Element, 6)
	for i := 0; i < 6; i++ {
		sets[i] = NewElement()
		sets[i].Data = (6 - i) * 100
	}

	// Define a function that merges two sets while retaining the
	// larger data value.
	merge := func(e1, e2 *Element) *Element {
		v1 := e1.Data.(int)
		vm := e2.Data.(int)
		if v1 > vm {
			vm = v1
		}
		Union(e1, e2)
		e := e1.Find()
		e.Data = vm
		return e
	}

	// Merge sets 1, 0, and 4 in that order.
	e1 := merge(merge(sets[1], sets[0]), sets[4])
	if e1.Data.(int) != 600 {
		t.Fatalf("Expected 600 but observed %d", e1.Data.(int))
	}

	// Merge sets 2, 3, and 5 in that order.
	e2 := merge(merge(sets[2], sets[3]), sets[5])
	if e2.Data.(int) != 400 {
		t.Fatalf("Expected 400 but observed %d", e2.Data.(int))
	}

	// Merge the preceding two sets.
	e12 := merge(e1, e2)
	if e12.Data.(int) != 600 {
		t.Fatalf("Expected 600 but observed %d", e12.Data.(int))
	}
}

// createElements returns a slice of Elements.
func createElements(n int) []*Element {
	elts := make([]*Element, n)
	for i := range elts {
		elts[i] = NewElement()
	}
	return elts
}

// selectIndexes returns a list of N pairs of indexes into a slice of length N.
func selectIndexes(n int) [][2]int {
	idxes := make([][2]int, n)
	if n < 2 {
		return idxes
	}
	for i := range idxes {
		idxes[i][0] = i
		if i == 0 {
			idxes[i][1] = rand.Intn(n)
		} else {
			idxes[i][1] = rand.Intn(i)
		}
	}
	return idxes
}

// pairwiseUnions repeatedly takes pairwise unions of a number of sets.
func pairwiseUnions(elts []*Element, idxes [][2]int) {
	for _, idx := range idxes {
		e1 := elts[idx[0]]
		e2 := elts[idx[1]]
		Union(e1, e2)
	}
}

// BenchmarkUnion measures the time to perform N union operations.
func BenchmarkUnion(b *testing.B) {
	b.StopTimer()
	elts := createElements(b.N)
	idxes := selectIndexes(b.N)
	b.StartTimer()
	pairwiseUnions(elts, idxes)
}

// BenchmarkUnionFind measures the time to perform N union operations followed
// by N find operations.  Find operations are so fast, we run out of memory
// trying to allocate enough Elements as to stress Find's performance.
func BenchmarkUnionFind(b *testing.B) {
	b.StopTimer()
	elts := createElements(b.N)
	idxes := selectIndexes(b.N)
	b.StartTimer()
	pairwiseUnions(elts, idxes)
	for _, e := range elts {
		_ = e.Find()
	}
}
