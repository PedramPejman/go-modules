package main

import (
  "fmt"
	"math"
  "math/rand"
  "os"
	"os/exec"
  "os/signal"
	"strings"
  "time"
)

/*
 * Dumps i '_' characters, hits carriage return and 
 * writes an i digit number from stdin to answer
 */
func receive(answer *int, i int) {
  // Disable input buffering
  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
  // Do not display entered characters on the screen
  exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
  var b []byte= make([]byte, 1)
  os.Stdin.Read(b)
  // Restore the echoing state when exiting
  exec.Command("stty", "-F", "/dev/tty", "echo").Run()
  // Flush screen
  fmt.Printf("\r%s\r", strings.Repeat("_", i))
  // Wait for number
  fmt.Scanf("%d", answer)
}

func main() {
  // Register signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for _ = range c {
			// Restore the echoing state when exiting
      fmt.Println()
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
      os.Exit(0)
		}
	}()

  // Start timer
  start := time.Now()

  // Seed random number generator
  seed := rand.NewSource(time.Now().UnixNano())
  gen := rand.New(seed)

  var answer int
  for i := 1; i <= 10; i++ {
    // Generate i-digit random number
    num := gen.Intn((int) (math.Pow10(i) - math.Pow10(i-1))) + (int) (math.Pow10(i-1))
    // Print number on screen
    fmt.Printf("\n%d", num)
    // Receive answer
    receive(&answer, i)
    // Compare and retry if necessary
		for answer != num {
      fmt.Println("\rWrong, try again")
      fmt.Printf("%d", num)
      receive(&answer, i)
    }
  }

  fmt.Printf("%s\n", time.Now().Sub(start))

}
