package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		os.Exit(1)
	}
	defer f.Close()
	b := make([]byte, 100)
	n, err := f.Read(b)

	fmt.Printf("Chars:%d % x\n", n, b) /// in Hex
	fmt.Printf("Chars:%d % c\n", n, b) // in Char
	stringVersion := string(b)
	fmt.Printf("Chars:%d %s\n", n, stringVersion)

	fmt.Printf("============================\n")

	// conversion to byte string
	//someString := "Foo bar"
	//f.write([]byte(someString))

	//f.Write([]byte(stringVersion))

	/*Anthonys-MacBook-Pro:byeSlices mcclayac$ godoc os WriteString
	type File struct {
		// contains filtered or unexported fields
	}

	File represents an open file descriptor.

	func (f *File) WriteString(s string) (n int, err error)
	WriteString is like Write, but writes the contents of string s rather
	than a slice of bytes.


		Anthonys-MacBook-Pro:byeSlices mcclayac$ godoc os ReadString
	No match found.

		Anthonys-MacBook-Pro:byeSlices mcclayac$ godoc os Read
	type File struct {
		// contains filtered or unexported fields
	}

	File represents an open file descriptor.

	func (f *File) Read(b []byte) (n int, err error)
	Read reads up to len(b) bytes from the File. It returns the number of
	bytes read and any error encountered. At end of file, Read returns 0,
		io.EOF.
	*/

}
