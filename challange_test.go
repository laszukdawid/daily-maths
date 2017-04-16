package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x+y) != add(args) {
        panic("Simple function is not working.")
    }
}

func TestSubtract(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x-y) != subtract(args) {
        panic("Simple function is not working.")
    }
}

func TestMutiply(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x*y) != multiply(args) {
        panic("Simple function is not working.")
    }
}

func TestDivide(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x/y) != divide(args) {
        panic("Simple function is not working.")
    }
}
