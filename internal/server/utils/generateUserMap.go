package utils

// GenerateUserMap generates a map of user:null
func GenerateUserMap(userPasswd map[string]string) map[string]string {
	userMap := make(map[string]string)
	for user := range userPasswd {
		userMap[user] = ""
	}
	return userMap
}
