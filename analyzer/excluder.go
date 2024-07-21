package analyzer

import (
	"bufio"
	"bytes"
	"flag"
	"go/token"
	"os"
	"strings"
	"sync"

	"github.com/eltonjr/json-interface-linter/internal/logger"
	"golang.org/x/tools/go/analysis"
)

var (
	excluderspath    string
	defaultExcluders = []string{}
	once             sync.Once
)

func RegisterFlags() {
	flag.StringVar(&excluderspath, "exclude", "", "file containing a list of interfaces to ignore")
}

func InitExcluders() {
	once.Do(func() {
		if excluderspath != "" {
			logger.Debugf("reading excluders from %s", excluderspath)
			d, err := ReadExcluders(excluderspath)
			if err != nil {
				logger.Debugf("failed to read excluders: %s", err)
			}

			defaultExcluders = d
		} else {
			logger.Debugf("no excluders file provided")
		}

		logger.Debugf("will ignore interfaces: %v", defaultExcluders)
	})
}

func ReadExcluders(path string) ([]string, error) {
	if path == "" {
		return nil, nil
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(buf))

	excluders := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if len(line) >= 2 && line[0] == '/' && line[1] == '/' {
			continue
		}

		excluders = append(excluders, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return excluders, nil
}

func reportOrExclude(pass *analysis.Pass, pos token.Pos, i, msg string) {
	logger.Debugf("found interface %s", i)
	if isExcluded(i, defaultExcluders) {
		logger.Debugf("found interface %s ignored by exclusion list", i)
		return
	}

	pass.Reportf(pos, msg)
}

func isExcluded(i string, excluders []string) bool {
	// for small slices, linear search is faster than map
	for _, m := range excluders {
		if strings.EqualFold(m, i) {
			return true
		}
	}
	return false
}
