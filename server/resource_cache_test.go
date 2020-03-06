package server

import (
	"testing"
)

func TestCacheCreation(t *testing.T) {
	files, _ := SearchFiles("../tests/app")
	cache := NewCache(files)

	if cache.Size() != len(files) {
		t.Errorf("Expected %v, actual %v", files, cache)
	}
}

func TestCacheGetResource(t *testing.T) {
	files, _ := SearchFiles("../tests/app")
	cache := NewCache(files)
	resource, _ := cache.GetResource("/index.html")

	if resource.Name != "/index.html" {
		t.Errorf("Expected %v, actual %v", "/index.html", resource.Name)
	}

	if resource.CompressionType == "" {
		t.Errorf("Expected %v, actual %v", "deflate", resource.CompressionType)
	}

	if resource.Type != "text/html; charset=utf-8" {
		t.Errorf("Expected %v, actual %v", "text/html; charset=utf-8", resource.Type)
	}
}

func TestCacheGetResourceNotExists(t *testing.T) {
	files := []ResourceFile{
		ResourceFile{
			AbsPath: "/fake-file",
			RelPath: "/fake-file",
		},
	}

	cache := NewCache(files)
	_, err := cache.GetResource("/index.html2")
	if err == nil {
		t.Errorf("Expected error, actual %v", err)
	}
}
