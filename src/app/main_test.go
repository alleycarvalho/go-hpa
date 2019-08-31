package main

import (
	"testing"
)

func TestCalculo(t *testing.T) {
	got := calculo(49)
	want := 7

	if got != want {
		t.Errorf("Raiz quadrada inválida: obteve %s e requer %s", got, want)
	}
}
