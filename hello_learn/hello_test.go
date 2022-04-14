package main

import "testing"

/*func TestHello(t *testing.T) {
	got := Hello("Pasham")
	want := "Hello, Pasham"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}*/

func TestHello(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		if got != want{	
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Pasham","")
		want := "Hello, Pasham"	
		assertCorrectMessage(t,got,want)
	})
	t.Run("say 'Hello, World' when empy string is supplied",func(t *testing.T){
		got := Hello("","")
		want := "Hello, world"
		assertCorrectMessage(t,got,want)
	})
	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Pasham", "Spanish")
		want := "Hola, Pasham"
		assertCorrectMessage(t,got,want)
	})
	t.Run("in French",func(t *testing.T){
		got := Hello("Pasham", "French")
		want := "Bonjour, Pasham"
		assertCorrectMessage(t,got,want)
	})
}

