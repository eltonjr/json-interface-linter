package marshal

import (
	"bufio"
	"bytes"
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
		line := scanner.Text()
		// Skip comments and empty lines.
		if strings.HasPrefix(line, "//") || line == "" {
			continue
		}

		ss := strings.Split(line, "[")
		var argIndex int
		if len(ss) > 1 {
			i, err := strconv.Atoi(strings.TrimRight(ss[1], "]"))
			if err != nil {
				return nil, err
			}
			argIndex = i
		}

		marshalers = append(marshalers, marshaler{
			functionPath: ss[0],
			argument:     argIndex,
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
