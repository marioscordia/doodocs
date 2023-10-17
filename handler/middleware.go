package handler

import (
	"github.com/gofiber/fiber/v2"
)


func (h *Handler) LogRequest(c *fiber.Ctx) error {
	h.infoLogger.Printf("%s - %s %s %s\n", c.IP(), c.Protocol(), c.Method(), c.OriginalURL())
	return c.Next()
}