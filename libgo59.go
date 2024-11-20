package execShell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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
