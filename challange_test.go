package main

import (
    "fmt"
    "testing"
)
////////////////////
func assertPanic(t *testing.T, f func(), msg string) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic.\n%s",msg)
        }
    }()
    f()
}
func assertNotPanic(t *testing.T, f func(), msg string) {
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("The code paniced, but it shouldn't!\n%s",msg)
        }
    }()
    f()
}
func assertEqual(t *testing.T, a interface{}, b interface{}, msg string) {
  if a == b {
      return
  }
  msg = msg + fmt.Sprintf("\n%v != %v", a, b)
  t.Fatal(msg)
}
////////////////////

func TestLevel(t *testing.T) {
    // Testing Level values, which should be in range
    // 0 -- 10 (inclusive)
    var config Config
    var msg string

    // Fail - level = -1
    config = Config{"User", -1, 5}
    msg = "Testing `Initialize` with level=-1"
    assertPanic(t, func(){initialize(config)}, msg)

    // Fail - level = 11
    config = Config{"User", 11, 5}
    msg = "Testing `Initialize` with level=11"
    assertPanic(t, func(){initialize(config)}, msg)

    // Success - level = 0
    config = Config{"User", 0, 5}
    msg = "Testing `Initialize` with level=0"
    assertNotPanic(t, func(){initialize(config)}, msg)

    // Success - level = 10
    config = Config{"User", 10, 5}
    msg = "Testing `Initialize` with level=10"
    assertNotPanic(t, func(){initialize(config)}, msg)
}

func TestIteration(t *testing.T) {
    // Testing iterations in configuration.
    // No even convinced that it should be passed.

    var config Config
    var msg string

    // Fail - iter < 0
    config = Config{"User", 5, -1}
    msg = "Testing `initialize` with iter=-1"
    assertPanic(t, func(){initialize(config)}, msg)

    // Fail - iter = 0
    config = Config{"User", 0, -1}
    msg = "Testing `initialize` with iter=0"
    assertPanic(t, func(){initialize(config)}, msg)

    // Success - iter = 1
    config = Config{"User", 1, 5}
    msg = "Testing `initialize` with iter=1"
    assertNotPanic(t, func(){initialize(config)}, msg)

}

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

func TestOperations(t *testing.T) {
    // Mathematical operations are randomly selected from a set
    // where its content depends on the level.

    user := "User"
    iter := 5

    initialize(Config{user, 0, iter})
    assertEqual(t, len(posOps), 1, "Level 0 - only +")

    initialize(Config{user, 2, iter})
    assertEqual(t, len(posOps), 2, "Level 2 - +-")

    initialize(Config{user, 5, iter})
    assertEqual(t, len(posOps), 3, "Level 5 - +-*")

    initialize(Config{user, 8, iter})
    assertEqual(t, len(posOps), 4, "Level 8 - +-*/")
}
