package util

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		d: "foo", e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		d: "foo", e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

/*
  How to run benchmark:
  Running tool #1: go test -benchmem -run=^$ -bench ^BenchmarkMemoryStack$ awsdisc/client/util
  Running tool #2: go test ./... -bench=BenchmarkMemoryStack -benchmem -run=^$ -count=10 > stack.txt && benchstat stack.txt
*/
func BenchmarkMemoryStack(b *testing.B) {
	var s S

	f, err := os.Create("stack.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		s = byCopy()
	}

	_ = fmt.Sprintf("%+v", s)
}

/*
  How to run benchmark:
  Running tool #1: go test -benchmem -run=^$ -bench ^BenchmarkMemoryStack$ awsdisc/client/util
  Running tool #2: go test ./... -bench=BenchmarkMemoryHeap -benchmem -run=^$ -count=10 > heap.txt && benchstat heap.txt
*/
func BenchmarkMemoryHeap(b *testing.B) {
	var s *S

	f, err := os.Create("heap.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		s = byPointer()
	}

	_ = fmt.Sprintf("%+v", s)
}
