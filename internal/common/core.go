package common

import (
	olc "github.com/google/open-location-code/go"
)

// ** ─── Location MANAGEMENT ────────────────────────────────────────────────────────
func (l *Location) OLC(scope int) string {
	return olc.Encode(float64(l.GetLatitude()), float64(l.GetLongitude()), scope)
}
