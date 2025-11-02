package generators_test

import (
	"testing"

	"github.com/CSKU-Lab/generators"
)

func TestUUID(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("UUID() panicked: %v", r)
		}
	}()

	id := generators.UUID()

	if len(id) != 36 {
		t.Errorf("Expected UUID length of 36, got %d", len(id))
	}
}
