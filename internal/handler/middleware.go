package handler

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"backend/internal/util/config"
)

var Cors = func() fiber.Handler {
	// origins is the value of allowed CORS addresses, separated by comma (,).
	origins := ""
	for i, s := range config.C.Cors {
		origins += s
		if i < len(config.C.Cors)-1 {
			origins += ", "
		}
	}

	config := cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(config)
}()

var Jwt = func() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.C.Secret)},
	})
}()
