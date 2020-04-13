package main

import (
	"strings"
	"testing"
)

func Test_decode(t *testing.T) {
	tables := []struct {
		expected string
		input    string
	}{
		{"SOS", "... --- ..."},
		{"Hello world", ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."}}

	for _, table := range tables {
		decoded := decode(table.input)
		if !strings.EqualFold(decoded, table.expected) {
			t.Errorf("Invalid response, got: %s, want: %s.", decoded, table.expected)
		}
	}
}

func Test_encode(t *testing.T) {
	tables := []struct {
		expected string
		input    string
	}{
		{"... --- ...", "SOS"},
		{".... . .-.. .-.. --- / .-- --- .-. .-.. -..", "Hello world"},
		{".... . .-.. .-.. --- / .-- --- .-. .-.. -..", "Hello.world?"}}

	for _, table := range tables {
		encoded := encode(table.input)
		if !strings.EqualFold(encoded, table.expected) {
			t.Errorf("Invalid response, got: %s, want: %s.", encoded, table.expected)
		}
	}
}
