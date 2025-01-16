package config

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

var STORE *session.Store

func InitSession() {
	STORE = session.New()
}
