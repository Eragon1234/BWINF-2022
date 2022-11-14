package utils

import "testing"

func TestIndexOf(t *testing.T) {
	type args[T comparable] struct {
		s []T
		e T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[string]{
		{
			name: "testing basic indexOf",
			args: args[string]{s: []string{"this", "is", "a", "find-this", "test"}, e: "find-this"},
			want: 3,
		},
		{
			name: "testing indexOf on first element",
			args: args[string]{s: []string{"find-this", "this", "is", "a", "test"}, e: "find-this"},
			want: 0,
		},
		{
			name: "testing indexOf on last element",
			args: args[string]{s: []string{"this", "is", "a", "test", "find-this"}, e: "find-this"},
			want: 4,
		},
		{
			name: "testing find of not existing element",
			args: args[string]{s: []string{"this", "is", "a", "test"}, e: "find-this"},
			want: -1,
		},
		{
			name: "testing indexOf on empty slice",
			args: args[string]{s: []string{}, e: "find-this"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
