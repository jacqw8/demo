package lines

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func Hi() {
	fmt.Println("HI new package!")
}

func CountLines() {
	// use flag to get filename from command line argument
	var filename string
	flag.StringVar(&filename, "file", "", "name of file")
	flag.Parse()
	// open file
	f, err := os.Open(filename)
	// check for errors
	if err != nil {
		fmt.Println("Error:", err)
	}
	// defer closing the file for end of function
	defer f.Close()
	// initialize scanner
	scanner := bufio.NewScanner(f)
	// initialize counter
	var count int
	for scanner.Scan() {
		count++
	}
	fmt.Println(filename, count)
}
