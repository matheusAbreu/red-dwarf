package healthCheck

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
)

type TalkingDwarf struct {
	DwarfSaid	string
}

const (
	PathRouter = "/health"
)

// @Summary 						Check if the service is running.
// @Description 				Health cheack
// @Tag 								default
// @Produce 						json
// @Success 						200 {object} TalkingDwarf
// @Router 							/health [get]
func Check(context *fiber.Ctx) error {
	dwarf := TalkingDwarf{
		DwarfSaid: "It's all good and red around here.",
	}
	return context.Status(http.StatusOK).JSON(dwarf)
}