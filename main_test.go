package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateUsingMap(t *testing.T) {
	assert.Equal(t, 10, CalculateUsingMap(6, "+", 4))
	assert.Equal(t, 2, CalculateUsingMap(6, "-", 4))
	assert.Equal(t, 24, CalculateUsingMap(6, "*", 4))
	assert.Equal(t, 1, CalculateUsingMap(6, "/", 4))

}

func TestCalculateUsingSwitch(t *testing.T) {
	assert.Equal(t, 10, CalculateUsingSwitch(6, "+", 4))
	assert.Equal(t, 2, CalculateUsingSwitch(6, "-", 4))
	assert.Equal(t, 24, CalculateUsingSwitch(6, "*", 4))
	assert.Equal(t, 1, CalculateUsingSwitch(6, "/", 4))
	assert.Equal(t, 2, CalculateUsingSwitch(6, "%", 4))
}

func TestCalculateUsingIfelse(t *testing.T) {
	assert.Equal(t, 10, CalculateUsingIfelse(6, "+", 4))
	assert.Equal(t, 2, CalculateUsingIfelse(6, "-", 4))
	assert.Equal(t, 24, CalculateUsingIfelse(6, "*", 4))
	assert.Equal(t, 1, CalculateUsingIfelse(6, "/", 4))
	assert.Equal(t, 2, CalculateUsingIfelse(6, "%", 4))
}

func TestCalculateUsing3OperandsAndSameOp(t *testing.T) {
	assert.Equal(t, 13, CalculateUsing3OperandsAndSameOp(6, "+", 4, "+", 3))
	assert.Equal(t, -1, CalculateUsing3OperandsAndSameOp(6, "-", 4, "-", 3))
	assert.Equal(t, 72, CalculateUsing3OperandsAndSameOp(6, "*", 4, "*", 3))
	assert.Equal(t, 0, CalculateUsing3OperandsAndSameOp(6, "/", 4, "/", 3))
	assert.Equal(t, 2, CalculateUsing3OperandsAndSameOp(6, "%", 4, "%", 3))
}

func TestCalculateUsing3Operands(t *testing.T) {
	assert.Equal(t, 10, CalculateUsing3Operands(2, "+", 4, "+", 4))
	assert.Equal(t, 6, CalculateUsing3Operands(6, "+", 4, "-", 4))
	assert.Equal(t, 22, CalculateUsing3Operands(6, "+", 4, "*", 4))
	assert.Equal(t, 7, CalculateUsing3Operands(6, "+", 4, "/", 4))
	assert.Equal(t, 5, CalculateUsing3Operands(6, "-", 4, "+", 3))
	assert.Equal(t, -1, CalculateUsing3Operands(6, "-", 4, "-", 3))
	assert.Equal(t, -6, CalculateUsing3Operands(6, "-", 4, "*", 3))
	assert.Equal(t, 5, CalculateUsing3Operands(6, "-", 4, "/", 3))
	assert.Equal(t, 29, CalculateUsing3Operands(6, "*", 4, "+", 5))
	assert.Equal(t, 19, CalculateUsing3Operands(6, "*", 4, "-", 5))
	assert.Equal(t, 120, CalculateUsing3Operands(6, "*", 4, "*", 5))
	assert.Equal(t, 4, CalculateUsing3Operands(6, "*", 4, "/", 5))
	assert.Equal(t, 3, CalculateUsing3Operands(6, "/", 4, "+", 2))
	assert.Equal(t, -1, CalculateUsing3Operands(6, "/", 4, "-", 2))
	assert.Equal(t, 2, CalculateUsing3Operands(6, "/", 4, "*", 2))
	assert.Equal(t, 0, CalculateUsing3Operands(6, "/", 4, "/", 2))

}
