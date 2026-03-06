package nodes

import (
	"context"
	"strings"

	"axiom-text-ops/axiom"
	gen "axiom-text-ops/gen"
)

// Chunk splits the input text into individual word tokens and emits one
// TokensResult frame per word, demonstrating a pipeline node.
func Chunk(ctx context.Context, log axiom.Logger, in <-chan *gen.TextRequest, emit func(*gen.TokensResult) error) error {
	for input := range in {
		words := strings.Fields(input.GetText())
		for _, w := range words {
			if err := emit(&gen.TokensResult{
				Tokens: []string{w},
				Count:  1,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
