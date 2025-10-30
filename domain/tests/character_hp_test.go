package test

import (
	"herkansing/onion/domain"
	"testing"
)

func TestCalculateMaxHP(t *testing.T) {
	tests := []struct {
		name   string
		hitDie string
		level  int
		con    int
		wantHP int
	}{
		{"Level 2 Halfling Rogue (1d8, 12 CON)", "1d8", 2, 12, 15},
		{"Level 6 Hill Dwarf Barbarian (1d12, 14 CON)", "1d12", 6, 14, 59},
	}

	for _, tt := range tests {
		got := domain.CalculateMaxHP(tt.hitDie, tt.level, tt.con)
		if got != tt.wantHP {
			t.Errorf("%s: got %d, want %d", tt.name, got, tt.wantHP)
		}
	}
}
