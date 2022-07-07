package config

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuthConfig() basicauth.Config {
	return basicauth.Config{
		Users: map[string]string{
			"admin": os.Getenv("ADMIN_PASSWORD"),
		},
	}
}
