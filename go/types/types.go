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
		Body        Code
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
		Name        string   `yaml:"name" validate:"required,namepattern"`
		Description string   `yaml:"description" validate:"required,descriptionpattern"`
		Version     string   `yaml:"version" validate:"omitempty,versionpattern"`
		License     string   `yaml:"license" validate:""`
		Keywords    Keywords `yaml:"keywords" validate:""`
	}

	Code struct {
		Language     string   `json:"language"`
		Body         []byte   `json:"body"`
		Dependencies []string `json:"dependencies"`
	}

	Variable struct {
		Name        string `json:"name" validate:"required,variablenamepattern"`
		Description string `json:"description" validate:"required,variabledescriptionpattern"`
	}

	Schema struct {
		// Optional. The value should be validated against any (one or more) of the subschemas in the list.
		AnyOf []*Schema `json:"anyOf,omitempty"`
		// Optional. Default value of the data.
		Default any `json:"default,omitempty"`
		// Optional. The description of the data.
		Description string `json:"description,omitempty"`
		// Optional. Possible values of the element of primitive type with enum format. Examples:
		// 1. We can define direction as : {type:STRING, format:enum, enum:["EAST", NORTH",
		// "SOUTH", "WEST"]} 2. We can define apartment number as : {type:INTEGER, format:enum,
		// enum:["101", "201", "301"]}
		Enum []string `json:"enum,omitempty"`
		// Optional. Example of the object. Will only populated when the object is the root.
		Example any `json:"example,omitempty"`
		// Optional. The format of the data. Supported formats: for NUMBER type: "float", "double" for INTEGER type: "int32", "int64" for STRING type: "email", "byte", etc
		Format string `json:"format,omitempty"`
		// Optional. SCHEMA FIELDS FOR TYPE ARRAY Schema of the elements of Type.ARRAY.
		Items *Schema `json:"items,omitempty"`
		// Optional. Maximum number of the elements for Type.ARRAY.
		MaxItems *int64 `json:"maxItems,omitempty"`
		// Optional. Maximum length of the Type.STRING
		MaxLength *int64 `json:"maxLength,omitempty"`
		// Optional. Maximum number of the properties for Type.OBJECT.
		MaxProperties *int64 `json:"maxProperties,omitempty"`
		// Optional. Maximum value of the Type.INTEGER and Type.NUMBER
		Maximum *float64 `json:"maximum,omitempty"`
		// Optional. Minimum number of the elements for Type.ARRAY.
		MinItems *int64 `json:"minItems,omitempty"`
		// Optional. SCHEMA FIELDS FOR TYPE STRING Minimum length of the Type.STRING
		MinLength *int64 `json:"minLength,omitempty"`
		// Optional. Minimum number of the properties for Type.OBJECT.
		MinProperties *int64 `json:"minProperties,omitempty"`
		// Optional. Minimum value of the Type.INTEGER and Type.NUMBER.
		Minimum *float64 `json:"minimum,omitempty"`
		// Optional. Indicates if the value may be null.
		Nullable *bool `json:"nullable,omitempty"`
		// Optional. Pattern of the Type.STRING to restrict a string to a regular expression.
		Pattern string `json:"pattern,omitempty"`
		// Optional. SCHEMA FIELDS FOR TYPE OBJECT Properties of Type.OBJECT.
		Properties map[string]*Schema `json:"properties,omitempty"`
		// Optional. The order of the properties. Not a standard field in open API spec. Only used to support the order of the properties.
		PropertyOrdering []string `json:"propertyOrdering,omitempty"`
		// Optional. Required properties of Type.OBJECT.
		Required []string `json:"required,omitempty"`
		// Optional. The title of the Schema.
		Title string `json:"title,omitempty"`
		// Optional. The type of the data.
		Type string `json:"type,omitempty"`
	}

	Keywords     []string
	Instructions string

	IFrontmatter interface {
		Validate() error
	}
)
