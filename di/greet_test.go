package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Thiago")

	got := buffer.String()
	want := "Hello, Thiago"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
