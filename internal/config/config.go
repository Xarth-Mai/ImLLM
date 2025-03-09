package config

// Init initializes the configuration from the yaml file.
func Init(addr *string, port *string, userPasswd *map[string]string) {
	*port = "4000"
	*userPasswd = map[string]string{
		"admin":  "password",
		"admin1": "password1",
	}
	*addr = "localhost"
}
