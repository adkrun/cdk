package types

import (
	"bytes"
	"fmt"
	"os"

	"github.com/adkrun/cdk/go/constants"
	"github.com/goccy/go-yaml"
)

func ReadMarkdownFile[T IFrontmatter](filename string) (*MarkdownFile[T], error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read the markdownfile: %s", filename)
	}

	return NewMarkdownFileFromBytes[T](filename, content)
}

func NewMarkdownFileFromBytes[T IFrontmatter](filename string, content []byte) (*MarkdownFile[T], error) {
	// Standard Frontmatter usually looks like:
	// ---
	// yaml: data
	// ---
	//
	// body content

	parts := bytes.SplitN(content, []byte("---"), 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid format of the markdownfile, expected frontmatter and body separated by '---'")
	}

	var frontmatter T
	if err := yaml.Unmarshal(parts[1], &frontmatter); err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	if err := frontmatter.Validate(); err != nil {
		return nil, fmt.Errorf("frontmatter validation failed: %w", err)
	}

	return &MarkdownFile[T]{
		Frontmatter: frontmatter,
		Body:        NewInstructionsFromBytes(parts[2]),
		filename:    filename,
	}, nil
}

func ReadAgentFile(path string) (*MarkdownFile[AgentFrontmatter], error) {
	return ReadMarkdownFile[AgentFrontmatter](constants.GetAgentFilename(path))
}

func ReadIdentityFile(path string) (*MarkdownFile[IdentityFrontmatter], error) {
	return ReadMarkdownFile[IdentityFrontmatter](constants.GetIdentityFilename(path))
}

func ReadInsightFile(path, name string) (*MarkdownFile[InsightFrontmatter], error) {
	return ReadMarkdownFile[InsightFrontmatter](constants.GetInsightFilename(path, name))
}

func ReadPromptFile(path, name string) (*MarkdownFile[PromptFrontmatter], error) {
	return ReadMarkdownFile[PromptFrontmatter](constants.GetPromptFilename(path, name))
}

func ReadSkillFile(path, name string) (*MarkdownFile[SkillFrontmatter], error) {
	return ReadMarkdownFile[SkillFrontmatter](constants.GetSkillFilename(path, name))
}

func ReadSoulFile(path string) (*MarkdownFile[SoulFrontmatter], error) {
	return ReadMarkdownFile[SoulFrontmatter](constants.GetSoulFilename(path))
}

func ReadUserFile(path, name string) (*MarkdownFile[UserFrontmatter], error) {
	return ReadMarkdownFile[UserFrontmatter](constants.GetUserFilename(path, name))
}

func (m *MarkdownFile[T]) Save() error {
	frontmatter, err := yaml.Marshal(m.Frontmatter)
	if err != nil {
		return fmt.Errorf("failed to marshal frontmatter: %w", err)
	}

	content := bytes.Join([][]byte{
		[]byte("---"),
		frontmatter,
		[]byte("---"),
		m.Body.Bytes(),
	}, []byte("\n"))

	return os.WriteFile(m.filename, content, 0644)
}
