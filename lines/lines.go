package lines

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func CountLines() {
	// use flag to get filename from command line argument
	var filename string
	// the command will look like "go run main.go -file lines/lines.go"
	// where the name of file goes after the flag "-file"
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
