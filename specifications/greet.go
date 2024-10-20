package specifications

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func Greet(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("John")
	assert.NoError(t, err)
	assert.Equal(t, "Hello, John!", got)
}
