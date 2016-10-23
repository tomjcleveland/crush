// Package crush crushes Go source files into a single line.
package crush

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

// Bytes is a convenience function wrapping Reader()
func Bytes(src []byte) ([]byte, error) {
	r := bytes.NewReader(src)
	out, err := Reader(r)
	if err != nil {
		return nil, err
	}
	outBytes, err := ioutil.ReadAll(out)
	if err != nil {
		return nil, err
	}
	return outBytes, nil
}

// String is a convenience function wrapping Reader()
func String(src string) (string, error) {
	outBytes, err := Bytes([]byte(src))
	return string(outBytes), err
}

// Reader crushes the provided Go source file into a
// single line, while still conforming to proper Go syntax.
func Reader(src io.Reader) (io.Reader, error) {
	out := bytes.NewBuffer(nil)
	s := bufio.NewScanner(src)
	for s.Scan() {
		line := s.Bytes()
		line = bytes.TrimSpace(line)

		// If line is empty, skip it
		if len(line) == 0 {
			continue
		}

		// If line is one-line comment, skip it
		if len(line) >= 2 && string(line[:2]) == "//" {
			continue
		}

		// If line starts multi-line comment, ignore block
		if bytes.Contains(line, []byte("/*")) {
			if err := consumeComment(s, line, out); err != nil {
				return nil, err
			}
			continue
		}

		writeLine(line, out)
	}
	return out, nil
}

// writeLine writes the give line, deciding whether or not to
// terminate it with a semicolon.
func writeLine(line []byte, out *bytes.Buffer) {
	out.Write(line)
	last := line[len(line)-1]

	// If last character is a comma, don't add semicolon
	if last == ',' {
		return
	}

	// If the last character is an open bracket,
	// don't add a semicolon
	if last == '{' || last == '(' || last == '[' {
		return
	}

	out.WriteByte(';')
}

// consumeComment consumes block comments until it finds
// a terminal symbol ('*/') on a line without another opening
// symobol ('/*') after it. It writes all non-comment tokens
// that occur *between* close- and open-comment symbols on the
// same line.
func consumeComment(s *bufio.Scanner, currLine []byte, out *bytes.Buffer) error {
	// Print anything before the comment starts
	beginning := bytes.SplitN(currLine, []byte("/*"), 2)[0]
	if strings.TrimSpace(string(beginning)) != "" {
		writeLine(beginning, out)
	}

	// Consume all block comments ending in the same line
	var done bool
	currLine, done = consumeSameLineComment(currLine)
	if done {
		writeLine(currLine, out)
		return nil
	}

	// Otherwise, scan through lines until you find `*/`
	for s.Scan() {
		line := s.Bytes()

		if bytes.Contains(line, []byte("*/")) {
			remainder := bytes.SplitN(line, []byte("*/"), 2)[1]
			if string(bytes.TrimSpace(remainder)) == "" {
				return nil
			}
			if bytes.Contains(remainder, []byte("/*")) {
				return consumeComment(s, remainder, out)
			}
			writeLine(remainder, out)
			return nil
		}
	}

	return errors.New("got EOF, expecting '*/'")
}

func consumeSameLineComment(currLine []byte) (clean []byte, done bool) {
	for bytes.Contains(currLine, []byte("*/")) {
		currLine = bytes.SplitN(currLine, []byte("*/"), 2)[1]
		if !bytes.Contains(currLine, []byte("/*")) {
			return currLine, true
		}
	}
	return currLine, false
}
