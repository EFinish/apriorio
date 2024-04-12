package main

import "testing"

func TestStringInSlice(t *testing.T) {
	tests := []struct {
		str  string
		list []string
		want bool
	}{
		{"foo", []string{"foo", "bar", "baz"}, true},
		{"qux", []string{"foo", "bar", "baz"}, false},
	}

	for _, test := range tests {
		got := stringInSlice(test.str, test.list)

		if got != test.want {
			t.Errorf("stringInSlice(%q, %v) = %v; want %v", test.str, test.list, got, test.want)
		}
	}
}
