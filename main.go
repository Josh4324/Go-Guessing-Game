package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number")

	// generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New((source))
	secretNumber := randomizer.Intn(100) // generates number between 0 and 100

	var guess int
	guessnum := 5
	fmt.Println("You have 5 tries to guess the correct number")
	for guessnum > 0 {

		fmt.Println("Please input your guess")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("Too Big")
		} else if guess < secretNumber {
			fmt.Println("Too small")
		} else {
			fmt.Println("You Win!")
			break
		}
		guessnum = guessnum - 1
	}

	if guessnum == 0 {
		fmt.Println("You lose!")
	}

}
