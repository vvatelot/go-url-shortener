package config

import (
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuthConfig() basicauth.Config {
	return basicauth.Config{
		Users: map[string]string{
			"admin": ENV.AdminPassword,
		},
	}
}
