package main

import (
 "fmt"
 "time"
 "math/rand"
 "os"
)

var score int

func finished(timer *time.Timer) {
  <-timer.C
  fmt.Printf("\nGame over - Score: %d\n", score)
  os.Exit(0)
}

func main() {
  // Set timer
  timer := time.NewTimer(time.Minute * 2)
  go finished(timer)

  // Seed random number generator
  seed := rand.NewSource(time.Now().UnixNano())
  gen := rand.New(seed)

  var answer int
  for {
    // Generate 2 random numbers
    num1 := gen.Intn(99)
    num2 := gen.Intn(99)
    fmt.Printf("%d * %d = ", num1, num2)

    // Wait for input
    fmt.Scanf("%d", &answer)

    // Compare and process
    if (num1*num2 == answer) {
      score++
    } else {
      fmt.Printf("Wrong : %d\n", num1*num2)
    }
  }
}
