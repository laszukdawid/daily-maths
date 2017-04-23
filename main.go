//+build !test
package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"

    "github.com/laszukdawid/daily-maths/challange"
)

func main() {
    file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)
    config := challange.Config{}

    err := decoder.Decode(&config)
    if err != nil {
        fmt.Println("Error while decoding config", err)
    }

    // Initialize
    challange.Initialize(config)

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
        ex := challange.GetExercise(config.Level)

        expected = challange.EvaluateExercise(ex)
        challange.DisplayExcercise(ex)
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
    challange.SaveResult(score, config)

}
