package types

type (
	MarkdownFile[Frontmatter any] struct {
		Frontmatter Frontmatter
		Body        Instructions
		//
		filename string
	}

	MemoryFile[State any] struct {
		Frontmatter MemoryFrontmatter
		Body        State
		//
		filename string
	}

	ToolFile struct {
		Frontmatter ToolFrontmatter
		Body        *Code
		//
		filename string
	}

	AgentFrontmatter struct {
		Metadata      `yaml:",inline"`
		Options       map[string]string `yaml:"options"`
		AllowedSkills []string          `yaml:"allowed-skills"`
	}

	IdentityFrontmatter struct {
		Metadata `yaml:",inline"`
	}

	InsightFrontmatter struct {
		Metadata `yaml:",inline"`
	}

	MemoryFrontmatter struct {
		Metadata `yaml:",inline"`
	}

	PromptFrontmatter struct {
		Metadata `yaml:",inline"`
	}

	SkillFrontmatter struct {
		Metadata     `yaml:",inline"`
		AllowedTools []string `yaml:"allowed-tools"`
	}

	SoulFrontmatter struct {
		Metadata `yaml:",inline"`
	}

	ToolFrontmatter struct {
		Metadata `yaml:",inline"`
		Input    *Schema `yaml:"input"`
		Output   *Schema `yaml:"output"`
	}

	UserFrontmatter struct {
		Metadata `yaml:",inline"`
	}

	Metadata struct {
		Name        string   `yaml:"name"`
		Description string   `yaml:"description"`
		Version     string   `yaml:"version"`
		License     string   `yaml:"license"`
		Keywords    Keywords `yaml:"keywords"`
	}

	Code struct {
		Language     string   `json:"language"`
		Body         []byte   `json:"body"`
		Dependencies []string `json:"dependencies"`
	}

	Variable struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Secret      bool   `json:"secret,omitempty"`
	}

	Schema struct {
		AnyOf            []*Schema          `json:"anyOf,omitempty"`
		Default          any                `json:"default,omitempty"`
		Description      string             `json:"description,omitempty"`
		Enum             []string           `json:"enum,omitempty"`
		Example          any                `json:"example,omitempty"`
		Format           string             `json:"format,omitempty"`
		Items            *Schema            `json:"items,omitempty"`
		MaxItems         *int64             `json:"maxItems,omitempty"`
		MaxLength        *int64             `json:"maxLength,omitempty"`
		MaxProperties    *int64             `json:"maxProperties,omitempty"`
		Maximum          *float64           `json:"maximum,omitempty"`
		MinItems         *int64             `json:"minItems,omitempty"`
		MinLength        *int64             `json:"minLength,omitempty"`
		MinProperties    *int64             `json:"minProperties,omitempty"`
		Minimum          *float64           `json:"minimum,omitempty"`
		Nullable         *bool              `json:"nullable,omitempty"`
		Pattern          string             `json:"pattern,omitempty"`
		Properties       map[string]*Schema `json:"properties,omitempty"`
		PropertyOrdering []string           `json:"propertyOrdering,omitempty"`
		Required         []string           `json:"required,omitempty"`
		Title            string             `json:"title,omitempty"`
		Type             string             `json:"type,omitempty"`
	}

	Layer struct {
		Name         string       `json:"name"`
		Description  string       `json:"description"`
		Instructions Instructions `json:"instructions"`
		Tools        Tools        `json:"tools,omitempty"`
	}

	Tool struct {
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		InputSchema  *Schema `json:"input_schema"`
		OutputSchema *Schema `json:"output_schema"`
		Handle       Handle  `json:"-"`
	}

	Blob struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Data []byte `json:"data"`
	}

	Source struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Type        string `json:"type"`
		URL         string `json:"url"`
	}

	Call struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Input  Input  `json:"input,omitempty"`
		Output Output `json:"output,omitempty"`
	}

	Part struct {
		Call   *Call   `json:"call,omitempty"`
		Blob   *Blob   `json:"blob,omitempty"`
		Source *Source `json:"source,omitempty"`
	}

	Message struct {
		Turn      Turn       `json:"turn"`
		Text      *string    `json:"text"`
		Structure *Structure `json:"structure,omitempty"`
		Parts     Parts      `json:"parts,omitempty"`
	}

	Price struct {
		Input  float64
		Output float64
	}

	Keywords     []string
	Instructions string
	Turn         string
	Modality     string

	Input     map[string]any
	Output    map[string]any
	Structure map[string]any
	Prices    map[Modality]Price

	Messages []*Message
	Layers   []*Layer
	Tools    []*Tool
	Parts    []*Part

	Handle func(Input) (Output, error)

	IFrontmatter interface {
		Validate() error
	}
)
