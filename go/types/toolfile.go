package types

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
)

func ReadToolFile(filename string, dependecies ...string) (*ToolFile, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read the toolfile: %s", filename)
	}

	return NewToolFileFromBytes(filename, content, dependecies...)
}

func NewToolFileFromBytes(filename string, content []byte, dependecies ...string) (*ToolFile, error) {
	// Standard Frontmatter usually looks like:
	// ---
	// yaml: data
	// ---
	//
	// code content

	parts := bytes.SplitN(content, []byte("---"), 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid format of the toolfile, expected frontmatter and body separated by '---'")
	}

	var frontmatter ToolFrontmatter
	if err := yaml.Unmarshal(parts[1], &frontmatter); err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	if err := frontmatter.Validate(); err != nil {
		return nil, fmt.Errorf("frontmatter validation failed: %w", err)
	}

	ext := strings.TrimPrefix(filepath.Ext(filename), ".")

	return &ToolFile{
		Frontmatter: frontmatter,
		Body:        NewCode(ext, parts[2], dependecies...),
		filename:    filename,
	}, nil
}

func (m *ToolFile) Save() error {
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
		return fmt.Errorf("failed to write toolfile: %w", err)
	}

	return nil
}
