package tree_sitter_kanata_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_kanata "github.com/postsolar/tree-sitter-kanata/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_kanata.Language())
	if language == nil {
		t.Errorf("Error loading kanata grammar")
	}
}
