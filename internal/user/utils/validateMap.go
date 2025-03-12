package utils

func ValidateMap(Map map[string]string, key string, value string) bool {
	if value == "" || key == "" {
		return false
	}

	if storedPass, ok := Map[key]; ok && storedPass == value {
		return true
	}
	return false
}
