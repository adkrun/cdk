package types

import (
	"fmt"

	"github.com/adkrun/Agentfile/go/constants"
)

func NewVariable(name, description string) Variable {
	return Variable{
		Name:        name,
		Description: description,
	}
}

func (v Variable) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("variable name is required")
	}

	if v.Description == "" {
		return fmt.Errorf("variable description is required for variable: %s", v.Name)
	}

	if !constants.VariableNameRegex.MatchString(v.Name) {
		return fmt.Errorf("variable name (%s) must be 3 to 32 characters, start and end with a lowercase letter, and contain only lowercase letters, numbers, or hyphens", v.Name)
	}

	if !constants.VariableDescriptionRegex.MatchString(v.Description) {
		return fmt.Errorf("variable description (%s) must be 1 to 256 characters and contain only letters, numbers, spaces, or . , ! ? ' \" ( ) -", v.Description)
	}

	return nil
}

func (v Variable) String() string {
	return fmt.Sprintf("{{%s|%s}}", v.Name, v.Description)
}

func (v Variable) Bytes() []byte {
	return []byte(v.String())
}
