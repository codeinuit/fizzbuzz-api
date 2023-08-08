package fizzbuzz_test

import (
	"testing"

	"github.com/codeinuit/fizzbuzz-api/pkg/fizzbuzz"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzIncorrectMultiplicationValue(t *testing.T) {
	var tests = []struct {
		name string

		int1, int2, limit int
		word1, word2      string
		want              string
	}{
		{
			"first multiplication value is negative",
			-2, 5, 16,
			"fizz", "buzz",
			fizzbuzz.ERROR_WRONG_MULTI_VALUE,
		},
		{
			"second multiplication value is negative",
			3, -5, 16,
			"fizz", "buzz",
			fizzbuzz.ERROR_WRONG_MULTI_VALUE,
		},
		{
			"both multiplication values are negative",
			-3, -5, 16,
			"fizz", "buzz",
			fizzbuzz.ERROR_WRONG_MULTI_VALUE,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := fizzbuzz.FizzBuzz(test.int1, test.int2, test.limit, test.word1, test.word2)

			assert.EqualError(t, err, test.want)
		})
	}
}

func TestFizzBuzzIncorrectLimitValue(t *testing.T) {
	var tests = []struct {
		name string

		int1, int2, limit int
		word1, word2      string
		want              string
	}{
		{
			"limit is equal to 1",
			3, 5, 1,
			"fizz", "buzz",
			fizzbuzz.ERROR_WRONG_LIMIT_VALUE,
		},
		{
			"limit is negative",
			3, 5, -32,
			"fizz", "buzz",
			fizzbuzz.ERROR_WRONG_LIMIT_VALUE,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := fizzbuzz.FizzBuzz(test.int1, test.int2, test.limit, test.word1, test.word2)

			assert.EqualError(t, err, test.want)
		})
	}
}

func TestFizzBuzz(t *testing.T) {
	output, err := fizzbuzz.FizzBuzz(3, 5, 16, "fizz", "buzz")

	assert.Nil(t, err)
	assert.Equal(t, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16", output)
}
