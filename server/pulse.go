package server

import (
	"fmt"

	"github.com/gopulse/pulse"
)

// pulse
func startPulse(port int) {
	app := pulse.New()
	router := pulse.NewRouter()

	app.Router = router

	router.Get("/hello", func(c *pulse.Context) error {
		c.String(messageStr)
		return nil
	})

	app.Run(fmt.Sprintf(":%d", port))
}
