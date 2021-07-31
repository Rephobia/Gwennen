package helloworld


import (
    "testing"
)

func TestHelloNotEmpty(t *testing.T) {

	if hello() == "" {
		t.Error("hello empty")
	}
}

func TestWorldNotEmpty(t *testing.T) {

	if world() == "" {
		t.Error("world empty")
	}
}
