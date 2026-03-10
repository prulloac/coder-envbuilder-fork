package devcontainer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// These tests intentionally live in package devcontainer (not
// devcontainer_test) so they can exercise unexported helper behavior directly.
// The external-package tests in devcontainer_test.go continue to validate
// user-facing behavior through the public API.

func TestBuildCanonicalToRefUnique(t *testing.T) {
	t.Parallel()

	canonicalToRefs := map[string][]string{
		"ghcr.io/example/features/a": {"ghcr.io/example/features/a:1"},
		"ghcr.io/example/features/b": {"ghcr.io/example/features/b:2"},
	}

	canonicalToRef, err := buildCanonicalToRef(canonicalToRefs)
	require.NoError(t, err)
	require.Equal(t, "ghcr.io/example/features/a:1", canonicalToRef["ghcr.io/example/features/a"])
	require.Equal(t, "ghcr.io/example/features/b:2", canonicalToRef["ghcr.io/example/features/b"])
}

func TestBuildCanonicalToRefAmbiguousDeterministicError(t *testing.T) {
	t.Parallel()

	canonicalToRefs := map[string][]string{
		"ghcr.io/example/features/late": {
			"ghcr.io/example/features/late:2.0.0",
			"ghcr.io/example/features/late:1.0.0",
		},
	}

	_, err := buildCanonicalToRef(canonicalToRefs)
	require.ErrorContains(t, err, "ambiguous canonical feature reference \"ghcr.io/example/features/late\"")
	require.ErrorContains(t, err, "ghcr.io/example/features/late:1.0.0, ghcr.io/example/features/late:2.0.0")
}
