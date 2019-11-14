package utils

// Int utils
type Int struct{}

// I Namespace for int slices utils
var I Int

// Index returns the index when element is found in slice, -1 when it is not
func (i Int) Index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include true when element is found, false when it is not
func (i Int) Include(vs []int, t int) bool {
	return i.Index(vs, t) >= 0
}

// Index returns the index when element is found in slice, -1 when it is not
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include true when element is found, false when it is not
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Map applies a function to a string collection
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
