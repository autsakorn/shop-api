package utils

// IndexOf to get index from slice
// Example
// data := []string{"onw", "twe", "three"}
// indexOf("two", data[:])
// output 1
func IndexOf(find string, data []string) int {
	for k, v := range data {
		if find == v {
			return k
		}
	}
	return -1
}
