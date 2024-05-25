package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Have name", func(t *testing.T) {
		got := Hello("Neyakki", "")
		want := "Hello, Neyakki!"

		assertEqualMsg(t, got, want)
	})

	t.Run("Have not name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertEqualMsg(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Neyakki", "Spanish")
		want := "Hola, Neyakki!"

		assertEqualMsg(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Neyakki", "French")
		want := "Bonjour, Neyakki!"

		assertEqualMsg(t, got, want)
	})
}

func assertEqualMsg(t testing.TB, got string, want string) {
	t.Helper() // Необходим, чтобы сообщить набору тестов, что этот метод является вспомогательным.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
