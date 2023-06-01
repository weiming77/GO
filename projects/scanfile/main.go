package main

import (
	"bufio"
	"os"
	"regexp"
)

var cmdRe = regexp.MustCompile(`;go ([a-z]+)`)

// cmdFreg returns the frequency of "go" subcommand usage in ZSH history
func cmdFreg(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fregs := make(map[string]int)
	s := bufio.NewScanner(file)
	for s.Scan() {
		matches := cmdRe.FindStringSubmatch(s.Text())
		if len(matches) == 0 {
			continue
		}
		cmd := matches[1]
		fregs[cmd]++
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return fregs, nil
}

func main() {
	cmdFreg()
}
