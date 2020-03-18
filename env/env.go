package env

import "os"

// IsProduction check if environment is production
func IsProduction() bool {
	return (os.Getenv("GO_ENVIRONMENT") == "production")
}
