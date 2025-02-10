package transformers

import (
	"github.com/eminano/greenmask/pkg/generators"
)

type Transformer interface {
	GetRequiredGeneratorByteLength() int
	SetGenerator(g generators.Generator) error
}
