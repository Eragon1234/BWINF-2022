package utils

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	type args[T any] struct {
		a []T
		i int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "testing basic remove",
			args: args[string]{a: []string{"this", "is", "a", "remove-this", "test"}, i: 3},
			want: []string{"this", "is", "a", "test"},
		},
		{
			name: "testing remove of first index",
			args: args[string]{a: []string{"remove-this", "this", "is", "a", "test"}, i: 0},
			want: []string{"this", "is", "a", "test"},
		},
		{
			name: "testing remove of last index",
			args: args[string]{a: []string{"this", "is", "a", "test", "remove-this"}, i: 4},
			want: []string{"this", "is", "a", "test"},
		},
		{
			name: "testing remove of not existing index",
			args: args[string]{a: []string{"this", "is", "a", "test"}, i: 8},
			want: []string{"this", "is", "a", "test"},
		},
		{
			name: "testing remove on empty slice",
			args: args[string]{a: []string{}, i: 0},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.a, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAndReturn(t *testing.T) {
	type args[T any] struct {
		a []T
		i int
	}
	type testCase[T any] struct {
		name  string
		args  args[T]
		want  []T
		want1 T
	}
	tests := []testCase[string]{
		{
			name:  "testings basic remove and return",
			args:  args[string]{a: []string{"this", "is", "a", "remove-this", "test"}, i: 3},
			want:  []string{"this", "is", "a", "test"},
			want1: "remove-this",
		},
		{
			name:  "testings remove of first index",
			args:  args[string]{a: []string{"remove-this", "this", "is", "a", "test"}, i: 0},
			want:  []string{"this", "is", "a", "test"},
			want1: "remove-this",
		},
		{
			name:  "testings remove of last index",
			args:  args[string]{a: []string{"this", "is", "a", "test", "remove-this"}, i: 4},
			want:  []string{"this", "is", "a", "test"},
			want1: "remove-this",
		},
		{
			name:  "testings remove of not existing index",
			args:  args[string]{a: []string{"this", "is", "a", "test"}, i: 8},
			want:  []string{"this", "is", "a", "test"},
			want1: "",
		},
		{
			name:  "testings remove on empty slice",
			args:  args[string]{a: []string{}, i: 0},
			want:  []string{},
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RemoveAndReturn(tt.args.a, tt.args.i)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAndReturn() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RemoveAndReturn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
