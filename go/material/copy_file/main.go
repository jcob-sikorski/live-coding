// The Safe File Copier (Level: Medium)
// The Goal: Practice real-world resource management and error handling.

// Task: Write a function CopyFile(dstName, srcName string) error.

// Open the source file.

// Defer the closing of the source file.

// Create the destination file.

// Defer the closing of the destination file.

// Use io.Copy to move the data.

// Bonus: Ensure that if the destination file creation fails, the source file still gets closed properly.

// The Safe File Copier (Level: Medium)
// The Goal: Practice real-world resource management and error handling.

// Task: Write a function CopyFile(dstName, srcName string) error.

// Open the source file.

// Defer the closing of the source file.

// Create the destination file.

// Defer the closing of the destination file.

// Use io.Copy to move the data.

// Bonus: Ensure that if the destination file creation fails, the source file still gets closed properly.

package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) error {
	// 1. Open the source file
	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}
	// This ensures src is closed when CopyFile returns
	defer src.Close()

	// 2. Create the destination file
	dst, err := os.Create(dstName)
	if err != nil {
		// Bonus: src.Close() will still run here because it was deferred above!
		return fmt.Errorf("failed to create destination: %w", err)
	}
	// This ensures dst is closed when CopyFile returns
	defer dst.Close()

	// 3. Use io.Copy to move the data
	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("error during copy: %w", err)
	}

	// 4. Flush and sync (Optional but good practice for "Safe" copying)
	return dst.Sync()
}

func main() {
	err := CopyFile("destination.txt", "source.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File copied successfully!")
	}
}
