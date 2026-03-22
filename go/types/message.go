package types

func NewMessage(turn Turn, text *string, structure *Structure, parts ...*Part) *Message {
	return &Message{
		Turn:      turn,
		Text:      text,
		Structure: structure,
		Parts:     parts,
	}
}

func NewMessageText(turn Turn, text string, parts ...*Part) *Message {
	return &Message{
		Turn:  turn,
		Text:  &text,
		Parts: parts,
	}
}

func NewMessageStructure(structure *Structure, parts ...*Part) *Message {
	return &Message{
		Turn:      TurnAgent,
		Structure: structure,
		Parts:     parts,
	}
}

func NewMessageSources(text string, sources ...*Source) *Message {
	parts := make([]*Part, len(sources))
	for i, source := range sources {
		parts[i] = &Part{Source: source}
	}

	return &Message{
		Turn:  TurnAgent,
		Text:  &text,
		Parts: parts,
	}
}
