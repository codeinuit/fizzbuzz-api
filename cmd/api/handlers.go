package main

import (
	"net/http"

	"github.com/codeinuit/fizzbuzz-api/pkg/fizzbuzz"
	logger "github.com/codeinuit/fizzbuzz-api/pkg/log"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	log logger.Logger
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, fb)
}
