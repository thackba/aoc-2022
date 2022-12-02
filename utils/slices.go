package utils

func FindEmptyLine(content []string) int {
	for i, line := range content {
		if line == "" {
			return i
		}
	}
	return -1
}
