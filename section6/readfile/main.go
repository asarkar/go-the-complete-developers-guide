/*
Create a program that reads the contents of a text file then prints its contents
to the terminal.

The file to open should be provided as an argument to the program when it is executed
at the terminal. For example, 'go run main.go myfile.txt' should open up the 'myfile.txt'
file.

To read in the arguments provided to the program, you can reference the variable os.Args,
which is a slice of type string.

To open a file, check out the documentation for the 'Open' function in the 'os' package.

What interfaces does the 'File' type implement?

if the 'File' type implements the 'Reader' interface, you might be able to reuse that
'io.Copy' function!
*/
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	name, err := readFileName()
	if err != nil {
		log.Fatal(err)
	}
	buf, err := readFile(*name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func readFileName() (*string, error) {
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("usage: main.go <filename>")
	}
	return &os.Args[1], nil
}

func readFile(name string) (*bytes.Buffer, error) {
	/* Open file */
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	/* Get size */
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	/* Grow to size of file plus extra slop to ensure no realloc. */
	buf.Grow(int(fi.Size()) + bytes.MinRead)

	/* ReadFrom(io.Reader) allocates an initial []byte of some size and reads into that slice.
	If there's more data in the reader than space in the slice, then the function
	allocates a larger slice, copies existing data to that slice and reads more.
	This repeats until EOF.
	To avoid this, we allocate a slice big enough.

	An alternative implementation is to use a 'strings.Builder' instead of 'bytes.Buffer'
	with the 'io.Copy' function, like so: io.Copy(buf, f), but that seems to use more
	memory. See https://stackoverflow.com/q/64830650/839733.
	*/
	_, err = buf.ReadFrom(f)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
