package tree_sitter_kanata_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/postsolar/tree-sitter-kanata"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_kanata.Language())
	if language == nil {
		t.Errorf("Error loading Kanata grammar")
	}
}
