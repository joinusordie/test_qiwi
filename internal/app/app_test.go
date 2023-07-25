package app

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetExchangeRate(t *testing.T) {
	// Arrange
	testTable := []struct {
		data     []string
		expected []string
	}{
		{
			data:     []string{"currency_rates", "--code=USD", "--date=2022-10-08"},
			expected: []string{"USD", "Доллар США", "61,2475"},
		},
		{
			data:     []string{"currency_rates", "--code=USD", "--date=2027-10-08"},
			expected: nil,
		},
		{
			data:     []string{"currency_rates", "--code=D", "--date=2022-10-08"},
			expected: nil,
		},
	}
	// Act
	for _, testCase := range testTable {
		result, _ := GetExchangeRate(testCase.data)

		//Assert
		if !cmp.Equal(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.expected, result)
		}
	}
}
