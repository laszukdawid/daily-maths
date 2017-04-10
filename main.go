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
    Level   string
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
    level := config.Level[0:1]
    text := fmt.Sprintf("%s %d %s\n", dataStr, score, level)

    if _, err := f.WriteString(text); err != nil {
        panic(err)
    }
}

func evaluateExercise(ex exercise) (float32) {
    return ex.op.Func(ex.args)
}

func displayExcercise(ex exercise) {
    fmt.Printf("%g %s %g = ", ex.args[0], ex.op.name, ex.args[1])
}

////////////////////////////////////////
// Random select
func getRandom(randRange [3]float32) (float32) {
    step := randRange[2]
    min := int(randRange[0]/step)
    max := int(randRange[1]/step)
    return step*float32(rand.Intn(max)+min)
}

func getRandomValues(level string, n int) ([]float32){
    var r []float32
    var randRange [3]float32
    switch level {
    case "Easy":
        randRange = [3]float32{0, 20, 1}
    case "Normal":
        randRange = [3]float32{-20, 20, 0.5}
    case "Hard":
        randRange = [3]float32{-50, 50, 0.5}
    default:
        panic("DEFINE YOURSELF!")
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

func getExercise(level string) (exercise) {
    var ex exercise

    // TODO: Make harder exercise could have more than one eq
    op := getFunction()
    args := getRandomValues(level, op.argsNum)

    // Divide by zero
    for op.name == "/" && args[1] == 0 {
        fmt.Println("Repeat")
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

func substract(args []float32) (float32) {
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
    fmt.Printf("Selected difficulty: %s\n", config.Level)

    // Random generator initition
    rand.Seed(time.Now().UnixNano())

    // Define operations depending on level
    posOps = append(posOps,  operation{"+", add, 2})
    posOps = append(posOps,  operation{"-", substract, 2})

    if config.Level != "Easy" {
        posOps = append(posOps, operation{"*", multiply, 2})
        posOps = append(posOps, operation{"/", divide, 2})
    }
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
        fmt.Printf("%d. ", i)

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
            fmt.Println("Nope. Excpected:", expected)
        }
    }
    fmt.Printf("Final score: %d/%d\n", score, config.Num)

    // Saving
    saveResult(score, config)

}
