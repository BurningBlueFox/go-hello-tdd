package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		spyOperation := &SpyCountdownOperation{}

		Countdown(spyOperation, spyOperation)

		want := []string{sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write}

		if !reflect.DeepEqual(want, spyOperation.Calls) {
			t.Errorf("wanted calls %v got %v", want, spyOperation.Calls)
		}
	})

	t.Run("print 3 2 1 go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spyOperation := &SpyCountdownOperation{}

		Countdown(buffer, spyOperation)

		got := buffer.String()
		want := `3
2
1
go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
