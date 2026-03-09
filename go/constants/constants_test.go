package constants

import (
	"strings"
	"testing"
)

func TestRegexPatterns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		re      interface{ MatchString(string) bool }
		valid   []string
		invalid []string
	}{
		{
			name: "NameRegex",
			re:   NameRegex,
			valid: []string{
				"abc",
				"my-agenta",
				"my.agenta",
			},
			invalid: []string{
				"ab",
				"My-agent",
				"my-agent1",
				"my.agent2",
				"my..agent",
				"my--agent",
				"-myagent",
				"myagent-",
			},
		},
		{
			name: "DescriptionRegex",
			re:   DescriptionRegex,
			valid: []string{
				"A",
				"Clearly describes what this component does.",
				strings.Repeat("a", 1024),
			},
			invalid: []string{
				"",
				strings.Repeat("a", 1025),
				"contains@symbol",
			},
		},
		{
			name: "VariableNameRegex",
			re:   VariableNameRegex,
			valid: []string{
				"abc",
				"my-vara",
			},
			invalid: []string{
				"ab",
				"my.var",
				"my-var1",
				"my--var",
				"1myvar",
			},
		},
		{
			name: "VariableDescriptionRegex",
			re:   VariableDescriptionRegex,
			valid: []string{
				"x",
				"Variable used to configure retries.",
				strings.Repeat("b", 256),
			},
			invalid: []string{
				"",
				strings.Repeat("b", 257),
				"bad#char",
			},
		},
		{
			name: "VersionRegex",
			re:   VersionRegex,
			valid: []string{
				"1.2.3",
				"0.0.0",
				"2.10.34-alpha",
				"3.4.5-alpha-1",
			},
			invalid: []string{
				"v1.2.3",
				"1.2",
				"1.2.3.4",
				"1.2.3+build",
				"01.2.3",
				"1.02.3",
				"1.2.03",
				"1.2.3-",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for _, value := range tc.valid {
				if !tc.re.MatchString(value) {
					t.Fatalf("expected %q to match %s", value, tc.name)
				}
			}

			for _, value := range tc.invalid {
				if tc.re.MatchString(value) {
					t.Fatalf("expected %q to not match %s", value, tc.name)
				}
			}
		})
	}
}
