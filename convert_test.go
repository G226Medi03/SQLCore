package sqlcore

import (
	"encoding/json"
	"testing"
)

func TestInt64(t *testing.T) {
	if got := Int64(42); got != 42 {
		t.Fatalf("int: %d", got)
	}
	if got := Int64(int64(7)); got != 7 {
		t.Fatalf("int64: %d", got)
	}
	if got := Int64(float64(3)); got != 3 {
		t.Fatalf("float64: %d", got)
	}
	if got := Int64(json.Number("9")); got != 9 {
		t.Fatalf("json.Number: %d", got)
	}
	if got := Int64("bad"); got != 0 {
		t.Fatalf("unsupported: %d", got)
	}
}
