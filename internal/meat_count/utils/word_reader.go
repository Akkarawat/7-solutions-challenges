package utils

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

type WordReader struct {
	reader *bufio.Reader
	buffer bytes.Buffer
}

func NewWordReader(r io.Reader) *WordReader {
	return &WordReader{
		reader: bufio.NewReader(r),
	}
}

func (wr *WordReader) NextWord() (string, error) {
	wr.buffer.Reset()

	for {
		// Read next byte
		b, err := wr.reader.ReadByte()
		if err == io.EOF {
			if wr.buffer.Len() > 0 {
				return wr.buffer.String(), nil // Return last word if any
			}
			return "", io.EOF // No more words
		}
		if err != nil {
			return "", err
		}

		// Convert byte to rune for proper Unicode handling
		r := rune(b)

		// Check if character is valid (A-Z, a-z, or '-')
		if unicode.IsLetter(r) || r == '-' {
			wr.buffer.WriteRune(r) // Append to current word
		} else {
			// If we encounter a separator and have a word, return it
			if wr.buffer.Len() > 0 {
				return wr.buffer.String(), nil
			}
			// Otherwise, skip consecutive separators
			continue
		}
	}
}
