package utils

// NullableFloat returns the float if not empty, null if empty
func NullableFloat(fl float64) *float64 {
	if fl == 0 {
		return nil
	}
	return &fl
}
