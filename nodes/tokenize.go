package nodes

import (
	"context"
	"strings"

	"axiom-text-ops/axiom"
	gen "axiom-text-ops/gen"
)

// Tokenize splits input text on whitespace and returns the individual tokens
// along with the total token count.
func Tokenize(ctx context.Context, log axiom.Logger, input *gen.TextRequest) (*gen.TokensResult, error) {
	tokens := strings.Fields(input.GetText())
	return &gen.TokensResult{
		Tokens: tokens,
		Count:  int32(len(tokens)),
	}, nil
}
