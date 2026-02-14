// The Attendance Sheet (Easy)
// Goal: Check if a key exists using the "comma ok" idiom.

// Create a map of student names and their "Present" status (bool).

// Write a function that takes a name and checks if that student is in the system.

// If they are, print their status; if not, print "Student not found."

package main

import "fmt"

func isPresent(attendanceList map[string]bool, studentName string) {
	// The "comma ok" idiom: 'ok' is a boolean that tells us if the key exists
	status, ok := attendanceList[studentName]

	if ok {
		fmt.Printf("Student %s found. Status: %t\n", studentName, status)
	} else {
		fmt.Printf("Student %s not found.\n", studentName)
	}
}
