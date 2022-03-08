package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeromelesaux/fizzbuzz/engine"
	"github.com/jeromelesaux/fizzbuzz/model"
	"github.com/jeromelesaux/fizzbuzz/persistence"
)

type Fizzbuzz struct {
}

func NewFizzbuzz() *Fizzbuzz {
	return &Fizzbuzz{}
}

// Fizzbuzz godoc
// @Summary      get the fizzbuzz slice of string
// @Description  get the fizzbuzz slice of string
// @Tags         fizzbuzz
// @Accept       json
// @Produce      json
// @Param        int1    query     string  false  "first integer"  Format(int)
// @Param        int2    query     string  false  "second integer"  Format(int)
// @Param        str1    query     string  false  "first string to replace"  Format(string)
// @Param        str2    query     string  false  "second string to replace"  Format(string)
// @Param        limit    query     string  false  "limit to reach"  Format(int)
// @Success      200  {array}   model.Parameters
// @Router       /fizzbuzz [get]
func (f *Fizzbuzz) Fizzbuzz(c *gin.Context) {
	p, err := model.ParseParameters(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	select {
	case persistence.AddParametersChan <- *p:
		log.Printf("Sending to channel the parameters %v\n", p)
	default:
		log.Printf("Warn: channel is full")
	}
	result := engine.DoFizzbuzz(p)
	c.JSON(http.StatusAccepted, result)
}

// GetStats godoc
// @Summary      return the maximum hits request statistics
// @Description  return the maximum hits request statistics
// @Tags         stats
// @Produce      json
// @Success      200  {array}   model.Parameters
// @Router       /stats [get]
func (f *Fizzbuzz) GetStats(c *gin.Context) {
	p, occurence := persistence.GetMostFrequent()
	p.Hits = int64(occurence)
	c.JSON(http.StatusOK, p)
}
