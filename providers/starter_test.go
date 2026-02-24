package providers

import (
	"testing"
)

func TestInitStarterProvider(t *testing.T) {
	sp, err := InitStarterProvider()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sp.starterMap) == 0 {
		t.Error("expected non-empty pet map")
	}
}
