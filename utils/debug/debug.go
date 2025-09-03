package debug

import (
	"fmt"
	"os"
	"strings"
)

// Debug is a method that display logs, it is useful for developer to trace program running
// details when troubleshooting
type Debug func(format string, v ...interface{})

var hookGetEnv = func() string {
	return os.Getenv("DEBUG")
}

var hookPrint = func(input string) {
	fmt.Println(input)
}

// Init returns a debug method that based the enviroment variable DEBUG value
func Init(flag string) Debug {
	enable := false

	env := hookGetEnv()
	parts := strings.Split(env, ",")
	for _, part := range parts {
		if part == flag {
			enable = true
			break
		}
	}

	return func(format string, v ...interface{}) {
		if enable {
			hookPrint(fmt.Sprintf(format, v...))
		}
	}
}
