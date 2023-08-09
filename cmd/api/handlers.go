package main

import (
	"net/http"

	"github.com/codeinuit/fizzbuzz-api/pkg/fizzbuzz"
	logger "github.com/codeinuit/fizzbuzz-api/pkg/log"
	"github.com/codeinuit/fizzbuzz-api/pkg/models"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	log logger.Logger
	db  *Database
}

// getFizzBuzzBody represent the input structure for
// the FizzBuzz endpoint
type getFizzBuzzBody struct {
	Int1    int    `json:"int1" binding:"required"`
	Int2    int    `json:"int2" binding:"required"`
	Int3    int    `json:"int3" binding:"required"`
	String1 string `json:"str1" binding:"required"`
	String2 string `json:"str2" binding:"required"`
}

// healthcheck works as a ping and returns a OK status
// GET /health
func (h handlers) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// fizzbuzz handles fizzbuzz
// GET /fizzbuzz
func (h handlers) fizzbuzz(c *gin.Context) {
	var v getFizzBuzzBody

	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter"})
		return
	}

	fb, err := fizzbuzz.FizzBuzz(v.Int1, v.Int2, v.Int3, v.String1, v.String2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.UsageUpdate(models.Stats{
		Int1:    uint8(v.Int1),
		Int2:    uint8(v.Int2),
		Int3:    uint8(v.Int3),
		String1: v.String1,
		String2: v.String2,
	})

	c.String(http.StatusOK, fb)
}

// stats return the fizzbuzz stats
// GET /stats
func (h handlers) stats(c *gin.Context) {
	res, err := h.db.CountUsage()
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"int1": res.Int1, "int2": res.Int2, "int3": res.Int3, "string1": res.String1, "string2": res.String2, "used": res.Use})
}
