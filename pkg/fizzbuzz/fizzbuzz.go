package fizzbuzz

import (
	"errors"
	"strconv"
)

const (
	ERROR_WRONG_LIMIT_VALUE string = "limit must be a positive integer at least supperior or equal to 1"
	ERROR_WRONG_MULTI_VALUE string = "multiplications values must be a positive integer at least supperior to 1"
)

func FizzBuzz(mult1, mult2, limit int, word1, word2 string) (output string, err error) {
	var count int = 1

	if limit <= 1 {
		return output, errors.New(ERROR_WRONG_LIMIT_VALUE)
	}

	if mult1 < 1 || mult2 < 1 {
		return output, errors.New(ERROR_WRONG_MULTI_VALUE)
	}

	for count <= limit {
		buzz := count%mult1 == 0
		fizz := count%mult2 == 0

		if !buzz && !fizz {
			output += strconv.Itoa(count)
		} else {
			if buzz {
				output += word1
			}

			if fizz {
				output += word2
			}
		}

		if count != limit {
			output += ","
		}
		count++
	}

	return output, err
}
