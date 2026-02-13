// The Short-Circuit: Use the "short statement" syntax (if err := doSomething(); err != nil) to check if a boolean variable isError is true.

package main

func checkIfError(isError bool) bool {
	// We initialize 'err' and then check if it is true
	if err := isError; err == true {
		return true
	} else {
		return false
	}
}
