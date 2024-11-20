package execShell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"math/rand"
  "time"
  "path/filepath"
	"io/ioutil"
)

// This function takes a string of text as input and returns a new string with any blank or empty lines removed.
// The purpose of this function is to remove any unnecessary whitespace from command output when printing to the console.
func Strip(output string) string {
	// Split the input string into an array of strings, one for each line of text
	lines := strings.Split(output, "\n")

	// Initialize an empty array of strings that will store the non-empty lines
	var nonEmptyLines []string

	// Iterate over the array of lines
	for _, line := range lines {
		// Check if the line is not empty or blank
		if strings.TrimSpace(line) != "" {
			// If the line is non-empty, add it to the array of non-empty lines
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	// Join the non-empty lines back together with a newline character in between each line
	return strings.Join(nonEmptyLines, "\n")
}

// This function runs a shell command with the given arguments and handles any errors.
// It takes 3 parameters:
// 1. command: the command to run
// 2. suppressOutput: a boolean indicating whether to suppress the output of the command
// 3. args: a variable number of arguments to pass to the command
func ExecShell(command string, suppressOutput bool, args ...string) {
	// Create a new command to run
	cmd := exec.Command(command, args...)

	// Run the command and capture its output and any errors
	output, err := cmd.CombinedOutput()

	// Check if there was an error
	if err != nil {
		// If suppressOutput is false, print the command output even if there was an error
		if !suppressOutput {
			// Print the command output and any error message
			fmt.Println("Command output:", Strip(string(output)))
		}

		// Print the error message
		fmt.Println("Error:", err)

		// Return to prevent further execution of the function
		return
	}

	// If there was no error and suppressOutput is false, print the command output
	if !suppressOutput {
		// Print the command output, stripped of any blank or empty lines
		fmt.Println(Strip(string(output)))
	}
}

// This function checks if a file exists in the file system.
// It takes 1 parameter:
// 1. filename: the name of the file to check for
// It returns a boolean indicating whether the file exists.
func FileExists(filename string) bool {
    // os.Stat() returns information about the file, or an error if the file does not exist
    // os.IsNotExist() checks if the error returned by os.Stat() indicates that the file does not exist
    // The ! operator negates the result of os.IsNotExist(), so if the error indicates that the file does not exist, the function returns false
    // Otherwise, the function returns true
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

// This function generates a random string of a given length.
// It takes 1 parameter:
// 1. length: the length of the string to generate
// It returns a string of the given length, consisting of random characters.
func StringRNG(length int) string {
    // The characters that can be used in the generated string
    // This string includes all alphanumeric characters, both uppercase and lowercase
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // Seed the random number generator with the current time
    // This is necessary because the random number generator is not truly random,
    // but rather a pseudorandom number generator that uses a seed value
    // to generate a sequence of numbers
    rand.Seed(time.Now().UnixNano())

    // Create a byte array of the given length
    // This array will be used to store the generated string
    result := make([]byte, length)

    // Iterate over the array and fill it with random characters
    // The characters are chosen from the charset string
    for i := range result {
        // Get a random index from the charset string
        // This index is used to select a character from the charset
        // The character is then assigned to the current index in the result array
        result[i] = charset[rand.Intn(len(charset))]
    }

    // Convert the byte array to a string and return it
    // This is the final generated string
    return string(result)
}

// Simple filepath.Join wrapper I guess, its so simple I won't even explain it. 
// Join("a", "b", "c") -> "a/b/c" AS A EXAMPLE
func Join(parts ...string) string {
    return filepath.Join(parts...)
}

// Function to read a file and return its contents as a string
func Read(filePath string) (string, error) {
    // Read the file
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }
    return string(content), nil
}
