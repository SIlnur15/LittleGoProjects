// guess is a game in which the player must guess a random number.

package main

import ("fmt"; "math/rand"; "bufio"; "log"; "os"; "strconv"; "strings")

func main() {
	target := rand.Intn(100) + 1 							// Generating an integer from 1 to 100
	reader := bufio.NewReader(os.Stdin) 						// Creating a bufio.Reader to read keyboard input.
	var guesses int
	success:=true 									// Setting the program to display a loss message by default.
	for guesses = 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "attempts")
		fmt.Print("Make a guess: ") 						// Requesting a number
		input, err := reader.ReadString('\n') 					// Reading the data entered by the user before pressing the Enter key.
		if err != nil { 							// If an error occurs, display a message and terminate the program.
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) 					// Remove newline character from entered data.
		guess, err := strconv.Atoi(input) 					// Convert the input string to an integer.
		if err != nil { 										// If an error occurs, display a message and terminate the program.
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {		
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("please take congratulations !!!")
			success=false
			break 								// Exit from the loop.
		}
	}
	if success {
		fmt.Println("I'm sorry, but you loose *_*, it was ", target)
	}
}

