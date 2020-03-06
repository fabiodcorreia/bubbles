package server

import (
	"testing"
)

func TestInMemoryResourceString(t *testing.T) {
	var content = []byte("File Content to Test aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	var fr, _ = NewInMemoryResource("index.html", content, "deflate")
	var str = "{ name: 'index.html', size: 30, initial-size: 77, contentType: 'text/html; charset=utf-8', compressionType: 'deflate' }"

	if str != fr.String() {
		t.Errorf("Expected %s, actual %v", str, fr)
	}
}

func Benchmark(b *testing.B) {
	var content = []byte("File Content to Test")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewInMemoryResource("index.html", content, "deflate")
	}
}
