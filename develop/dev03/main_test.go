package main

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	files := &Files{}

	wd, _ := os.Getwd()

	filepath, err := files.GetFilePath("temporary/test.txt")
	expected := path.Join(wd, "temporary/test.txt")

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	if filepath != expected {
		t.Errorf("expected %s got %s\n", expected, filepath)
	}

	_, err = files.GetFilePath("nonexistentdir/test.txt")

	if err == nil {
		t.Error("expected an error but got nil\n")
	}
}

func TestGetSortedStringsWithArgumentNumeric(t *testing.T) {
	application := &Application{}
	flags := &Flags{Numeric: true}

	data := []string{"3", "1", "2", "4"}

	result, err := application.GetSortedStringsWithArguments(data, flags)

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	expected := []string{"1", "2", "3", "4"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v\n", expected, result)
	}
}

func TestGetSortedStringsWithArgumentReverse(t *testing.T) {
	application := &Application{}
	flags := &Flags{Reverse: true}

	data := []string{"1", "2", "3", "4"}

	result, err := application.GetSortedStringsWithArguments(data, flags)

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	expected := []string{"4", "3", "2", "1"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v\n", expected, result)
	}
}

func TestGetSortedStringsWithArgumentIgnoreTrailingMonth(t *testing.T) {
	application := &Application{}
	flags := &Flags{Month: true}

	data := []string{
		"september",
		"august",
		"july",
		"june",
		"may",
		"april",
		"march",
		"february",
		"january",
	}

	result, err := application.GetSortedStringsWithArguments(data, flags)

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	expected := []string{
		"january",
		"february",
		"march",
		"april",
		"may",
		"june",
		"july",
		"august",
		"september",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v\n", expected, result)
	}
}

func TestGetSortedStringsWithArgumentIgnoreTrailingSpace(t *testing.T) {
	application := &Application{}
	flags := &Flags{IgnoreTrailingSpace: true}

	data := []string{"1", "2", "3", "4               "}

	result, err := application.GetSortedStringsWithArguments(data, flags)

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	expected := []string{"1", "2", "3", "4"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v\n", expected, result)
	}
}

func TestGetSortedStringsWithArgumentColumnKey(t *testing.T) {
	application := &Application{}
	flags := &Flags{ColumnKey: 2}

	data := []string{"1 9", "2 8", "3 7", "4 6"}

	result, err := application.GetSortedStringsWithArguments(data, flags)

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	expected := []string{"6", "7", "8", "9"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v\n", expected, result)
	}
}

func TestGetSortedStringsWithArgumentUnique(t *testing.T) {
	application := &Application{}
	flags := &Flags{Unique: true}

	data := []string{"1", "2", "3", "4", "4"}

	result, err := application.GetSortedStringsWithArguments(data, flags)

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	expected := []string{"1", "2", "3", "4"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v\n", expected, result)
	}
}

func TestGetSortedStringsWithArgumentCheck(t *testing.T) {
	application := &Application{}
	flags := &Flags{Check: true}

	data := []string{"1", "2", "3", "4"}

	_, err := application.GetSortedStringsWithArguments(data, flags)

	if err == nil {
		t.Errorf("unexpected error %v\n", err)
	}
}
