package di

import (
	"fmt"
	"io"
	"os"
)

func GreetToConsole() {
	Greet(os.Stdout, "Thiago")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
