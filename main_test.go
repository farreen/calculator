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

// func TestCalculateUsing4OperandsAndSameOp(t *testing.T) {
// 	assert.Equal(t, 15, CalculateUsing4OperandsAndSameOp(6, "+", 4, "+", 3, "+", 2))
// 	assert.Equal(t, -3, CalculateUsing4OperandsAndSameOp(6, "-", 4, "-", 3, "-", 2))
// 	assert.Equal(t, 144, CalculateUsing4OperandsAndSameOp(6, "*", 4, "*", 3, "*", 2))
// 	assert.Equal(t, 0, CalculateUsing4OperandsAndSameOp(6, "/", 4, "/", 3, "/", 2))
// }
//
// func TestCalculateUsing4Operands(t *testing.T) {
// 	assert.Equal(t, 12, CalculateUsing4Operands(2, "+", 4, "+", 4, "+", 2))
// 	assert.Equal(t, 12, CalculateUsing4Operands(6, "+", 4, "+", 4, "-", 2))
// 	assert.Equal(t, 18, CalculateUsing4Operands(6, "+", 4, "+", 4, "*", 2))
// 	assert.Equal(t, 12, CalculateUsing4Operands(6, "+", 4, "+", 4, "/", 2))
//
// 	assert.Equal(t, 11, CalculateUsing4Operands(6, "+", 4, "-", 3, "+", 4))
// 	assert.Equal(t, 3, CalculateUsing4Operands(6, "+", 4, "-", 3, "-", 4))
// 	assert.Equal(t, -2, CalculateUsing4Operands(6, "+", 4, "-", 3, "*", 4))
// 	assert.Equal(t, 10, CalculateUsing4Operands(6, "+", 4, "-", 3, "/", 4))
//
// 	assert.Equal(t, 32, CalculateUsing4Operands(6, "+", 4, "*", 5, "+", 6))
// 	assert.Equal(t, 20, CalculateUsing4Operands(6, "+", 4, "*", 5, "-", 6))
// 	assert.Equal(t, 126, CalculateUsing4Operands(6, "+", 4, "*", 5, "*", 6))
// 	assert.Equal(t, 9, CalculateUsing4Operands(6, "+", 4, "*", 5, "/", 6))
//
// 	assert.Equal(t, 9, CalculateUsing4Operands(6, "+", 4, "/", 2, "+", 1))
// 	assert.Equal(t, 7, CalculateUsing4Operands(6, "+", 4, "/", 2, "-", 1))
// 	assert.Equal(t, 8, CalculateUsing4Operands(6, "+", 4, "/", 2, "*", 1))
// 	assert.Equal(t, 8, CalculateUsing4Operands(6, "+", 4, "/", 2, "/", 1))
//
// 	assert.Equal(t, 5, CalculateUsing4Operands(6, "-", 4, "+", 2, "+", 1))
// 	assert.Equal(t, 3, CalculateUsing4Operands(6, "-", 4, "+", 2, "-", 1))
// 	assert.Equal(t, 4, CalculateUsing4Operands(6, "-", 4, "+", 2, "*", 1))
// 	assert.Equal(t, 4, CalculateUsing4Operands(6, "-", 4, "+", 2, "/", 1))
//
// 	assert.Equal(t, 1, CalculateUsing4Operands(6, "-", 4, "-", 2, "+", 1))
// 	assert.Equal(t, -1, CalculateUsing4Operands(6, "-", 4, "-", 2, "-", 1))
// 	assert.Equal(t, 0, CalculateUsing4Operands(6, "-", 4, "-", 2, "*", 1))
// 	assert.Equal(t, 0, CalculateUsing4Operands(6, "-", 4, "-", 2, "/", 1))
//
// 	assert.Equal(t, -1, CalculateUsing4Operands(6, "-", 4, "*", 2, "+", 1))
// 	assert.Equal(t, -3, CalculateUsing4Operands(6, "-", 4, "*", 2, "-", 1))
// 	assert.Equal(t, -2, CalculateUsing4Operands(6, "-", 4, "*", 2, "*", 1))
// 	assert.Equal(t, -2, CalculateUsing4Operands(6, "-", 4, "*", 2, "/", 1))
//
// 	assert.Equal(t, 5, CalculateUsing4Operands(6, "-", 4, "/", 2, "+", 1))
// 	assert.Equal(t, 3, CalculateUsing4Operands(6, "-", 4, "/", 2, "-", 1))
// 	assert.Equal(t, 4, CalculateUsing4Operands(6, "-", 4, "/", 2, "*", 1))
// 	assert.Equal(t, 4, CalculateUsing4Operands(6, "-", 4, "/", 2, "/", 1))
//
// 	assert.Equal(t, 27, CalculateUsing4Operands(6, "*", 4, "+", 2, "+", 1))
// 	assert.Equal(t, 25, CalculateUsing4Operands(6, "*", 4, "+", 2, "-", 1))
// 	assert.Equal(t, 26, CalculateUsing4Operands(6, "*", 4, "+", 2, "*", 1))
// 	assert.Equal(t, 26, CalculateUsing4Operands(6, "*", 4, "+", 2, "/", 1))
//
// 	assert.Equal(t, 23, CalculateUsing4Operands(6, "*", 4, "-", 2, "+", 1))
// 	assert.Equal(t, 20, CalculateUsing4Operands(6, "*", 4, "-", 2, "-", 1))
// 	assert.Equal(t, 22, CalculateUsing4Operands(6, "*", 4, "-", 2, "*", 1))
// 	assert.Equal(t, 22, CalculateUsing4Operands(6, "*", 4, "-", 2, "/", 1))
//
// 	assert.Equal(t, 49, CalculateUsing4Operands(6, "*", 4, "*", 2, "+", 1))
// 	assert.Equal(t, 47, CalculateUsing4Operands(6, "*", 4, "*", 2, "-", 1))
// 	assert.Equal(t, 48, CalculateUsing4Operands(6, "*", 4, "*", 2, "*", 1))
// 	assert.Equal(t, 48, CalculateUsing4Operands(6, "*", 4, "*", 2, "/", 1))
//
// 	assert.Equal(t, 13, CalculateUsing4Operands(6, "*", 4, "/", 2, "+", 1))
// 	assert.Equal(t, 11, CalculateUsing4Operands(6, "*", 4, "/", 2, "-", 1))
// 	assert.Equal(t, 12, CalculateUsing4Operands(6, "*", 4, "/", 2, "*", 1))
// 	assert.Equal(t, 12, CalculateUsing4Operands(6, "*", 4, "/", 2, "/", 1))
//
// 	assert.Equal(t, 4, CalculateUsing4Operands(6, "/", 4, "+", 2, "+", 1))
// 	assert.Equal(t, 2, CalculateUsing4Operands(6, "/", 4, "+", 2, "-", 1))
// 	assert.Equal(t, 3, CalculateUsing4Operands(6, "/", 4, "+", 2, "*", 1))
// 	assert.Equal(t, 3, CalculateUsing4Operands(6, "/", 4, "+", 2, "/", 1))
//
// 	assert.Equal(t, 0, CalculateUsing4Operands(6, "/", 4, "-", 2, "+", 1))
// 	assert.Equal(t, -2, CalculateUsing4Operands(6, "/", 4, "-", 2, "-", 1))
// 	assert.Equal(t, -1, CalculateUsing4Operands(6, "/", 4, "-", 2, "*", 1))
// 	assert.Equal(t, 0, CalculateUsing4Operands(6, "/", 4, "-", 2, "/", 1))
//
// 	assert.Equal(t, 3, CalculateUsing4Operands(6, "/", 4, "*", 2, "+", 1))
// 	assert.Equal(t, 1, CalculateUsing4Operands(6, "/", 4, "*", 2, "-", 1))
// 	assert.Equal(t, 2, CalculateUsing4Operands(6, "/", 4, "*", 2, "*", 1))
// 	assert.Equal(t, 2, CalculateUsing4Operands(6, "/", 4, "*", 2, "/", 1))
//
// 	assert.Equal(t, 1, CalculateUsing4Operands(6, "/", 4, "/", 2, "+", 1))
// 	assert.Equal(t, -1, CalculateUsing4Operands(6, "/", 4, "/", 2, "-", 1))
// 	assert.Equal(t, 0, CalculateUsing4Operands(6, "/", 4, "/", 2, "*", 1))
// 	assert.Equal(t, 0, CalculateUsing4Operands(6, "/", 4, "/", 2, "/", 1))
// }
