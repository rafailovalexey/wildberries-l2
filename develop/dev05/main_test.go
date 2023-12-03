package main

import (
	"io"
	"os"
	"testing"
)

func TestGrep(t *testing.T) {
	application := &Application{}
	flags := &Flags{Context: 1}

	data := []string{
		"apple",
		"banana",
		"orange",
		"grape",
		"peach",
	}

	pattern := "an"

	expected := []string{
		"apple",
		"banana",
		"orange",
		"grape",
	}

	result, err := application.Grep(data, pattern, flags)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("expected length %d got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %s got %s", expected[i], v)
		}
	}
}

func TestMatch(t *testing.T) {
	application := &Application{}
	flags := &Flags{IgnoreCase: true}

	line := "Hello, World!"
	pattern := "hello, world!"

	result, err := application.Match(line, pattern, flags)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !result {
		t.Error("expected a match got no match")
	}

	flags.IgnoreCase = false

	result, err = application.Match(line, pattern, flags)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if result {
		t.Error("expected no match got a match")
	}
}

func TestPrintFileData(t *testing.T) {
	application := &Application{}
	flags := &Flags{LineNumbers: true}

	data := []string{"apple", "banana", "orange"}

	output := CaptureOutput(func() {
		application.PrintFileData(data, flags)
	})

	expected := "1:apple\n2:banana\n3:orange\n"

	if output != expected {
		t.Errorf("expected output %s got %s", expected, output)
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
