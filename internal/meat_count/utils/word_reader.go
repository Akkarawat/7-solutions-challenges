package utils

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

// WordReader efficiently reads words from an io.Reader
type WordReader struct {
	reader *bufio.Reader
	buffer bytes.Buffer
}

// NewWordReader initializes a WordReader with an io.Reader
func NewWordReader(r io.Reader) *WordReader {
	return &WordReader{
		reader: bufio.NewReader(r),
	}
}

// NextWord reads the next word from the stream
func (wr *WordReader) NextWord() (string, error) {
	// Clear buffer from previous word
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

		// Check if it's a word separator (whitespace, full stop, semicolon)
		if unicode.IsSpace(rune(b)) || b == '.' || b == ',' {
			if wr.buffer.Len() > 0 {
				return wr.buffer.String(), nil // Return the current word
			}
			continue // Skip consecutive separators
		}

		// Append character to buffer
		wr.buffer.WriteByte(b)
	}
}
