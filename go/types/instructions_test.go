package types

import "testing"

func TestInstructionsFill(t *testing.T) {
	t.Parallel()

	t.Run("fills new variable syntax", func(t *testing.T) {
		t.Parallel()

		input := Instructions("Hello {{agent-name|The current agent name}}. Topic: {{topic|Current discussion topic}}.")
		got := input.Fill(map[string]string{
			"agent-name": "copilot",
			"topic":      "regex updates",
		})

		want := "Hello copilot. Topic: regex updates."
		if got != want {
			t.Fatalf("unexpected fill output: got %q, want %q", got, want)
		}
	})

	t.Run("uses literal replacement value", func(t *testing.T) {
		t.Parallel()

		input := Instructions("Cost: {{price|Price value}}")
		got := input.Fill(map[string]string{"price": "$10"})

		want := "Cost: $10"
		if got != want {
			t.Fatalf("unexpected fill output: got %q, want %q", got, want)
		}
	})
}
