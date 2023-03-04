package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	//res := count(b, false, false)
	res := count(b, false)
	if res != exp {
		t.Errorf("Expected %d,got %d instead\n", exp, res)
	}

}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 \n line2 \nline3 word1")
	exp := 3
	res := count(b, true)
	//res := count(b, true, false)
	if exp != res {
		t.Errorf("Expected %d ,got %d instead.\n", exp, res)
	}
}

// func TestCountBytes(t *testing.T) {
// 	b := bytes.NewBufferString("gopher")
// 	exp := 6
// 	res := count(b, false, true)
// 	if exp != res {
// 		t.Errorf("Expected %d,got%d istead\n", exp, res)
// 	}
// }
