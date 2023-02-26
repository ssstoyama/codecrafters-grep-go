package main

import "testing"

func TestMatchLine(t *testing.T) {
	tests := []struct {
		line    []byte
		pattern string
		ok      bool
	}{
		{
			line:    []byte("apple"),
			pattern: "a",
			ok:      true,
		},
		{
			line:    []byte("apple123"),
			pattern: `\d`,
			ok:      true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.pattern, func(t *testing.T) {
			// t.Parallel()

			ok, err := matchLine(tt.line, tt.pattern)
			if err != nil {
				t.Fatal(err)
			}
			if tt.ok != ok {
				t.Fatalf("expected=%v, actual=%v\n", tt.ok, ok)
			}
		})
	}
}
