package types

import "strings"

func NewCode(language string, body []byte, dependencies ...string) Code {
	return Code{
		Language:     strings.ToLower(language),
		Body:         body,
		Dependencies: dependencies,
	}
}

func NewGoCode(body []byte, dependencies ...string) Code {
	return NewCode("go", body, dependencies...)
}

func NewTsCode(body []byte, dependencies ...string) Code {
	return NewCode("ts", body, dependencies...)
}
