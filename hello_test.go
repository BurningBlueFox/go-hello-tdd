package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Thiago", "")
		want := "Hello, Thiago!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying Hello, World! when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In portuguese", func(t *testing.T) {
		got := Hello("Thiago", "portuguese")
		want := "Olá, Thiago!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Thiago", "spanish")
		want := "Hola, Thiago!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In french", func(t *testing.T) {
		got := Hello("Thiago", "french")
		want := "Bonjour, Thiago!"
		assertCorrectMessage(t, got, want)
	})

}
