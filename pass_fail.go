// pass_fail tells us whether the user has passed the exam

package main

import (
    "bufio"; "fmt"; "log"; "os"; "strings"; "strconv")

func main() {                                       // The "main" function is called when the program starts.
    fmt.Print("Enter a grade: ")                    // Asking the user for a value.
    reader := bufio.NewReader(os.Stdin)             // Create a bufio.Reader to read keyboard input.
    input, err := reader.ReadString('\n')           // Reading the data entered by the user before pressing the Enter key.
    if err != nil {                                 // If an error occurs, display a message and terminate the program.
        log.Fatal(err)
    }
    input = strings.TrimSpace(input)                // Remove newline character from entered data.
    grade, err := strconv.ParseFloat(input, 64)     // Convert the input string to a float64 value (number).
    if err != nil {                                 // If an error occurs, display a message and terminate the program.
        log.Fatal(err)
    }
    var status string                               // The variable "status" is declared here so that it is in scope within the function.
    if grade >= 60 {                                // If grade is 60 or greater, the variable status is assigned the string "passing". Otherwise, the variable is assigned the string "failing"
        status = "passing"
    } else {
        status = "failing"
    }
    fmt.Println("A grade of", grade, "is", status)
}