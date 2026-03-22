package types

func NewLayer(name, description, instructions string, tools ...*Tool) *Layer {
	return &Layer{
		Name:         name,
		Description:  description,
		Instructions: NewInstructions(instructions),
		Tools:        tools,
	}
}

func (l *Layer) GetTool(name string) (*Tool, bool) {
	for _, tool := range l.Tools {
		if tool.Name == name {
			return tool, true
		}
	}

	return nil, false
}
