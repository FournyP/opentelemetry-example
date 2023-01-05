package controllers

import (
	"fmt"

	"test/usecases"

	"github.com/gofiber/fiber/v2"
)

type PingController struct {
	pingUseCase *usecases.PingUseCase
}

func NewPingController(pingUseCase *usecases.PingUseCase) *PingController {
	return &PingController{
		pingUseCase: pingUseCase,
	}
}

func (controller *PingController) Ping(c *fiber.Ctx) error {
	c.Accepts("json", "text")

	pingMessage := new(usecases.PingMessage)

	if err := c.QueryParser(pingMessage); err != nil {
		return fmt.Errorf("invalid format : %v", err)
	}

	response, err := controller.pingUseCase.Execute(pingMessage)

	if err != nil {
		return err
	}

	return c.JSON(response)
}
