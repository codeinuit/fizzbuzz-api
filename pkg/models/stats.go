package models

import (
	"gorm.io/gorm"
)

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
