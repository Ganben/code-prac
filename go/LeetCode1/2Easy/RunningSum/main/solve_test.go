package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	assert := assert.New(t)
	i1 := []int{1, 2, 3, 4}
	r1 := []int{1, 3, 6, 10}

	r0 := runningSum(i1)
	t.Log(r0)
	assert.Equal(r0, r1)

}
