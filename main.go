package main

import (
    "encoding/json"
    "fmt"
    "os"
    "math/rand"
    //"strconv"
    "time"
)

type Config struct {
    User    string
    Level   string
    Num     int
}

type operation struct {
    name string
    fname func(args []float32) (float32)
}

func save_result(score int, config Config) (){
    path := "./results/"+config.User

    // If no score for given user
    if _, err := os.Stat(path); os.IsNotExist(err) {
        fmt.Println("Creating score file for user", config.User)
        f, err := os.Create(path)
        if err != nil {
            panic(err)
        }
        f.Close()
    }
    // File exists, just open it
    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }

    defer f.Close()

    data_str := time.Now().Format("2006-01-02 15:04:05")
    level := config.Level[0:1]
    text := fmt.Sprintf("%s %d %s\n", data_str, score, level)
    if _, err = f.WriteString(text); err != nil {
        panic(err)
    }
}

func get_random_values(level string) (float32, float32){
    seed := rand.NewSource(time.Now().UnixNano())
    randgen := rand.New(seed)
    var r1, r2 float32
    switch level {
    case "Easy":
        r1 = float32(randgen.Intn(20))/2
        r2 = float32(randgen.Intn(20))/2
    case "Normal":
        r1 = float32(randgen.Intn(100)-50)/2
        r2 = float32(randgen.Intn(100)-50)/2
    case "Hard":
        r1 = float32(randgen.Intn(200)-100)/8
        r2 = float32(randgen.Intn(200)-100)/8
    default:
        panic("DEFINE YOURSELF!")
    }
    return r1, r2
}

func get_function() (operation) {
    seed := rand.NewSource(time.Now().UnixNano())
    randgen := rand.New(seed)

    flist := []operation {
                 operation{"+", add},
		             operation{"-", substract},
		             operation{"*", multiply},
		             operation{"/", divide}}

    f_index := randgen.Intn(len(flist))
    return flist[f_index]
}

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
    seed := rand.NewSource(time.Now().UnixNano())
    randgen := rand.New(seed)
    for r2 == 0 {
        r2 = float32(randgen.Intn(200)-100)/8
    }
    return r2
}

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

    var r1, r2 float32
    var op operation

    // Variables
    var answer, expected float32
    score := 0
    good_responses := []string{"Awesome.", "Obviously.", "Yep.",
        "10 points for "+config.User+"dor.",
        "Whoever thinks otherwise is stupid."}

    for i:=0; i<config.Num; i++ {
        fmt.Printf("%d. ", i)

        // TODO: Why only two variables?
        r1, r2 = get_random_values(config.Level)

        // TODO: Making more sensible operations
        // e.g. rounding when using "/"
        op = get_function()
        expected = op.fname([]float32{r1, r2})

        fmt.Printf("%g %s %g = ", r1, op.name, r2)
        _, err := fmt.Scanf("%g", &answer)
        if err != nil {
            fmt.Println("Cannot read answer:", err)
        }

        if answer==expected {
            response_idx := rand.Intn(len(good_responses))
            fmt.Println(good_responses[response_idx])
            score += 1
        } else {
            fmt.Println("Nope. Excpected:", expected)
        }
    }
    fmt.Printf("Final score: %d/%d\n", score, config.Num)

    // Saving
    save_result(score, config)

}
