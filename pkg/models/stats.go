package models

import (
	"time"

	"gorm.io/gorm"
)

type FizzBuzzRequestStat struct {
	gorm.Model

	// request informations
	Int1    int
	Int2    int
	Int3    int
	String1 string
	String2 string

	Time time.Time
}
