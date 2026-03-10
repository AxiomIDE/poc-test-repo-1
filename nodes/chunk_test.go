package nodes

import (
	"testing"

	gen "axiom-text-ops/gen"
)

func TestChunk(t *testing.T) {
	var results []*gen.TokensResult
	in := make(chan *gen.TextRequest, 1)
	in <- &gen.TextRequest{Text: "hello world"}
	close(in)

	err := Chunk(nil, nil, nil, in, func(r *gen.TokensResult) error {
		results = append(results, r)
		return nil
	})
	if err != nil {
		t.Fatalf("Chunk returned error: %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 frames, got %d", len(results))
	}
	if results[0].Tokens[0] != "hello" {
		t.Errorf("expected first token 'hello', got %q", results[0].Tokens[0])
	}
	if results[1].Tokens[0] != "world" {
		t.Errorf("expected second token 'world', got %q", results[1].Tokens[0])
	}
}
