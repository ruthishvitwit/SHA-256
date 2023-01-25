package main

import (
	"fmt"
	"testing"
)

func TestSha256TableDriven(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{"abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{"1234", "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4"},
		{"name", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
	}
	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := Sha256(tt.s)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sha256("hello world")
	}
}
