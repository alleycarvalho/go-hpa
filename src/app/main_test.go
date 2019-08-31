package main

import (
	"math"
	"testing"
)

func TestCalculo(t *testing.T) {
	got := math.Sqrt(49)
	want := 7 float64

	if got != want {
		t.Errorf("Raiz quadrada inválida: obteve %s e requer %s", got, want)
	}
}
