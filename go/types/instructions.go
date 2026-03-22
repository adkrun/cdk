package types

import (
	"regexp"
	"strings"

	"github.com/adkrun/cdk/go/constants"
)

func NewInstructions(value string) Instructions {
	cleaned := strings.TrimLeft(value, "\n")

	return Instructions(strings.TrimSpace(cleaned))
}

func NewInstructionsFromBytes(value []byte) Instructions {
	return NewInstructions(string(value))
}

func (i Instructions) Variables() ([]*Variable, error) {
	re := regexp.MustCompile(`\{\{([^|{}]+)\|([^{}]+)\}\}`)
	unique := make(map[string]bool)
	var variables []*Variable
	for _, match := range re.FindAllStringSubmatch(string(i), -1) {
		if len(match) > 2 {
			name := strings.TrimSpace(match[1])
			description := strings.TrimSpace(match[2])

			if !constants.VariableNameRegex.MatchString(name) || !constants.VariableDescriptionRegex.MatchString(description) {
				continue
			}

			if _, exists := unique[name]; !exists {
				unique[name] = true
				variable := NewVariable(name, description)
				if err := variable.Validate(); err != nil {
					return nil, err
				}
				variables = append(variables, variable)
			}
		}
	}

	return variables, nil
}

func (i Instructions) Fill(values map[string]string) string {
	result := string(i)
	for key, value := range values {
		pattern := regexp.MustCompile(`\{\{\s*` + regexp.QuoteMeta(key) + `\s*\|[^{}]+\}\}`)
		result = pattern.ReplaceAllLiteralString(result, value)
	}

	return result
}

func (i Instructions) String() string {
	return string(i)
}

func (i Instructions) Bytes() []byte {
	return []byte(i)
}
