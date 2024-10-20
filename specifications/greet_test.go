package specifications

import (
	"testing"
)

func TestGreet(t *testing.T) {
	GreetSpecification(t, GreetAdapter(Greet))
}
