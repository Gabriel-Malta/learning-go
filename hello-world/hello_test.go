package main

import "testing"

func TestHello(t *testing.T){
	verificationRightMessage := func(t *testing.T, resultado, esperado string){
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("Says Hello to people", func(t *testing.T){
		resultado := Hello("Gabriel")
		esperado := "Hello, Gabriel!"

		verificationRightMessage(t, resultado, esperado)
	})

	t.Run("Says Hello World when the name string is empty", func(t *testing.T){
		resultado := Hello("")
		esperado := "Hello, World!"

		verificationRightMessage(t, resultado, esperado)
	})
}