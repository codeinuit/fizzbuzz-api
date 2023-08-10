package models

import (
	"gorm.io/gorm"
)

// Stats is the model to store FizzBuzz calls and
// the numbers of usages
type Stats struct {
	gorm.Model

	// request informations
	Int1    int
	Int2    int
	Int3    int
	String1 string
	String2 string

	Use int
}
