package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myHandler myHandler
	h := NoSurf(&myHandler)

	switch h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error("type is not http.Handler")
	}
}

func TestSessionLoad(t *testing.T) {
	var myHandler myHandler
	h := SessionLoad(&myHandler)

	switch h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error("type is not http.Handler")
	}
}
