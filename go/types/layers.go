package types

import (
	"fmt"
	"strings"
)

func NewLayers(layers ...*Layer) Layers {
	return layers
}

func (l Layers) GetLayer(name string) (*Layer, bool) {
	for _, layer := range l {
		if layer.Name == name {
			return layer, true
		}
	}

	return nil, false
}

func (l Layers) WithTools() Layers {
	layers := Layers{}

	for _, layer := range l {
		if len(layer.Tools) > 0 {
			layers = append(layers, layer)
		}
	}

	return layers
}

func (l Layers) Instructions() Instructions {
	instructions := ""
	for _, layer := range l {
		instructions += fmt.Sprintf("# %s\n%s\n\n%s\n\n", layer.Name, layer.Description, layer.Instructions)
	}

	return Instructions(strings.TrimSpace(instructions))
}
