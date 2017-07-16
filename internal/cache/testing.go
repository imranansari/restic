package cache

import (
	"testing"

	restic "github.com/restic/restic/internal"
	"github.com/restic/restic/internal/test"
)

// TestNewCache returns a cache in a temporary directory which is removed when
// cleanup is called.
func TestNewCache(t testing.TB) (*Cache, func()) {
	dir, cleanup := test.TempDir(t)
	t.Logf("created new cache at %v", dir)
	cache, err := New(restic.NewRandomID().String(), dir)
	if err != nil {
		t.Fatal(err)
	}
	return cache, cleanup
}