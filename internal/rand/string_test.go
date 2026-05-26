package rand

import (
	"testing"
)

func TestRandomStr(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"small length", 10},
		{"medium length", 32},
		{"large length", 256},
		{"length 2", 2},
		{"length 0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomStr(tt.length)
			if err != nil {
				t.Errorf("RandomStr() error = %v", err)
				return
			}
			if len(got) != tt.length {
				t.Errorf("RandomStr() length = %d, want %d", len(got), tt.length)
			}
		})
	}
}

func TestRandomStrUniqueness(t *testing.T) {
	results := make(map[string]bool)
	iterations := 10000
	length := 32

	for range iterations {
		str, err := RandomStr(length)
		if err != nil {
			t.Fatalf("RandomStr() failed: %v", err)
		}
		if results[str] {
			t.Errorf("RandomStr() generated duplicate: %s", str)
		}
		results[str] = true
	}

	if len(results) != iterations {
		t.Errorf("RandomStr() duplicates detected: got %d unique from %d calls", len(results), iterations)
	}
}
