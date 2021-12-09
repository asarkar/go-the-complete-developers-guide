package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	f := "lorem_ipsum.txt"
	buf, err := readFile(f)
	if err != nil {
		t.Errorf("%v", err)
	}
	if buf == nil {
		t.Errorf("Faled to read file: %s", f)
	}
}
