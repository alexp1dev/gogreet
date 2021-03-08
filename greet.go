package greet

import (
	"fmt"
	"github.com/alex5p/gogreet/format"
)

func GreetingFor(name string) string {
	return fmt.Sprintf(format.GreetingFormat(), name)
}