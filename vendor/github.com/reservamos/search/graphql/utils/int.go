package utils

// NullableInt returns an int32 if not empty, null if empty
func NullableInt(i int32) *int32 {
	if i == 0 {
		return nil
	}
	return &i
}
