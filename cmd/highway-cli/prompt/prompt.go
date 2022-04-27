package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ForString(text string, isValid func(string) bool) string {
	var in string
	r := bufio.NewReader(os.Stdin)

	var valid bool
	for {
		fmt.Print(text)
		in, _ = r.ReadString('\n')
		in = strings.Trim(in, "\n")
		if valid = isValid(in); !valid {
			fmt.Println("Invalid value.")
		} else {
			break
		}
	}
	return in
}
