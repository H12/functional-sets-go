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

		InputValues    []int
		ExpectedResult bool
	}{
		{
			Name: "union contains both values of two singleton sets",

			Set1: SingletonSet(1),
			Set2: SingletonSet(2),

			InputValues:    []int{1, 2},
			ExpectedResult: true,
		},
		{
			Name: "union does not contain value that is in neither of two singleton sets",

			Set1: SingletonSet(1),
			Set2: SingletonSet(2),

			InputValues:    []int{3},
			ExpectedResult: false,
		},
	}

	for _, entry := range tests {
		t.Run(entry.Name, func(t *testing.T) {
			set := Union(entry.Set1, entry.Set2)

			for _, v := range entry.InputValues {
				result := set(v)
				if result != entry.ExpectedResult {
					t.Errorf("Expected set(%+v) to be %+v, but got %+v",
						v,
						entry.ExpectedResult,
						result,
					)
				}
			}
		})
	}
}

// TestIntersect

// TestDiff

// TestFilter

// TestForAll

// TestExists

// TestMap
