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
	tests := []struct {
		Name string

		Set1 Set
		Set2 Set

		ResultMap map[int]bool
	}{
		{
			Name: "union contains all values of two provided sets",

			Set1: setFromInts(1, 2),
			Set2: setFromInts(4, 5),

			ResultMap: map[int]bool{
				1: true,
				2: true,
				3: false,
				4: true,
				5: true,
			},
		},
	}

	for _, entry := range tests {
		t.Run(entry.Name, func(t *testing.T) {
			set := Union(entry.Set1, entry.Set2)

			for inputValue, expectedResult := range entry.ResultMap {
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
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		Name string

		Set1 Set
		Set2 Set

		ResultMap map[int]bool
	}{
		{
			Name: "intersect contains only element shared between two sets",

			Set1: setFromInts(1, 2, 3),
			Set2: setFromInts(3, 4, 5),

			ResultMap: map[int]bool{
				1: false,
				2: false,
				3: true,
				4: false,
				5: false,
				6: false, // a value contained in neither set, for good measure
			},
		},
	}

	for _, entry := range tests {
		t.Run(entry.Name, func(t *testing.T) {
			set := Intersect(entry.Set1, entry.Set2)

			for inputValue, expectedResult := range entry.ResultMap {
				result := set(inputValue)
				if result != expectedResult {
					t.Errorf("Expected set (%+v) to be %+v, but got %+v",
						inputValue,
						expectedResult,
						result,
					)
				}
			}
		})
	}
}

// TestDiff

// TestFilter

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
