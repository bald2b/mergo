package mergo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type StructIssue89 struct {
	Bool bool
}

func TestIssue89MergeFalseAtChildDefault(t *testing.T) {
	parent := StructIssue89{
		Bool: true,
	}
	child := StructIssue89{
		Bool: false,
	}
	Merge(&child, parent)
	assert.True(t, child.Bool)
}
func TestIssue89MergeFalseAtChildBoolIsNotEmpty(t *testing.T) {
	parent := StructIssue89{
		Bool: true,
	}
	child := StructIssue89{
		Bool: false,
	}
	Merge(&child, parent, WithBoolIsAlwaysNotEmpty)
	assert.False(t, child.Bool)
}

func TestIssue89MergeTrueAtChild(t *testing.T) {
	parent := StructIssue89{
		Bool: false,
	}
	child := StructIssue89{
		Bool: true,
	}
	Merge(&child, parent)
	assert.True(t, child.Bool)
}
