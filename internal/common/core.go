package common

import (
	olc "github.com/google/open-location-code/go"
)

// Fetch olc code from lat/lng
func (l *Location) OLC(scope int) string {
	return olc.Encode(float64(l.GetLatitude()), float64(l.GetLongitude()), scope)
}

// Checks if Enviornment is Development
func (e Environment) IsDev() bool {
	return e == Environment_DEVELOPMENT
}

// Checks if Enviornment is Development
func (e Environment) IsProd() bool {
	return e == Environment_PRODUCTION
}
