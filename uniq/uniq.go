// Uniq reads lines from stdin and writes unique lines to standard out.
// It does the stupid thing and holds everything it reads in memory.
// This was written, because the version of uniq in OS X 10.10.1 does
// not work properly.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	seen := map[string]struct{}{}

	var err error
	var line string
	for err == nil {
		line, err = reader.ReadString('\n')
		seen[line] = struct{}{}
	}

	writer := bufio.NewWriter(os.Stdout)
	for k, _ := range seen {
		_, err = writer.WriteString(k)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = writer.Write([]byte{'\n'})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
