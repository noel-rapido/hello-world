package main

import (
	"bytes"
	"testing"
)

func TestWriteResponse(t *testing.T) {
	var buff = bytes.NewBuffer([]byte{})
	expected := "hello world"
	WriteResponse(buff, "Hello World")

	responseOutput := buff.String()

	// test if output= input
	if responseOutput != expected {
		t.Fatalf("unexpected response received.\n  expected: %v\n  received: %v", expected, responseOutput)
	}

}
