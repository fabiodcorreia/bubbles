package server

import (
	"testing"
)

func TestGetResourceContentSuccess(t *testing.T) {
	_, err := GetResourceContent("../internal/test-app/index.html")
	if err != nil {
		t.Errorf("Expected %v, actual %v", "Error", err)
	}
}

func TestGetResourceContentFileNotFound(t *testing.T) {
	_, err := GetResourceContent("../internal/test-app/fake.txt")
	if err == nil {
		t.Errorf("Expected %v, actual %v", "Error", err)
	}
}

func TestFileSeekerRootNotFound(t *testing.T) {
	_, err := SearchFiles("/tests/app")
	if err == nil {
		t.Errorf("Expected %v, actual %v", "Error", err)
	}
}

func TestFileSeekerHiddenFolder(t *testing.T) {
	_, err := SearchFiles("..")
	if err != nil {
		t.Errorf("Expected %v, actual %v", "nil", err)
	}
}

func TestFileSeekerSuccess(t *testing.T) {
	expected := []string{
		"/asset-manifest.json",
		"/favicon.ico",
		"/index.html",
		"/logo192.png",
		"/logo512.png",
		"/manifest.json",
		"/precache-manifest.260e71b40620c651caefbcfc5866fe4b.js",
		"/robots.txt",
		"/service-worker.js",
		"/static/css/main.d1b05096.chunk.css",
		"/static/css/main.d1b05096.chunk.css.map",
		"/static/js/2.a69980a4.chunk.js",
		"/static/js/2.a69980a4.chunk.js.LICENSE.txt",
		"/static/js/2.a69980a4.chunk.js.map",
		"/static/js/main.a76cbfe2.chunk.js",
		"/static/js/main.a76cbfe2.chunk.js.map",
		"/static/js/runtime-main.9976fff3.js",
		"/static/js/runtime-main.9976fff3.js.map",
		"/static/media/logo.5d5d9eef.svg"}
	files, _ := SearchFiles("../tests/app")

	for i, file := range files {
		if file.RelPath != expected[i] {
			t.Errorf("Expected %v, actual %v", file, expected[i])
		}
	}
}
