package marshal

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strconv"
	"strings"
)

// marshaler represent which function is used to marshal a value.
// an argument index can be used in case the function receives many arguments.
type marshaler struct {
	// package.function
	functionPath string
	// 0-based
	argument int
}

var defaultMarshalers = []marshaler{
	{
		functionPath: "encoding/json.Marshal",
	},
	{
		functionPath: "encoding/json.MarshalIndent",
	},
	{
		functionPath: "encoding/json.Encode",
	},
}

var ErrMissingClosingBracket = errors.New("missing closing bracket ']'")

// ReadMarshalers reads an marshalers file, a newline delimited file that lists
// function calls to be checked
//
// Lines that start with two forward slashes are considered comments and are ignored.
// Use [] to represent argument index, 0-based.
//
// example:
// a function called this way mypkg.Marshal(ctx, bytes)
// should be represented as mypkg.Marshal[1]
func ReadMarshalers(path string) ([]marshaler, error) {
	if path == "" {
		return defaultMarshalers, nil
	}

	var marshalers []marshaler

	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(buf))

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		if len(line) >= 2 && line[0] == '/' && line[1] == '/' {
			continue
		}

		stringSize := bytes.IndexByte(line, '[')
		if stringSize == -1 {
			marshalers = append(marshalers, marshaler{
				functionPath: string(line),
			})
			continue
		}

		argSize := bytes.IndexByte(line[stringSize:], ']')
		if argSize == -1 {
			return nil, ErrMissingClosingBracket
		}

		i, err := strconv.Atoi(string(line[stringSize+1 : stringSize+argSize]))
		if err != nil {
			return nil, err
		}

		marshalers = append(marshalers, marshaler{
			functionPath: string(line[:stringSize]),
			argument:     i,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return marshalers, nil
}

func isMarshaler(path string, marshalers []marshaler) (marshaler, bool) {
	// for small slices, linear search is faster than map
	for _, m := range marshalers {
		if strings.EqualFold(m.functionPath, path) {
			return m, true
		}
	}
	return marshaler{}, false
}
