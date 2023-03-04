package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count Lines")
	//bytes := flag.Bool("b", false, "Counting bytes")
	flag.Parse()
	//flag.Parse() return the pointer to the bool,we need to dereference it in order use the value
	//hence we are passing the parameter as *lines
	//fmt.Println(lines, "lines")
	//fmt.Println(*lines, "pointer to line")
	//fmt.Println(count(os.Stdin, *lines, *bytes))
	fmt.Println(count(os.Stdin, *lines))
}

// func count(r io.Reader, countLines, countBytes bool) int {
func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		//if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	// if countBytes {
	// 	fmt.Println("inside countbytes")
	// 	fmt.Println(len(scanner.Bytes()))
	// 	return len(scanner.Bytes())
	// }

	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}
