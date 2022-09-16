// Package sets provides functions for working with Sets.
//
// In this package, a Set is defined as a function that accepts an integer N,
// and returns a boolean indicating whether N is in the Set.
package sets

type Set func(i int) bool

// SingletonSet returns the set of the one given element.
func SingletonSet(i int) Set {
	return func(j int) bool {
		return i == j
	}
}

// Union returns the union of the two given sets, the sets of all elements that
// are in either `set1` or `set2`.
func Union(set1, set2 Set) Set {
	return func(i int) bool {
		return set1(i) || set2(i)
	}
}

// Returns the intersection of the two given sets, the set of all elements that
// are both in `set1` and `set2`.
func Intersect(set1, set2 Set) Set {
	return func(i int) bool {
		return set1(i) && set2(i)
	}
}

// Returns the difference of the two given sets, the set of all elements of
// `set1` that are not in `set2`.
func Diff(set1, set2 Set) Set {
	union := Union(set1, set2)

	return func(i int) bool {
		return !union(i)
	}
}

// Returns the subset of `set` for which `predicateFunc` holds.
func Filter(set Set, predicateFunc func(i int) bool) Set {
	panic("Implement Filter")
}

// The bounds for `ForAll` and `Exists` are +/- 1000.
var BOUND = 1000

// Returns whether all bounded integers within `set` satisfy `predicateFunc`.
func ForAll(set Set, predicateFunc func(i int) bool) bool {
	panic("Implement ForAll")
}

// Returns whether there exists a bounded integer within `set` that satisfies
// `predicateFunc`.
func Exists(set Set, predicateFunc func(i int) bool) bool {
	panic("Implement Exists")
}

// Returns a set transformed by applying `fnc` to each element of `set`.
func Map(set Set, fnc func(i int) int) Set {
	panic("Implement Map")
}
