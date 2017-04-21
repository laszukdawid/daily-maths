package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"
    "time"
)

////////////////////////////////////////

var posOps []operation

type Config struct {
    User    string
    Level   int
    Num     int
}

type operation struct {
    name string
    Func func(args []float32) (float32)
    argsNum int
}

type exercise struct {
    op operation
    args []float32
}


////////////////////////////////////////
func saveResult(score int, config Config) (){
    path := "./results/"+config.User
    var f *os.File
    var err error

    // Creates or loads score file
    if _, err = os.Stat(path); os.IsNotExist(err) {
        fmt.Println("Creating score file for user", config.User)
        f, err = os.Create(path)
    } else {
        f, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
    }

    // Handle error in both cases
    if err != nil {
        panic(err)
    }

    defer f.Close()

    dataStr := time.Now().Format("2006-01-02 15:04:05")
    level := config.Level
    text := fmt.Sprintf("%s %d %d\n", dataStr, score, level)

    if _, err := f.WriteString(text); err != nil {
        panic(err)
    }
}

func evaluateExercise(ex exercise) (float32) {
    return ex.op.Func(ex.args)
}

func displayExcercise(ex exercise) {
    opName := ex.op.name
    args := ex.args

    if opName=="-" && args[1]<0 {
        opName = "+"
        args[1] *= -1
    }
    if opName=="+" && args[1]<0 {
        opName = "-"
        args[1] *= -1
    }

    fmt.Printf("%g %s %g = ", args[0], opName, args[1])
}

////////////////////////////////////////
// Random select
func getRandom(randRange [3]float32) (float32) {
    step := randRange[2]
    min := int(randRange[0]/step)
    max := int(randRange[1]/step)
    return step*float32(rand.Intn(max)+min)
}

func getRandomValues(level int, n int) ([]float32){
    var r []float32
    var randRange [3]float32
    switch {
    case (10 >= level && level >= 7):
        randRange = [3]float32{-30, 30, 0.5}
    case (level >= 4):
        randRange = [3]float32{-20, 20, 0.5}
    case (level >= 0):
        randRange = [3]float32{0, 20, 1}
    default:
        panic("Level value should be within range 0--10")
    }

    for len(r) < n {
        r = append(r, getRandom(randRange))
    }

    return r
}

func getFunction() (operation) {
    f_index := rand.Intn(len(posOps))
    return posOps[f_index]
}

func getExercise(level int) (exercise) {
    var ex exercise

    // TODO: Make harder exercise could have more than one eq
    op := getFunction()
    args := getRandomValues(level, op.argsNum)

    // Divide by zero
    for op.name == "/" && args[1] == 0 {
        args = getRandomValues(level, op.argsNum)
    }

    ex.op = op
    ex.args = args

    return ex
}
////////////////////////////////////////
// Maths operators/functions
func add(args []float32) (float32) {
    return args[0] + args[1]
}

func subtract(args []float32) (float32) {
    return args[0] - args[1]
}

func multiply(args []float32) (float32) {
    return args[0] * args[1]
}

func divide(args []float32) (float32) {
    return args[0] / args[1]
}

////////////////////////////////////////
func initialize(config Config) () {
    // Greetings are important
    fmt.Printf("Hello, %s. Let's do this!\n", config.User)
    fmt.Printf("Selected difficulty: %d\n", config.Level)

    // Random generator initiation
    rand.Seed(time.Now().UnixNano())

    // Check selected number of iterations
    num := config.Num
    if (num < 0) {
        panic("Negative number of iterations. Something's fishy.")
    }

    // Define operations depending on level
    level := config.Level
    if (level < 0 || level > 10) {
        panic("Level should be within range 0 -- 10.")
    }

    posOps = nil
    switch {
        case (level > 7): // HARD
            posOps = append(posOps, operation{"/", divide, 2})
            fallthrough
        case (level > 4): // MEDIUM
            posOps = append(posOps, operation{"*", multiply, 2})
            fallthrough
        case (level > 1): // EASY
            posOps = append(posOps,  operation{"-", subtract, 2})
    }
    posOps = append(posOps,  operation{"+", add, 2})
}
////////////////////////////////////////
// Magic starts here
func main() {
    file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)
    config := Config{}

    err := decoder.Decode(&config)
    if err != nil {
        fmt.Println("Error while decoding config", err)
    }

    // Initialize
    initialize(config)

    // Variables
    var answer, expected float32
    score := 0
    goodResponses := []string{"Awesome.", "Obviously.", "Yep.",
        "10 points for "+config.User+"dor.",
        "Whoever thinks otherwise is stupid."}

    for i:=0; i<config.Num; i++ {
        fmt.Printf("%d. ", (i+1))

        // TODO: Making more sensible operations
        // e.g. rounding when using "/"
        ex := getExercise(config.Level)

        expected = evaluateExercise(ex)
        displayExcercise(ex)
        _, err := fmt.Scanf("%g", &answer)
        if err != nil {
            fmt.Println("Cannot read answer:", err)
        }

        if answer==expected {
            responseIdx := rand.Intn(len(goodResponses))
            fmt.Println(goodResponses[responseIdx])
            score += 1
        } else {
            fmt.Println("Nope. Expected:", expected)
        }
    }
    fmt.Printf("Final score: %d/%d\n", score, config.Num)

    // Saving
    saveResult(score, config)

}
