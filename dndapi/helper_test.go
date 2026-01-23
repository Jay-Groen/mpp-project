package dndapi

import "testing"

func TestNameToIndex(t *testing.T) {
    tests := []struct{
        in string
        want string
    }{
        {"Magic Missile", "magic-missile"},
        {"Tasha's Hideous Laughter", "tashas-hideous-laughter"},
        {"  Weird  Name  ", "--weird--name--"}, // current behavior: spaces become hyphens without trimming
    }

    for _, tt := range tests {
        if got := NameToIndex(tt.in); got != tt.want {
            t.Fatalf("NameToIndex(%q) = %q, want %q", tt.in, got, tt.want)
        }
    }
}
