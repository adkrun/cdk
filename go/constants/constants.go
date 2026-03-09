package constants

import (
	"path/filepath"
	"regexp"
)

const (
	NamePattern        = `^[a-z](?:[a-z0-9]|[.-][a-z0-9]){1,62}[a-z]$`
	DescriptionPattern = `^[a-zA-Z0-9\s.,!?'"()\-]{1,1000}[a-zA-Z0-9\s.,!?'"()\-]{0,24}$`

	VariableNamePattern        = `^[a-z](?:[a-z0-9]|[-][a-z0-9]){1,30}[a-z]$`
	VariableDescriptionPattern = `^[a-zA-Z0-9\s.,!?'"()\-]{1,256}$`

	VersionPattern = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-[a-zA-Z0-9-]+)?$`

	DefaultVersion = "0.1.0"
	DefaultLicense = "MIT"

	DefaultAgentFilename    = "AGENT.md"
	DefaultIdentityFilename = "IDENTITY.md"
	DefaultInsightFilename  = "INSIGHT.md"
	DefaultMemoryFilename   = "MEMORY.json"
	DefaultPromptFilename   = "PROMPT.md"
	DefaultSoulFilename     = "SOUL.md"
	DefaultSkillFilename    = "SKILL.md"
	DefaultUserFilename     = "USER.md"

	DefaultMemoriesPath = "memories"
	DefaultPromptsPath  = "prompts"
	DefaultSkillsPath   = "skills"
	DefaultScriptsPath  = "scripts"
	DefaultUsersPath    = "users"
)

var (
	NameRegex                = regexp.MustCompile(NamePattern)
	DescriptionRegex         = regexp.MustCompile(DescriptionPattern)
	VariableNameRegex        = regexp.MustCompile(VariableNamePattern)
	VariableDescriptionRegex = regexp.MustCompile(VariableDescriptionPattern)
	VersionRegex             = regexp.MustCompile(VersionPattern)
)

func GetAgentFilename(path string) string {
	return filepath.Join(path, DefaultAgentFilename)
}

func GetIdentityFilename(path string) string {
	return filepath.Join(path, DefaultIdentityFilename)
}

func GetInsightFilename(path, name string) string {
	return filepath.Join(path, DefaultMemoriesPath, name, DefaultInsightFilename)
}

func GetMemoryFilename(path, name string) string {
	return filepath.Join(path, DefaultMemoriesPath, name, DefaultMemoryFilename)
}

func GetPromptFilename(path, name string) string {
	return filepath.Join(path, DefaultPromptsPath, name, DefaultPromptFilename)
}

func GetSkillFilename(path, name string) string {
	return filepath.Join(path, DefaultSkillsPath, name, DefaultSkillFilename)
}

func GetToolFilename(path, skillName, name string) string {
	return filepath.Join(path, DefaultSkillsPath, skillName, DefaultScriptsPath, name)
}

func GetSoulFilename(path string) string {
	return filepath.Join(path, DefaultSoulFilename)
}

func GetUserFilename(path, name string) string {
	return filepath.Join(path, DefaultUsersPath, name, DefaultUserFilename)
}
