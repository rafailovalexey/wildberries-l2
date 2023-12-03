package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCutInputStrings(t *testing.T) {
	input := "apple\torange\tbanana\n1\t2\t3\n"
	expected := "1\t2\t3\n"

	flags := &Flags{Fields: "1", Delimiter: "\t", Separated: false}

	scanner := bufio.NewScanner(strings.NewReader(input))

	application := &Application{}

	output := CaptureOutput(func() {
		application.CutInputStrings(scanner, flags)
	})

	if output != expected {
		t.Errorf("")
	}
}

func TestParseNumberColumns(t *testing.T) {
	input := "42"
	expected := 42

	application := &Application{}

	result := application.ParseNumberColumns(input)

	if result != expected {
		t.Errorf("expected result %d got result %d", expected, result)
	}
}

func CaptureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	f()

	w.Close()

	out, _ := io.ReadAll(r)

	os.Stdout = old

	return string(out)
}
