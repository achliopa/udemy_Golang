// Create a program that reads the contents of a text file
// then prints its contents to the terminal

// The file to open should be provided as an argument to the
// program when its executed at the terminal. e.g `go run main.go myfile.txt`
// shoudl open up myfile.txt file

// To read the args provided to a program we can reference the variable os.Args
// which is a slice of type string

// To open a file we can use the Open function at os pkg

// What interfaces does th File ype impelement?

// If the File type implements the Reader interface, we might be able to reuse the io.Copy func.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
}
