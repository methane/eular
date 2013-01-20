package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func score(s string) (score int) {
	for _, c := range s {
		score += int(c) - int('A') + 1
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic("Can't read a line.");
	}
	line = strings.Trim(line, "\" \n\r")
	names := strings.Split(line, "\",\"")
	sort.Strings(names)

	total_score := 0
	for pos, word := range names {
		total_score += (pos + 1) * score(word)
	}
	fmt.Printf("%v\n", total_score)
}
