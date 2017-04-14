package main

import (
    "testing"
    )

func TestAdd(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x+y) != add(args) {
        panic("Something is fucked up.")
    }
}

func TestSubtract(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x-y) != subtract(args) {
        panic("Something is fucked up.")
    }
}

func TestMutiply(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x*y) != multiply(args) {
        panic("Something is fucked up.")
    }
}

func TestDivide(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x/y) != divide(args) {
        panic("Something is fucked up.")
    }
}
