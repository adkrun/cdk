package types

import (
	"bytes"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-json"
)

func ReadMemoryFile[T any](filename string) (*MemoryFile[T], error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read the memoryfile: %s", filename)
	}

	return NewMemoryFileFromBytes[T](filename, content)
}

func NewMemoryFileFromBytes[T any](filename string, content []byte) (*MemoryFile[T], error) {
	// Standard Frontmatter usually looks like:
	// ---
	// yaml: data
	// ---
	//
	// { ... }

	parts := bytes.SplitN(content, []byte("---"), 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid format of the memoryfile, expected frontmatter and body separated by '---'")
	}

	var frontmatter MemoryFrontmatter
	if err := yaml.Unmarshal(parts[1], &frontmatter); err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	if err := frontmatter.Validate(); err != nil {
		return nil, fmt.Errorf("frontmatter validation failed: %w", err)
	}

	var body T
	if err := json.Unmarshal(parts[2], &body); err != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err)
	}

	return &MemoryFile[T]{
		Frontmatter: frontmatter,
		Body:        body,
		filename:    filename,
	}, nil
}

func (m *MemoryFile[T]) Save() error {
	frontmatter, err := yaml.Marshal(m.Frontmatter)
	if err != nil {
		return fmt.Errorf("failed to serialize frontmatter: %w", err)
	}

	body, err := json.Marshal(m.Body)
	if err != nil {
		return fmt.Errorf("failed to serialize body: %w", err)
	}

	content := bytes.Join([][]byte{
		[]byte("---"),
		frontmatter,
		[]byte("---"),
		body,
	}, []byte("\n"))

	if err := os.WriteFile(m.filename, content, 0644); err != nil {
		return fmt.Errorf("failed to write memoryfile: %w", err)
	}

	return nil
}
