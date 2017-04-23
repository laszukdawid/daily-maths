package challange

import (
    "fmt"
    "math"
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
    assertPanic(t, func(){Initialize(config)}, msg)

    // Fail - level = 11
    config = Config{"User", 11, 5}
    msg = "Testing `Initialize` with level=11"
    assertPanic(t, func(){Initialize(config)}, msg)

    // Success - level = 0
    config = Config{"User", 0, 5}
    msg = "Testing `Initialize` with level=0"
    assertNotPanic(t, func(){Initialize(config)}, msg)

    // Success - level = 10
    config = Config{"User", 10, 5}
    msg = "Testing `Initialize` with level=10"
    assertNotPanic(t, func(){Initialize(config)}, msg)
}

func TestIteration(t *testing.T) {
    // Testing iterations in configuration.
    // No even convinced that it should be passed.

    var config Config
    var msg string

    // Fail - iter < 0
    config = Config{"User", 5, -1}
    msg = "Testing `Initialize` with iter=-1"
    assertPanic(t, func(){Initialize(config)}, msg)

    // Fail - iter = 0
    config = Config{"User", 0, -1}
    msg = "Testing `Initialize` with iter=0"
    assertPanic(t, func(){Initialize(config)}, msg)

    // Success - iter = 1
    config = Config{"User", 1, 5}
    msg = "Testing `Initialize` with iter=1"
    assertNotPanic(t, func(){Initialize(config)}, msg)

}

/////////// TEST RANDOM ///////////
func TestRandomRange(t *testing.T) {
    testRepeats := 40
    testRange := [3]float32{0, 20, 0.25}

    for i:=0; i<testRepeats; i++ {
        val := GetRandom(testRange)
        if (val<testRange[0] || val>testRange[1]) {
            panic("Out of range")
        }
        // Caution: Mod doesn't round. Step needs to be 2**z
        if math.Mod(float64(val), float64(testRange[2])) != 0 {
            panic("Wrong step")
        }
    }
}

/////////// TEST OPERATIONS ///////////
func TestAdd(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x+y) != Add(args) {
        panic("Simple function is not working.")
    }
}

func TestSubtract(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x-y) != Subtract(args) {
        panic("Simple function is not working.")
    }
}

func TestMutiply(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x*y) != Multiply(args) {
        panic("Simple function is not working.")
    }
}

func TestDivide(t *testing.T) {
    x, y := float32(3), float32(5)
    args := []float32{x, y}
    if float32(x/y) != Divide(args) {
        panic("Simple function is not working.")
    }
}

func TestOperations(t *testing.T) {
    // Mathematical operations are randomly selected from a set
    // where its content depends on the level.

    user := "User"
    iter := 5

    Initialize(Config{user, 0, iter})
    assertEqual(t, len(posOps), 1, "Level 0 - only +")

    Initialize(Config{user, 2, iter})
    assertEqual(t, len(posOps), 2, "Level 2 - +-")

    Initialize(Config{user, 5, iter})
    assertEqual(t, len(posOps), 3, "Level 5 - +-*")

    Initialize(Config{user, 8, iter})
    assertEqual(t, len(posOps), 4, "Level 8 - +-*/")
}
