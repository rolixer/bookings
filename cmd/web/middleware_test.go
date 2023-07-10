package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh myHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Errorf("type is not http.Handler is %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var mh myHandler
	h := SessionLoad(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Errorf("type is not http.Handler is %T", v)
	}
}
