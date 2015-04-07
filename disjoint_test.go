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
	sets := make([]*Set, N)
	for i := 0; i < N; i++ {
		sets[i] = Singleton(i)
	}

	// Merge each even number with its predecessor and each odd number with
	// its predecessor.
	for i := 2; i < N; i++ {
		sets[i].Merge(sets[i-2])
	}
	for i := 3; i < N; i++ {
		sets[i].Merge(sets[i-2])
	}

	// Ensure that even numbers are in the same union as other even numbers
	// and odd numbers are in the same union as other oddn numbers.
	for i := 0; i < N*2; i++ {
		s1 := rand.Intn(N)
		s2 := rand.Intn(N)
		sameMod2 := s1%2 == s2%2
		sameRep := sets[s1].Representative() == sets[s2].Representative()
		if sameMod2 != sameRep {
			t.Fatalf("Are %d and %d in the same set?  The package incorrectly says %v",
				sets[s1], sets[s2], sameRep)
		}
	}
}
