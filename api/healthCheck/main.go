package healthCheck

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	PathRouter = "/health"
)

// @Summary 						Check if the service is running.
// @Description 				Health cheack
// @Tag 								default
// @Produce 						plain
// @Success 						200 {string} string "OK"
// @Router 							/health [get]
func Check(contex *fiber.Ctx) error {
	return contex.SendStatus(http.StatusOK)
}