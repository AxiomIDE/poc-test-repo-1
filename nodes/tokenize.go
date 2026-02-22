package nodes

import (
	"context"
	"strings"

	gen "axiom-text-ops/gen"
)

// Tokenize splits input text on whitespace and returns the individual tokens
// along with the total token count.
func Tokenize(ctx context.Context, input *gen.TextRequest) (*gen.TokensResult, error) {
	tokens := strings.Fields(input.GetText())
	return &gen.TokensResult{
		Tokens: tokens,
		Count:  int32(len(tokens)),
	}, nil
}
