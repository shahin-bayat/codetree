package main

import "testing"

func Test_countLines(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{name: "Test 1", args: args{path: "mock/dummy1.txt"}, expected: 150},
	}

	for _, tt := range tests {
		if got, err := countLines(tt.args.path); got != tt.expected || err != nil {
			t.Errorf("countLines(%s) = %v, want %v", tt.args.path, got, tt.expected)
		}
	}
}
