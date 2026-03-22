package types

import (
	"fmt"

	"github.com/adkrun/cdk/go/constants"
)

func NewCall(ID, name string, input Input, output Output) *Call {
	return &Call{
		ID:     ID,
		Name:   name,
		Input:  input,
		Output: output,
	}
}

func NewCallInput(ID, name string, input Input) *Call {
	return &Call{
		ID:    ID,
		Name:  name,
		Input: input,
	}
}

func NewCallOutput(ID, name string, output Output) *Call {
	return &Call{
		ID:     ID,
		Name:   name,
		Output: output,
	}
}

func (c *Call) Validate() error {
	if c.ID == "" {
		return fmt.Errorf("call ID is required")
	}

	if c.Name == "" {
		return fmt.Errorf("call name is required")
	}

	if !constants.NameRegex.MatchString(c.Name) {
		return fmt.Errorf("name must be 3 to 64 characters, start and end with a lowercase letter, and contain only lowercase letters, numbers, dots, or hyphens")
	}

	return nil
}
