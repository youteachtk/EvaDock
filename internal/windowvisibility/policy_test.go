package windowvisibility

import (
	"testing"
	"time"
)

func TestShouldReassertTopmost(t *testing.T) {
	now := time.Date(2026, 10, 27, 12, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		visible bool
		last    time.Time
		want    bool
	}{
		{"first time visible", true, time.Time{}, true},
		{"hidden never shown", false, time.Time{}, false},
		{"visible before interval", true, now.Add(-time.Second), false},
		{"visible at interval", true, now.Add(-2 * time.Second), true},
		{"visible after interval", true, now.Add(-3 * time.Second), true},
		{"hidden after interval", false, now.Add(-3 * time.Second), false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ShouldReassertTopmost(tc.visible, tc.last, now)
			if got != tc.want {
				t.Errorf("ShouldReassertTopmost(%v, %v, %v) = %v; want %v", tc.visible, tc.last, now, got, tc.want)
			}
		})
	}
}
