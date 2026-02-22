package nodes_test

import (
	"context"
	"testing"

	gen "axiom-text-ops/gen"
	"axiom-text-ops/nodes"
)

// TESTS — delete this block when done ─────────────────────────────────────────
// Tests are required to publish this package. The publish pipeline runs your
// tests as a quality gate — a package will not be published if tests fail or
// do not meet the minimum requirements.
//
// Requirements checked before publishing:
//   - At least one test per node
//   - All tests must pass
//   - Output fields must be meaningfully asserted — not just error-checked
//
// The generated test below is a starting point. Replace the TODO comment with
// real assertions that verify your node returns correct data for known inputs.
// Think: given a specific input, what should the output fields contain?
//
// Run your tests locally at any time:
//   axiom test

func TestTokenize(t *testing.T) {
	ctx := context.Background()
	input := &gen.TextRequest{}

	got, err := nodes.Tokenize(ctx, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_ = got // TODO: assert output fields — e.g. if got.SomeField != "expected" { t.Errorf(...) }
}
