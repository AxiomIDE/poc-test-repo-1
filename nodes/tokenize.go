package nodes

import (
	"context"
	"strings"

	"axiom-text-ops/axiom"
	gen "axiom-text-ops/gen"
)

// Tokenize splits input text on whitespace and returns the individual tokens
// along with the total token count.
func Tokenize(ctx context.Context, log axiom.Logger, secrets axiom.Secrets, input *gen.TextRequest) (*gen.TokensResult, error) {
	log.Info("Hello, from Tokenize! Modified.")
	if val, ok := secrets.Get("MY_SECRET"); ok {
		log.Info("MY_SECRET resolved", "value", val)
	} else {
		log.Warn("MY_SECRET not found in secrets")
	}
	tokens := strings.Fields(input.GetText())
	return &gen.TokensResult{
		Tokens: tokens,
		Count:  int32(len(tokens)),
	}, nil
}
