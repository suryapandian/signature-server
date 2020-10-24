package config

import (
	"os"
)

var (
	PORT        string
	PRIVATE_KEY string
)

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	PRIVATE_KEY = os.Getenv("PRIVATE_KEY")
	if PRIVATE_KEY == "" {
		PRIVATE_KEY = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIEFMEZrmlYxczXKFxIlNvNGR5JQvDhTkLovJYxwQd3ua
-----END PRIVATE KEY-----`
	}
}
