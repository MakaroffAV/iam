package main

import "testing"

func BenchmarkSepJoin(t *testing.B) {
	for i := 0; i < 1e+6; i++ {
		sepJoin()
	}
}

func BenchmarkBuiltinJoin(t *testing.B) {
	for i := 0; i < 1e+6; i++ {
		builtinJoin()
	}
}
