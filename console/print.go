package console

import (
	"fmt"
	"strings"
)

func Print(base string, args ...string) {
	finalString := base
	for _, arg := range args {
		finalString = strings.Replace(finalString, "?", arg, 1)
	}
	fmt.Printf("➜ %s %s\n", finalString, reset)
}
