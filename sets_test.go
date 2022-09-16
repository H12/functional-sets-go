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
		isEven := func(i int) bool { return i%2 == 0 }
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

// TestForAll

// TestExists

// TestMap

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
