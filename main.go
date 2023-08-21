package main

import (
	"github.com/jacqw8/demo/lines"
	"github.com/jacqw8/demo/print"
)

// imports from github
// I also tried a version making packages locally (rather than on github)
// Is it required that the imports are from a github repository
// for this tutorial?

// Wording on the tutorial:
// go mod init github.com/demo
// You should replace "" with your GitHub username.
// This will create a new Go module with the files "go.mod" and "go.sum".

// I think the wording above in the tutorial is a little confusing,
// especially since there no empty ""

func main() {
	print.Print()
	lines.CountLines()
}
