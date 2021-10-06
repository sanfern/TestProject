package main

import (
	"testing"
)

func TestAddValid(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "Positive Values",
			a:    10,
			b:    20,
			want: 30,
		}, {
			name: "Negative Values",
			a:    -10,
			b:    -20,
			want: -30,
		},
	}

	// skip if file DNE
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("NewBpfProgram() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
