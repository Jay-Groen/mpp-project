package test

import (
    "herkansing/onion/domain"
    "testing"
)

func TestCalculateConModifier(t *testing.T) {
    tests := []struct{
        con int
        want int
    }{
        {10, 0},
        {11, 0},
        {12, 1},
        {9, -1},
        {8, -1},
        {7, -2},
        {1, -5},
        {30, 10},
    }

    for _, tt := range tests {
        if got := domain.CalculateConModifier(tt.con); got != tt.want {
            t.Fatalf("CalculateConModifier(%d) = %d, want %d", tt.con, got, tt.want)
        }
    }
}

func TestCalculateMaxHP_EdgeCases(t *testing.T) {
    if got := domain.CalculateMaxHP("1d8", 0, 10); got != 0 {
        t.Fatalf("CalculateMaxHP(level 0) = %d, want 0", got)
    }

    // Very low CON shouldn't drop below 1 HP
    if got := domain.CalculateMaxHP("1d6", 1, 1); got < 1 {
        t.Fatalf("CalculateMaxHP should never be < 1 for positive level; got %d", got)
    }
}
