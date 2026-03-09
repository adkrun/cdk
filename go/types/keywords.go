package types

import "strings"

func NewKeywords(keyword ...string) Keywords {
	return Keywords(keyword)
}

func NewKeywordsFromString(value string) Keywords {
	parts := strings.Split(value, ",")
	var keywords []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			keywords = append(keywords, strings.ToLower(trimmed))
		}
	}

	return Keywords(keywords)
}

func NewKeywordsFromBytes(value []byte) Keywords {
	return NewKeywordsFromString(string(value))
}

func (k Keywords) String() string {
	return strings.Join(k, ", ")
}

func (k Keywords) Bytes() []byte {
	return []byte(k.String())
}
