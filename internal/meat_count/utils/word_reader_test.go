package utils

import (
	"io"
	"strings"
	"testing"
)

func testWordReader(t *testing.T, input string, expected []string) {
	reader := NewWordReader(strings.NewReader(input))

	var words []string
	for {
		word, err := reader.NextWord()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		words = append(words, word)
	}

	if len(words) != len(expected) {
		t.Fatalf("Expected %d words, got %d", len(expected), len(words))
	}
	for i, word := range expected {
		if words[i] != word {
			t.Errorf("Expected word '%s', got '%s'", word, words[i])
		}
	}
}

func TestWordReader_BasicSentence(t *testing.T) {
	input := "Hello world. This is, a test."
	expected := []string{"Hello", "world", "This", "is", "a", "test"}
	testWordReader(t, input, expected)
}

func TestWordReader_ConsecutiveSeparators(t *testing.T) {
	input := "Hello...world,,this  is   a test"
	expected := []string{"Hello", "world", "this", "is", "a", "test"}
	testWordReader(t, input, expected)
}

func TestWordReader_OnlySeparators(t *testing.T) {
	input := "   ... ,,   "
	expected := []string{}
	testWordReader(t, input, expected)
}

func TestWordReader_EmptyInput(t *testing.T) {
	input := ""
	expected := []string{}
	testWordReader(t, input, expected)
}

func TestWordReader_NewlinesAndSpaces(t *testing.T) {
	input := "Hello\nworld\tthis is  a test.\nNext line,"
	expected := []string{"Hello", "world", "this", "is", "a", "test", "Next", "line"}
	testWordReader(t, input, expected)
}

func TestWordReader_LargeInput(t *testing.T) {
	largeText := strings.Repeat("word, ", 100000) // 100,000 words
	reader := NewWordReader(strings.NewReader(largeText))

	count := 0
	for {
		_, err := reader.NextWord()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		count++
	}

	expectedCount := 100000
	if count != expectedCount {
		t.Errorf("Expected %d words, got %d", expectedCount, count)
	}
}
