package hello

import (
	"strings"
)

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	return "Hello, " + strings.Join(names, ", ") + "!"
	//return fmt.Sprintf("Hello, %s!", strings.Join(names, ", "))
}
