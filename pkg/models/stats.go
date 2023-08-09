package models

import (
	"gorm.io/gorm"
)

type Stats struct {
	gorm.Model

	// request informations
	Int1    uint8
	Int2    uint8
	Int3    uint8
	String1 string
	String2 string

	Use uint8
}
