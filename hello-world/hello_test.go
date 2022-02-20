package main

import "testing"

func TestHello(t *testing.T){
	resultado := Hello("Gabriel")
	esperado := "Hello Gabriel!"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}