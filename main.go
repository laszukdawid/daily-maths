package main

import (
    "encoding/json"
    "fmt"
    "os"
    "math/rand"
    "time"
)

////////////////////////////////////////
type Config struct {
    User    string
    Level   string
    Num     int
}

type operation struct {
    name string
    Func func(args []float32) (float32)
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

////////////////////////////////////////
// Random select
func getRandom(randRange [3]float32) (float32) {
    step := randRange[2]
    min := int(randRange[0]/step)
    max := int(randRange[1]/step)
    return step*float32(rand.Intn(max)+min)
}

func getRandomValues(level string) ([]float32){
    var r []float32
    var rangRange [3]float32
    var n int
    switch level {
    case "Easy":
        n = 2
        rangRange = [3]float32{0, 20, 1}
    case "Normal":
        n = 2
        rangRange = [3]float32{-20, 20, 0.5}
    case "Hard":
        n = 2
        rangRange = [3]float32{-50, 50, 0.5}
    default:
        panic("DEFINE YOURSELF!")
    }

    for len(r) < n {
        r = append(r, getRandom(rangRange))
    }

    return r
}

func getFunction() (operation) {

    flist := []operation {
                 operation{"+", add},
                 operation{"-", substract},
                 operation{"*", multiply},
                 operation{"/", divide}}

    f_index := rand.Intn(len(flist))
    return flist[f_index]
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
    return args[0] / zeroDivide(args[1])
}

func zeroDivide(r2 float32) (float32) {
    for r2 == 0 {
        r2 = float32(rand.Intn(200)-100)/8
    }
    return r2
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

    fmt.Printf("Hello, %s. Let's do this!\n", config.User)
    fmt.Printf("Selected difficulty: %s\n", config.Level)

    // Random generator initition
    rand.Seed(time.Now().UnixNano())

    // Variables
    var r []float32
    var op operation
    var answer, expected float32
    score := 0
    goodResponses := []string{"Awesome.", "Obviously.", "Yep.",
        "10 points for "+config.User+"dor.",
        "Whoever thinks otherwise is stupid."}

    for i:=0; i<config.Num; i++ {
        fmt.Printf("%d. ", i)

        r = getRandomValues(config.Level)

        // TODO: Making more sensible operations
        // e.g. rounding when using "/"
        op = getFunction()
        expected = op.Func(r)

        fmt.Printf("%g %s %g = ", r[0], op.name, r[1])
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
