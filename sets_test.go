package sets

import "testing"

func TestSingletonSet(t *testing.T) {
	tests := []struct {
		Name string

		SeedValue  int
		InputValue int

		ExpectedResult bool
	}{
		{
			Name: "new singleton set contains the value it was initialized with",

			SeedValue:  1,
			InputValue: 1,

			ExpectedResult: true,
		},
		{
			Name: "new singleton set does not contain a value it was not initialized with",

			SeedValue:  1,
			InputValue: 42,

			ExpectedResult: false,
		},
	}

	for _, entry := range tests {
		t.Run(entry.Name, func(t *testing.T) {
			set := SingletonSet(entry.SeedValue)
			result := set(entry.InputValue)

			if result != entry.ExpectedResult {
				t.Errorf("Expected set(%+v) to be %+v, but got %+v",
					entry.InputValue,
					entry.ExpectedResult,
					result,
				)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	t.Run("union contains all values of two provided sets", func(t *testing.T) {
		set := Union(setFromInts(1, 2), setFromInts(4, 5))
		resultMap := map[int]bool{
			1: true,
			2: true,
			3: false,
			4: true,
			5: true,
		}

		for inputValue, expectedResult := range resultMap {
			result := set(inputValue)
			if result != expectedResult {
				t.Errorf("Expected set(%+v) to be %+v, but got %+v",
					inputValue,
					expectedResult,
					result,
				)
			}
		}
	})
}

func TestIntersect(t *testing.T) {
	t.Run("intersect contains only element shared between two sets", func(t *testing.T) {
		set := Intersect(setFromInts(1, 2, 3), setFromInts(3, 4, 5))
		resultMap := map[int]bool{
			1: false,
			2: false,
			3: true,
			4: false,
			5: false,
			6: false, // a value contained in neither set, for good measure
		}

		for inputValue, expectedResult := range resultMap {
			result := set(inputValue)
			if result != expectedResult {
				t.Errorf("Expected set(%+v) to be %+v, but got %+v",
					inputValue,
					expectedResult,
					result,
				)
			}
		}
	})
}

func TestDiff(t *testing.T) {
	t.Run("intersect includes only values not in provided sets", func(t *testing.T) {
		set := Diff(setFromInts(1, 2), setFromInts(4, 5))
		resultMap := map[int]bool{
			1: false,
			2: false,
			3: true,
			4: false,
			5: false,
		}

		for inputValue, expectedResult := range resultMap {
			result := set(inputValue)
			if result != expectedResult {
				t.Errorf("Expected set(%+v) to be %+v, but got %+v",
					inputValue,
					expectedResult,
					result,
				)
			}
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("filtered set includes only values for which the predicate returns true", func(t *testing.T) {
		inputSet := setFromInts(1, 2, 3, 4, 5, 6, 7, 8)

		set := Filter(inputSet, isEven)
		resultMap := map[int]bool{
			1:  false,
			2:  true,
			3:  false,
			4:  true,
			5:  false,
			6:  true,
			7:  false,
			8:  true,
			10: false, // even number not included in original set should not be in filtered set
		}

		for inputValue, expectedResult := range resultMap {
			result := set(inputValue)
			if result != expectedResult {
				t.Errorf("Expected set(%+v) to be %+v, but got %+v",
					inputValue,
					expectedResult,
					result,
				)
			}
		}
	})
}

func TestForAll(t *testing.T) {
	bigSlice := make([]int, BOUND, BOUND)
	for i := range bigSlice {
		bigSlice[i] = i + 1
	}

	tests := []struct {
		Name string

		InputSet      Set
		PredicateFunc func(i int) bool

		ExpectedResult bool
	}{
		{
			Name: "returns true when predicate holds for all values in the set",

			InputSet:      setFromInts(2, 4, 6, 8, 10),
			PredicateFunc: isEven,

			ExpectedResult: true,
		},
		{
			Name: "returns false when predicate does not hold for all values in the set",

			InputSet:      setFromInts(1, 2, 3, 4, 5),
			PredicateFunc: isEven,

			ExpectedResult: false,
		},
		{
			Name: "returns true for an empty set",

			InputSet:      func(i int) bool { return false },
			PredicateFunc: func(i int) bool { return true },

			ExpectedResult: true,
		},
		{
			Name: "works for a set of all bounded integers",

			InputSet:      setFromInts(bigSlice...),
			PredicateFunc: func(i int) bool { return true },

			ExpectedResult: true,
		},
	}

	for _, entry := range tests {
		t.Run(entry.Name, func(t *testing.T) {
			result := ForAll(entry.InputSet, entry.PredicateFunc)
			if result != entry.ExpectedResult {
				t.Errorf("Expected ForAll to return %+v for the provided set and predicate, but got %+v",
					entry.ExpectedResult,
					result,
				)
			}
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		Name string

		InputSet      Set
		PredicateFunc func(i int) bool

		ExpectedResult bool
	}{
		{
			Name: "returns true when the predicate function holds for all entries in the set",

			InputSet:      setFromInts(2, 4, 6),
			PredicateFunc: isEven,

			ExpectedResult: true,
		},
		{
			Name: "returns true when the predicate function holds for some entries in the set",

			InputSet:      setFromInts(1, 2, 3),
			PredicateFunc: isEven,

			ExpectedResult: true,
		},
		{
			Name: "returns false when the predicate function holds for no entries in the set",

			InputSet:      setFromInts(1, 3, 5),
			PredicateFunc: isEven,

			ExpectedResult: false,
		},
	}

	for _, entry := range tests {
		result := Exists(entry.InputSet, entry.PredicateFunc)
		if result != entry.ExpectedResult {
			t.Errorf("Expected Exists to return %+v for the provided set and predicate, but got %+v",
				entry.ExpectedResult,
				result,
			)
		}
	}
}

func TestMap(t *testing.T) {
	t.Run("returned Set returns the appropriate results for the mapped and unmapped values", func(t *testing.T) {
		inputSet := setFromInts(1, 2, 3, 4, 5)
		doubleValue := func(i int) int {
			return i * 2
		}

		set := Map(inputSet, doubleValue)
		reusultMap := map[int]bool{
			1:  false,
			2:  true,
			3:  false,
			4:  true,
			5:  false,
			6:  true,
			7:  false,
			8:  true,
			9:  false,
			10: true,
		}

		for inputValue, expectedResult := range reusultMap {
			result := set(inputValue)
			if result != expectedResult {
				t.Errorf("expected mapped set to return %+v for value %+v, but got %+v",
					expectedResult,
					inputValue,
					result,
				)
			}
		}
	})
}

func isEven(i int) bool {
	return i%2 == 0
}

func setFromInts(ints ...int) Set {
	return func(v int) bool {
		for _, i := range ints {
			if i == v {
				return true
			}
		}

		return false
	}
}
