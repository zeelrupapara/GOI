package utils

func Contains(item string, items []string) bool {
	for _, i := range items{
		if i == item{
			return true
		}
	}
	return false
}
