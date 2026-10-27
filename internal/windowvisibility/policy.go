package windowvisibility

import "time"

const ReassertInterval = 2 * time.Second

// ShouldReassertTopmost periodically restores a visible pet to the topmost
// window band, without showing it after the user explicitly hid it.
func ShouldReassertTopmost(visible bool, last, now time.Time) bool {
	return visible && (last.IsZero() || now.Sub(last) >= ReassertInterval)
}
