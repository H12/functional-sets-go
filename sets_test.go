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

// TestUnion

// TestIntersect

// TestDiff

// TestFilter

// TestForAll

// TestExists

// TestMap
