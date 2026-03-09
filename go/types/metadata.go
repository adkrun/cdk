package types

import (
	"fmt"

	"github.com/adkrun/Agentfile/go/constants"
)

func NewMetadata(name, description string, keyword ...string) Metadata {
	return Metadata{
		Name:        name,
		Description: description,
		Version:     constants.DefaultVersion,
		License:     constants.DefaultLicense,
		Keywords:    NewKeywords(keyword...),
	}
}

func (m Metadata) Validate() error {
	if m.Name == "" {
		return fmt.Errorf("name is required")
	}
	if m.Description == "" {
		return fmt.Errorf("description is required")
	}

	if !constants.NameRegex.MatchString(m.Name) {
		return fmt.Errorf("name must be 3 to 64 characters, start and end with a lowercase letter, and contain only lowercase letters, numbers, dots, or hyphens")
	}

	if !constants.DescriptionRegex.MatchString(m.Description) {
		return fmt.Errorf("description must be 1 to 1024 characters and contain only letters, numbers, spaces, or . , ! ? ' \" ( ) -")
	}

	if m.Version != "" && !constants.VersionRegex.MatchString(m.Version) {
		return fmt.Errorf("version must use semantic format: major.minor.patch (optional -suffix), for example: 1.2.3 or 1.2.3-beta")
	}

	return nil
}
