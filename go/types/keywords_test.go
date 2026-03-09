package types

import (
	"reflect"
	"testing"
)

func TestNewKeywords(t *testing.T) {
	t.Parallel()

	got := NewKeywords("AI", "agent", "tooling")
	want := Keywords{"AI", "agent", "tooling"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected keywords: got %+v, want %+v", got, want)
	}
}

func TestNewKeywordsFromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  Keywords
	}{
		{
			name:  "normalizes and trims",
			input: "AI, Agent, TOOLING",
			want:  Keywords{"ai", "agent", "tooling"},
		},
		{
			name:  "skips empty entries",
			input: " ai, ,  ,agent ,, tooling ",
			want:  Keywords{"ai", "agent", "tooling"},
		},
		{
			name:  "empty input",
			input: "",
			want:  nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := NewKeywordsFromString(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("unexpected keywords: got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestNewKeywordsFromBytes(t *testing.T) {
	t.Parallel()

	got := NewKeywordsFromBytes([]byte("One, TWO, three"))
	want := Keywords{"one", "two", "three"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected keywords: got %+v, want %+v", got, want)
	}
}

func TestKeywordsString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   Keywords
		want string
	}{
		{
			name: "joins with comma and space",
			in:   Keywords{"ai", "agent", "tooling"},
			want: "ai, agent, tooling",
		},
		{
			name: "empty keywords",
			in:   Keywords{},
			want: "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.in.String()
			if got != tc.want {
				t.Fatalf("unexpected string: got %q, want %q", got, tc.want)
			}
		})
	}
}
