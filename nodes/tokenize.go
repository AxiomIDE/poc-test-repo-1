package nodes

import (
	"context"
	"strings"

	"axiom-text-ops/axiom"
	gen "axiom-text-ops/gen"
)

// Tokenize splits input text on whitespace and returns the individual tokens
// along with the total token count.
func Tokenize(ctx context.Context, ax axiom.Context, input *gen.TextRequest) (*gen.TokensResult, error) {
	ax.Log().Info("Hello, from Tokenize! Modified.")
	if val, ok := ax.Secrets().Get("MY_SECRET"); ok {
		ax.Log().Info("MY_SECRET resolved", "value", val)
	} else {
		ax.Log().Warn("MY_SECRET not found in secrets")
	}
	tokens := strings.Fields(input.GetText())
	return &gen.TokensResult{
		Tokens: tokens,
		Count:  int32(len(tokens)),
	}, nil
}
